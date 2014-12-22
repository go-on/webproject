package config

import (
	"encoding/json"
	"time"
)

// shortcut for MustNewOption of type bool
func (c *Config) NewBool(name, helpText string, opts ...func(*Option)) BoolGetter {
	return BoolGetter{
		opt: c.MustNewOption(name, "bool", helpText, opts),
		cfg: c,
	}
}

// shortcut for MustNewOption of type int32
func (c *Config) NewInt32(name, helpText string, opts ...func(*Option)) Int32Getter {
	return Int32Getter{
		opt: c.MustNewOption(name, "int32", helpText, opts),
		cfg: c,
	}
}

// shortcut for MustNewOption of type float32
func (c *Config) NewFloat32(name, helpText string, opts ...func(*Option)) Float32Getter {
	return Float32Getter{
		opt: c.MustNewOption(name, "float32", helpText, opts),
		cfg: c,
	}
}

// shortcut for MustNewOption of type string
func (c *Config) NewString(name, helpText string, opts ...func(*Option)) StringGetter {
	return StringGetter{
		opt: c.MustNewOption(name, "string", helpText, opts),
		cfg: c,
	}
}

// shortcut for MustNewOption of type datetime
func (c *Config) NewDateTime(name, helpText string, opts ...func(*Option)) DateTimeGetter {
	return DateTimeGetter{
		opt: c.MustNewOption(name, "datetime", helpText, opts),
		cfg: c,
	}
}

func (c *Config) NewDate(name, helpText string, opts ...func(*Option)) DateTimeGetter {
	return DateTimeGetter{
		opt: c.MustNewOption(name, "date", helpText, opts),
		cfg: c,
	}
}

func (c *Config) NewTime(name, helpText string, opts ...func(*Option)) DateTimeGetter {
	return DateTimeGetter{
		opt: c.MustNewOption(name, "time", helpText, opts),
		cfg: c,
	}
}

// shortcut for MustNewOption of type json
func (c *Config) NewJSON(name, helpText string, opts ...func(*Option)) JSONGetter {
	return JSONGetter{
		opt: c.MustNewOption(name, "json", helpText, opts),
		cfg: c,
	}
}

func Required(o *Option) { o.Required = true }

func Default(val interface{}) func(*Option) {
	return func(o *Option) { o.Default = val }
}

func Shortflag(s rune) func(*Option) {
	return func(o *Option) { o.Shortflag = string(s) }
}

// panics for invalid values
func (c *Config) MustNewOption(name, type_, helpText string, opts []func(*Option)) *Option {
	o, err := c.NewOption(name, type_, helpText, opts)
	if err != nil {
		panic(err)
	}
	return o
}

// adds a new option
func (c *Config) NewOption(name, type_, helpText string, opts []func(*Option)) (*Option, error) {
	o := &Option{Name: name, Type: type_, Help: helpText}

	for _, s := range opts {
		s(o)
	}

	if err := o.Validate(); err != nil {
		return nil, err
	}

	if err := c.addOption(o); err != nil {
		return nil, err
	}
	return o, nil
}

type Option struct {
	// Name must consist of words that are joined by the underscore character _
	// Each word must consist of uppercase letters [A-Z] and may have numbers
	// A word must consist of two ascii characters or more.
	// A name must at least have one word
	Name string `json:"name"`

	// Required indicates, if the Option is required
	Required bool `json:"required"`

	// Type must be one of "bool","int32","float32","string","datetime","json"
	Type string `json:"type"`

	// The Help string is part of the documentation
	Help string `json:"help"`

	// The Default value for the Config. The value might be nil for optional Options.
	// Otherwise, it must have the same type as the Type property indicates
	Default interface{} `json:"default,omitempty"`

	// A Shortflag for the Option. Shortflags may only be used for commandline flags
	// They must be a single lowercase ascii character
	Shortflag string `json:"shortflag,omitempty"`
}

// ValidateDefault checks if the default value is valid.
// If it does, nil is returned, otherwise
// ErrInvalidDefault is returned or a json unmarshalling error if the type is json
func (c Option) ValidateDefault() error {
	if c.Default == nil {
		return nil
	}
	err := c.ValidateValue(c.Default)
	if err != nil {
		return InvalidDefault{c.Name, c.Type, c.Default}
	}
	return nil
}

// ValidateValue checks if the given value is valid.
// If it does, nil is returned, otherwise
// ErrInvalidValue is returned or a json unmarshalling error if the type is json
func (c Option) ValidateValue(val interface{}) error {
	invalidErr := InvalidValueError{c.Name, val}
	// value may only be nil for optional Options
	if val == nil && c.Required {
		return invalidErr
	}

	if val == nil {
		return nil
	}
	switch ty := val.(type) {
	case bool:
		if c.Type != "bool" {
			return invalidErr
		}
	case int32:
		if c.Type != "int32" {
			return invalidErr
		}
	case float32:
		if c.Type != "float32" {
			return invalidErr
		}
	case string:
		if c.Type != "string" && c.Type != "json" {
			return invalidErr
		}
		if c.Type == "json" {
			var v interface{}
			if err := json.Unmarshal([]byte(ty), &v); err != nil {
				return err
			}
		}
	case time.Time:

		switch c.Type {
		case "date", "time", "datetime":
			// ok
		default:
			return invalidErr
		}

	default:
		return invalidErr
	}
	return nil
}

// Validate checks if the Option is valid.
// If it does, nil is returned, otherwise
// the error is returned
func (c Option) Validate() error {
	if err := ValidateName(c.Name); err != nil {
		return err
	}
	if err := ValidateType(c.Name, c.Type); err != nil {
		return err
	}
	if err := c.ValidateDefault(); err != nil {
		return err
	}
	if c.Help == "" {
		return ErrMissingHelp
	}
	return nil
}
