package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	DateFormat     = "2006-01-02"
	TimeFormat     = "15:04:05"
	DateTimeFormat = "2006-01-02 15:04:05"
)

var (
	NameRegExp      = regexp.MustCompile("^[a-z][a-z0-9]+$")
	VersionRegexp   = regexp.MustCompile("^[a-z0-9-.]+$")
	ShortflagRegexp = regexp.MustCompile("^[a-z]$")
)

func ValidateShortflag(shortflag string) error {
	if shortflag == "" || ShortflagRegexp.MatchString(shortflag) {
		return nil
	}
	return ErrInvalidShortflag
}

// ValidateName checks if the given name conforms to the
// naming convention. If it does, nil is returned, otherwise
// ErrInvalidName is returned
func ValidateName(name string) error {
	if name == "" {
		return InvalidNameError(name)
	}

	if !NameRegExp.MatchString(name) {
		return InvalidNameError(name)
	}

	return nil
}

func ValidateVersion(version string) error {
	if !VersionRegexp.MatchString(version) {
		return ErrInvalidVersion
	}
	return nil
}

// ValidateType checks if the given type is valid.
// If it does, nil is returned, otherwise
// ErrInvalidType is returned
func ValidateType(option, typ string) error {
	switch typ {
	case "bool", "int32", "float32", "string", "datetime", "date", "time", "json":
		return nil
	default:
		return InvalidTypeError{option, typ}
	}
}

//var delim = []byte("\u220e\n")
var delim = []byte("\n$")

// var delim = []byte("\n\n")

func stringToValue(typ string, in string) (out interface{}, err error) {
	switch typ {
	case "bool":
		return strconv.ParseBool(in)
	case "int32":
		i, e := strconv.ParseInt(in, 10, 32)
		return int32(i), e
	case "float32":
		fl, e := strconv.ParseFloat(in, 32)
		return float32(fl), e
	case "datetime":
		return time.Parse(DateTimeFormat, in)
	case "date":
		return time.Parse(DateFormat, in)
	case "time":
		return time.Parse(TimeFormat, in)
	case "string":
		return in, nil
	case "json":
		var v interface{}
		err = json.Unmarshal([]byte(in), &v)
		if err != nil {
			return nil, err
		}
		return in, nil
	default:
		return nil, errors.New("unknown type " + typ)
	}

}

func keyToArg(key string) string {
	return "--" + key
}

func argToKey(arg string) string {
	return strings.TrimLeft(arg, "-")
}

func err2Stderr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
