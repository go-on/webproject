package pgsql

import "fmt"

type ValidationError struct {
	Field   *Field
	Table   *Table
	Err     error
	Details error // error details, not shown to the user
}

func (ve *ValidationError) Error() string {
	if ve.Field != nil {
		return fmt.Sprintf(
			"validation error for field %s in table %s: %s",
			ve.Field.Name,
			ve.Table.Name,
			ve.Err.Error(),
		)
	}

	if ve.Table != nil {
		return fmt.Sprintf(
			"validation error for table %s: %s",
			ve.Table.Name,
			ve.Err.Error(),
		)
	}

	return fmt.Sprintf(
		"validation error: %s",
		ve.Err.Error(),
	)
}

func convertError(field *Field, val interface{}, detailedError error) *ValidationError {
	return &ValidationError{
		Field:   field,
		Table:   field.Table,
		Err:     fmt.Errorf("can't convert %#v to %s", val, field.Type.String()),
		Details: detailedError,
	}
}

func nullNotAllowedError(field *Field, val interface{}) *ValidationError {
	return &ValidationError{
		Field: field,
		Table: field.Table,
		Err:   fmt.Errorf("can't set to value %#v: Null is not allowed \n", val),
	}
}

func fieldError(field *Field, err error) *ValidationError {
	return &ValidationError{
		Field:   field,
		Table:   field.Table,
		Err:     err,
		Details: err,
	}
}

func aliasConvertError(as *AsStruct, val interface{}, detailedError error) *ValidationError {
	return &ValidationError{
		Err:     fmt.Errorf("can't convert %#v to %s for alias %s", val, as.Type.String(), as.As),
		Details: detailedError,
	}
}
