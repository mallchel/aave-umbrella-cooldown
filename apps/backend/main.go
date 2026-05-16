package main

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"image/color"
	"log"
	"net/http"
	"time"

	"1-task/internal/envutil"
	"1-task/internal/storage/postgres"

	"github.com/go-chi/chi/v5"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type pageData struct {
	RenderedAt time.Time
	SVG        template.HTML
}

var pageTpl = template.Must(template.New("home").Parse(`<!doctype html>
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
			<p class="muted">Daily volumes from PostgreSQL. Queued vs withdrawable dates.</p>
			<div class="chart">{{.SVG}}</div>
			<p class="footer">Rendered at: {{.RenderedAt.Format "2006-01-02 15:04:05 MST"}}</p>
		</section>
	</main>
</body>
</html>`))

func main() {
	addr := envutil.Get("HTTP_ADDR", ":8080")
	dsn := envutil.Get("POSTGRES_DSN", "postgresql://umbrella_user:umbrella_pass@localhost:5432/umbrella_db?sslmode=disable")

	repo, err := postgres.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("init postgres repository: %v", err)
	}
	defer func() {
		if err := repo.Close(); err != nil {
			log.Printf("close postgres repository: %v", err)
		}
	}()

	r := chi.NewRouter()

	r.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"status":"ok"}`))
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		renderHome(w, r, repo)
	})

	r.Get("/chart", func(w http.ResponseWriter, r *http.Request) {
		renderHome(w, r, repo)
	})

	log.Printf("backend listening on %s", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal(err)
	}
}

func renderHome(w http.ResponseWriter, r *http.Request, repo *postgres.Repository) {
	points, err := repo.ListDailyQueuePoints(r.Context())
	if err != nil {
		http.Error(w, fmt.Sprintf("query chart data: %v", err), http.StatusInternalServerError)
		return
	}

	svg, err := buildChartSVG(points)
	if err != nil {
		http.Error(w, fmt.Sprintf("render chart: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := pageTpl.Execute(w, pageData{RenderedAt: time.Now(), SVG: template.HTML(svg)}); err != nil {
		http.Error(w, fmt.Sprintf("render page: %v", err), http.StatusInternalServerError)
		return
	}
}

func buildChartSVG(points []postgres.DailyQueuePoint) (string, error) {
	p := plot.New()
	p.Title.Text = "Umbrella Queue Daily Volumes"
	p.X.Label.Text = "Day"
	p.Y.Label.Text = "Amount (raw units)"
	p.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02"}
	p.Add(plotter.NewGrid())

	queued := make(plotter.XYs, 0, len(points))
	withdrawable := make(plotter.XYs, 0, len(points))
	for _, point := range points {
		x := float64(point.Day.Unix())
		queued = append(queued, plotter.XY{X: x, Y: point.QueuedVolume})
		withdrawable = append(withdrawable, plotter.XY{X: x, Y: point.WithdrawableVolume})
	}

	queuedLine, err := plotter.NewLine(queued)
	if err != nil {
		return "", fmt.Errorf("build queued line: %w", err)
	}
	queuedLine.Color = color.RGBA{R: 191, G: 77, B: 53, A: 255}
	queuedLine.Width = vg.Points(2)

	withdrawableLine, err := plotter.NewLine(withdrawable)
	if err != nil {
		return "", fmt.Errorf("build withdrawable line: %w", err)
	}
	withdrawableLine.Color = color.RGBA{R: 18, G: 102, B: 82, A: 255}
	withdrawableLine.Width = vg.Points(2)

	p.Add(queuedLine, withdrawableLine)
	p.Legend.Add("Queued", queuedLine)
	p.Legend.Add("Withdrawable", withdrawableLine)
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
