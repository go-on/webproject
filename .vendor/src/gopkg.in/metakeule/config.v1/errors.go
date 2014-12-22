package config

import (
	"errors"
	"fmt"
)

var (
	// ErrInvalidName      = errors.New("invalid name")
	ErrInvalidVersion   = errors.New("invalid version")
	ErrInvalidShortflag = errors.New("invalid shortflag")
	ErrCommandCommand   = errors.New("command of command is not supported")

	//ErrInvalidDefault = errors.New("invalid default")
	// ErrInvalidValue   = errors.New("invalid value")
	ErrMissingHelp = errors.New("missing help text")
)

type EmptyValueError string

func (e EmptyValueError) Error() string {
	return fmt.Sprintf("invalid value: empty string for %#v", string(e))
}

type InvalidNameError string

func (e InvalidNameError) Error() string {
	return fmt.Sprintf("invalid name %#v", string(e))
}

type InvalidTypeError struct {
	Option string
	Type   string
}

func (e InvalidTypeError) Error() string {
	return fmt.Sprintf("invalid type %#v for option %#v", e.Type, e.Option)
}

type InvalidDefault struct {
	Option  string
	Type    string
	Default interface{}
}

func (e InvalidDefault) Error() string {
	return fmt.Sprintf("invalid default value %#v option %s of type %s", e.Default, e.Option, e.Type)
}

type MissingOptionError struct {
	Version string
	Option  string
}

func (e MissingOptionError) Error() string {
	return fmt.Sprintf("missing option %s is not allowed in version %s", e.Option, e.Version)
}

type InvalidConfigEnv struct {
	Version string
	EnvKey  string
	Err     error
}

func (e InvalidConfigEnv) Error() string {
	return fmt.Sprintf("env variable %s is not compatible with version %s: %s", e.EnvKey, e.Version, e.Err.Error())
}

type InvalidConfigFlag struct {
	Version string
	Flag    string
	Err     error
}

func (e InvalidConfigFlag) Error() string {
	return fmt.Sprintf("flag %s is not compatible with version %s: %s", e.Flag, e.Version, e.Err.Error())
}

type InvalidConfig struct {
	Version string
	Err     error
}

func (e InvalidConfig) Error() string {
	return fmt.Sprintf("config is not compatible with version %s: %s", e.Version, e.Err.Error())
}

type InvalidConfigFileError struct {
	ConfigFile string
	Version    string
	Err        error
}

func (e InvalidConfigFileError) Error() string {
	return fmt.Sprintf("config file %s is not compatible with version %s: %s", e.ConfigFile, e.Version, e.Err.Error())
}

type InvalidValueError struct {
	Option string
	Value  interface{}
}

func (e InvalidValueError) Error() string {
	return fmt.Sprintf("value %#v is invalid for option %s", e.Value, e.Option)
}

type ErrInvalidOptionName string

func (e ErrInvalidOptionName) Error() string {
	return fmt.Sprintf("invalid option name %s", string(e))
}

type ErrInvalidAppName string

func (e ErrInvalidAppName) Error() string {
	return fmt.Sprintf("invalid app name %s", string(e))
}

type UnknownOptionError struct {
	Version string
	Option  string
}

func (e UnknownOptionError) Error() string {
	return fmt.Sprintf("option %s is unknown in version %s", e.Option, e.Version)
}

type ErrDoubleOption string

func (e ErrDoubleOption) Error() string {
	return fmt.Sprintf("option %s is set twice", string(e))
}

type ErrDoubleShortflag string

func (e ErrDoubleShortflag) Error() string {
	return fmt.Sprintf("shortflag %s is set twice", string(e))
}
