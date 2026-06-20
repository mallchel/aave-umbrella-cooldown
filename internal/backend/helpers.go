package backend

import (
	"1-task/internal/storage/postgres"
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"image/color"
	"net/http"
	"strconv"
	"strings"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func writeJSON(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(body)
}

func badParam(message string) error {
	return &ParamError{message: message}
}

type ParamError struct {
	message string
}

func (e *ParamError) Error() string {
	return e.message
}

type dayTicks struct{}

func (dayTicks) Ticks(min, max float64) []plot.Tick {
	if max <= min {
		return nil
	}

	const secondsPerDay = 24 * 60 * 60
	spanDays := int((max - min) / secondsPerDay)

	// Keep roughly up to 9 labels to avoid overcrowding.
	stepDays := 1
	candidates := []int{1, 2, 3, 7, 14, 30, 60, 90}
	for _, c := range candidates {
		if spanDays/c <= 15 {
			stepDays = c
			break
		}
		stepDays = c
	}

	start := time.Unix(int64(min), 0).UTC().Truncate(24 * time.Hour)
	for start.Unix() < int64(min) {
		start = start.Add(24 * time.Hour)
	}

	ticks := make([]plot.Tick, 0, spanDays/stepDays+2)
	for t := start; float64(t.Unix()) <= max; t = t.Add(time.Duration(stepDays) * 24 * time.Hour) {
		ticks = append(ticks, plot.Tick{
			Value: float64(t.Unix()),
			Label: t.Format("2006-01-02"),
		})
	}

	return ticks
}

type usdtTicks struct{}

func (usdtTicks) Ticks(min, max float64) []plot.Tick {
	ticks := plot.DefaultTicks{}.Ticks(min, max)
	for i := range ticks {
		stringV := strconv.FormatFloat(ticks[i].Value, 'f', 6, 64)
		ticks[i].Label = addThousandsSeparators(strings.TrimRight(strings.TrimRight(stringV, "0"), "."))
	}
	return ticks
}

func addThousandsSeparators(s string) string {
	negative := strings.HasPrefix(s, "-")
	if negative {
		s = s[1:]
	}

	parts := strings.SplitN(s, ".", 2)
	intPart := parts[0]

	for i := len(intPart) - 3; i > 0; i -= 3 {
		intPart = intPart[:i] + "," + intPart[i:]
	}

	if len(parts) == 2 && parts[1] != "" {
		intPart += "." + parts[1]
	}

	if negative {
		return "-" + intPart
	}
	return intPart
}

func buildRequestedChartSVG(points []postgres.DailyFlowPoint) (string, error) {
	p := plot.New()
	p.Title.Text = "Umbrella Requested Volumes"
	p.X.Label.Text = "Day"
	p.Y.Label.Text = "Amount (USDT)"
	p.X.Tick.Marker = dayTicks{}
	p.Y.Tick.Marker = usdtTicks{}
	p.Add(plotter.NewGrid())

	requested := make(plotter.XYs, 0, len(points))

	for _, point := range points {
		x := float64(point.Day.Unix())
		requested = append(requested, plotter.XY{X: x, Y: point.Requested})
	}

	requestedLine, err := plotter.NewLine(requested)
	if err != nil {
		return "", fmt.Errorf("build requested line: %w", err)
	}
	requestedLine.Color = color.RGBA{R: 191, G: 77, B: 53, A: 255}
	requestedLine.Width = vg.Points(2)

	p.Add(requestedLine)
	p.Legend.Add("Requested", requestedLine)
	p.Legend.Top = true

	writerTo, err := p.WriterTo(11*vg.Inch, 4.6*vg.Inch, "svg")
	if err != nil {
		return "", fmt.Errorf("create svg writer: %w", err)
	}

	var buf bytes.Buffer
	if _, err := writerTo.WriteTo(&buf); err != nil {
		return "", fmt.Errorf("write svg: %w", err)
	}

	return buf.String(), nil
}

func buildWithdrawnChartSVG(points []postgres.DailyFlowPoint) (string, error) {
	p := plot.New()
	p.Title.Text = "Umbrella Withdrawn Volumes"
	p.X.Label.Text = "Day"
	p.Y.Label.Text = "Amount (USDT)"
	p.X.Tick.Marker = dayTicks{}
	p.Y.Tick.Marker = usdtTicks{}
	p.Add(plotter.NewGrid())

	withdrawn := make(plotter.XYs, 0, len(points))
	for _, point := range points {
		x := float64(point.Day.Unix())
		withdrawn = append(withdrawn, plotter.XY{X: x, Y: point.Withdrawn})
	}

	withdrawnLine, err := plotter.NewLine(withdrawn)
	if err != nil {
		return "", fmt.Errorf("build withdrawn line: %w", err)
	}
	withdrawnLine.Color = color.RGBA{R: 18, G: 102, B: 82, A: 255}
	withdrawnLine.Width = vg.Points(2)

	p.Add(withdrawnLine)
	p.Legend.Add("Withdrawn", withdrawnLine)
	p.Legend.Top = true

	writerTo, err := p.WriterTo(11*vg.Inch, 4.6*vg.Inch, "svg")
	if err != nil {
		return "", fmt.Errorf("create svg writer: %w", err)
	}

	var buf bytes.Buffer
	if _, err := writerTo.WriteTo(&buf); err != nil {
		return "", fmt.Errorf("write svg: %w", err)
	}

	return buf.String(), nil
}

func buildRequestCountChartSVG(points []postgres.DailyFlowPoint) (string, error) {
	p := plot.New()
	p.Title.Text = "Umbrella Withdraw Request Count"
	p.X.Label.Text = "Day"
	p.Y.Label.Text = "Count"
	p.X.Tick.Marker = dayTicks{}
	p.Y.Tick.Marker = usdtTicks{}
	p.Add(plotter.NewGrid())

	requestCount := make(plotter.XYs, 0, len(points))
	for _, point := range points {
		x := float64(point.Day.Unix())
		requestCount = append(requestCount, plotter.XY{X: x, Y: point.RequestCount})
	}

	requestCountLine, err := plotter.NewLine(requestCount)
	if err != nil {
		return "", fmt.Errorf("build request count line: %w", err)
	}
	requestCountLine.Color = color.RGBA{R: 18, G: 102, B: 82, A: 255}
	requestCountLine.Width = vg.Points(2)

	p.Add(requestCountLine)
	p.Legend.Add("Request Count", requestCountLine)
	p.Legend.Top = true

	writerTo, err := p.WriterTo(11*vg.Inch, 4.6*vg.Inch, "svg")
	if err != nil {
		return "", fmt.Errorf("create svg writer: %w", err)
	}

	var buf bytes.Buffer
	if _, err := writerTo.WriteTo(&buf); err != nil {
		return "", fmt.Errorf("write svg: %w", err)
	}

	return buf.String(), nil
}

type chartPageData struct {
	RenderedAt      time.Time
	RequestedSVG    template.HTML
	WithdrawnSVG    template.HTML
	RequestCountSVG template.HTML
}

var chartPageTpl = template.Must(template.New("home").Parse(`<!doctype html>
<html lang="en">
<head>
	<meta charset="utf-8" />
	<meta name="viewport" content="width=device-width, initial-scale=1" />
	<title>Umbrella Queue</title>
	<style>
		:root {
			--bg-1: #f8f4ea;
			--bg-2: #efe5d0;
			--ink: #2d2a23;
			--muted: #6c6658;
			--card: rgba(255, 255, 255, 0.72);
			--border: rgba(58, 50, 36, 0.14);
		}
		body {
			margin: 0;
			color: var(--ink);
			font-family: "Avenir Next", "Gill Sans", "Trebuchet MS", sans-serif;
			background: radial-gradient(circle at 20% 0%, var(--bg-1), var(--bg-2));
			min-height: 100vh;
		}
		.wrap {
			max-width: 1100px;
			margin: 0 auto;
			padding: 32px 18px 44px;
		}
		.card {
			background: var(--card);
			border: 1px solid var(--border);
			border-radius: 18px;
			padding: 20px;
			box-shadow: 0 16px 44px rgba(62, 49, 22, 0.08);
		}
		h1 {
			margin: 0 0 8px;
			font-size: 34px;
			letter-spacing: 0.02em;
		}
		.muted {
			margin: 0;
			color: var(--muted);
		}
		.chart {
			margin-top: 18px;
			overflow-x: auto;
		}
		.footer {
			margin-top: 12px;
			color: var(--muted);
			font-size: 14px;
		}
	</style>
</head>
<body>
	<main class="wrap">
		<section class="card">
			<h1>Aave Umbrella Queue</h1>
			<p class="muted">Daily requested volume to withdraw</p>
			<div class="chart">{{.RequestedSVG}}</div>
			<p class="muted">Daily withdrawn volume</p>
			<div class="chart">{{.WithdrawnSVG}}</div>
			<p class="muted">Daily request count to withdraw</p>
			<div class="chart">{{.RequestCountSVG}}</div>
			<p class="footer">Rendered at: {{.RenderedAt.Format "2006-01-02 15:04:05 MST"}}</p>
		</section>
	</main>
</body>
</html>`))
