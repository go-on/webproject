package pgsql

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

type RowValidator interface {
	ValidateRow(value map[*Field]interface{}) error
}

type FieldValidator interface {
	Validate(value interface{}) error
}

type TypeValidator struct{ *Field }

func (ø *TypeValidator) Validate(value interface{}) error {

	/*
		// don#t check it here, let the database do the check.
		// this way we don't get false positives on updates that do not change values
		if value == nil {
			if !ø.Field.Is(NullAllowed) && ø.Field.Default == nil {
				return fmt.Errorf("nil (null) is not allowed")

			}
			return nil
		}
	*/
	valString := ToString(value)
	switch ø.Field.Type {
	case IntType:
		_, err := strconv.ParseInt(valString, 10, 32)
		if err != nil {
			return fmt.Errorf("%#v is no Int: %s", value, err.Error())
		}
	case FloatType:
		_, err := strconv.ParseFloat(valString, 32)
		if err != nil {
			return fmt.Errorf("%#v is no Float: %s", value, err.Error())
		}
	default:
		if IsVarChar(ø.Field.Type) {
			if no := utf8.RuneCountInString(valString); no > int(ø.Field.Type) {
				return fmt.Errorf("%#v with length %v is too long for a varchar(%v)", value, no, int(ø.Field.Type))
			}
		}
	}
	return nil
}

type SelectionValidator struct{ *Field }

func (ø *SelectionValidator) Validate(value interface{}) error {
	if value == nil {
		return nil
	}
	if !ø.Field.InSelection(value) {
		return fmt.Errorf("%#v is not in the selection: %#v", value, ø.Field.Selection)
	}
	return nil
}

type OrFieldValidator []FieldValidator

func (ø OrFieldValidator) Validate(value interface{}) (err error) {
	for _, v := range ø {
		if err = v.Validate(value); err == nil {
			return nil
		}
	}
	return
}

type OrRowValidator []RowValidator

func (ø OrRowValidator) ValidateRow(value map[*Field]interface{}) (err error) {
	for _, v := range ø {
		if err = v.ValidateRow(value); err == nil {
			return nil
		}
	}
	return
}
