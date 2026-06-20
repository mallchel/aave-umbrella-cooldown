package backend

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"1-task/internal/storage/postgres"

	"github.com/go-playground/validator/v10"
)

const (
	withdrawFlowEventTypeRequest  = "request"
	withdrawFlowEventTypeWithdraw = "withdraw"
)

// listWithdrawFlowsFilter is the HTTP-layer filter model for listing withdraw flows.
type listWithdrawFlowsFilter struct {
	SenderAddress *string
	EventType     *string
	FromTime      *time.Time
	ToTime        *time.Time
	Limit         int
	Offset        int
}

type listWithdrawFlowsParamsInput struct {
	SenderAddress *string    `validate:"omitempty,min=1"`
	EventType     *string    `validate:"omitempty,oneof=request withdraw"`
	FromTime      *time.Time `validate:"omitempty,lteoptionalfield=ToTime"`
	ToTime        *time.Time `validate:"omitempty"`
	Limit         *int       `validate:"omitempty,min=1"`
	Offset        *int       `validate:"omitempty,min=0"`
}

var withdrawFlowsParamsValidator = newWithdrawFlowsParamsValidator()

func newWithdrawFlowsParamsValidator() *validator.Validate {
	validate := validator.New()
	if err := validate.RegisterValidation("lteoptionalfield", validateTimeLTEOptionalField); err != nil {
		panic(err)
	}
	return validate
}

func validateTimeLTEOptionalField(fieldLevel validator.FieldLevel) bool {
	field := fieldLevel.Field()
	otherField := fieldLevel.Parent().FieldByName(fieldLevel.Param())
	if isNilValue(field) || isNilValue(otherField) {
		return true
	}

	fieldTime, ok := field.Interface().(*time.Time)
	if !ok {
		return false
	}
	otherTime, ok := otherField.Interface().(*time.Time)
	if !ok {
		return false
	}

	return !fieldTime.After(*otherTime)
}

func isNilValue(value reflect.Value) bool {
	return !value.IsValid() || (value.Kind() == reflect.Pointer && value.IsNil())
}

func parseListWithdrawFlowsParams(params ListWithdrawFlowsParams) (listWithdrawFlowsFilter, error) {
	input := listWithdrawFlowsParamsInput{
		FromTime: params.FromTime,
		ToTime:   params.ToTime,
		Limit:    params.Limit,
		Offset:   params.Offset,
	}

	if params.SenderAddress != nil {
		sender := strings.TrimSpace(*params.SenderAddress)
		input.SenderAddress = &sender
	}

	if params.EventType != nil {
		eventType := strings.ToLower(strings.TrimSpace(string(*params.EventType)))
		input.EventType = &eventType
	}

	if err := validateListWithdrawFlowsParams(input); err != nil {
		return listWithdrawFlowsFilter{}, err
	}

	filter := listWithdrawFlowsFilter{
		SenderAddress: input.SenderAddress,
		EventType:     input.EventType,
		FromTime:      input.FromTime,
		ToTime:        input.ToTime,
	}

	limit := 100
	if input.Limit != nil {
		limit = min(*input.Limit, 500)
	}

	offset := 0
	if input.Offset != nil {
		offset = *input.Offset
	}

	filter.Limit = limit
	filter.Offset = offset

	return filter, nil
}

func validateListWithdrawFlowsParams(input listWithdrawFlowsParamsInput) error {
	if err := withdrawFlowsParamsValidator.Struct(input); err != nil {
		var validationErrors validator.ValidationErrors
		if errors.As(err, &validationErrors) && len(validationErrors) > 0 {
			return badParam(listWithdrawFlowsParamValidationMessage(validationErrors[0]))
		}
		return err
	}

	return nil
}

func listWithdrawFlowsParamValidationMessage(fieldError validator.FieldError) string {
	switch fieldError.Field() {
	case "SenderAddress":
		return "sender_address must not be empty"
	case "EventType":
		return fmt.Sprintf("event_type must be one of: %s, %s", withdrawFlowEventTypeRequest, withdrawFlowEventTypeWithdraw)
	case "FromTime":
		return "from_time must be before or equal to to_time"
	case "Limit":
		return "limit must be a positive integer"
	case "Offset":
		return "offset must be a non-negative integer"
	default:
		return fmt.Sprintf("%s is invalid", fieldError.Field())
	}
}

func (f listWithdrawFlowsFilter) toStorageFilter() postgres.ListWithdrawFlowsFilter {
	return postgres.ListWithdrawFlowsFilter{
		SenderAddress: f.SenderAddress,
		EventType:     f.EventType,
		FromTime:      f.FromTime,
		ToTime:        f.ToTime,
		Limit:         f.Limit,
		Offset:        f.Offset,
	}
}
