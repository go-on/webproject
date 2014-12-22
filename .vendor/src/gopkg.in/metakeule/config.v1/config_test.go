package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func withTempConfig(fn func()) error {
	dir, err := ioutil.TempDir(os.TempDir(), "config_test")
	// fmt.Println(dir)
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	subdirs := [...]string{"user", "local", "global"}

	for _, subdir := range subdirs {
		err = os.Mkdir(filepath.Join(dir, subdir), 0755)
		if err != nil {
			return err
		}
	}

	USER_DIR = filepath.Join(dir, "user")
	GLOBAL_DIRS = filepath.Join(dir, "global")
	WORKING_DIR = filepath.Join(dir, "local")

	fn()
	return nil
}

func init() {
	CONFIG_EXT = ".tmp"
}

func TestConfig(t *testing.T) {
	tests := [...]struct {
		Option    string
		Help      string
		Type      string
		Default   interface{}
		Required  bool
		Shortflag rune
		ENV       string
		Global    string
		Local     string
		User      string
		Arg       string
		expected  interface{}
	}{
		{
			Option:   "name",
			Type:     "string",
			Help:     "Test default",
			Default:  "Donald",
			expected: "Donald",
		},
		{
			Option:   "name",
			Type:     "string",
			Help:     "Test global override",
			Default:  "Donald",
			Global:   "Daisy",
			expected: "Daisy",
		},
		{
			Option:   "name",
			Type:     "string",
			Help:     "Test user override",
			Default:  "Donald",
			Global:   "Daisy",
			User:     "Mickey",
			expected: "Mickey",
		},
		{
			Option:   "name",
			Type:     "string",
			Help:     "Test local override",
			Default:  "Donald",
			Global:   "Daisy",
			User:     "Mickey",
			Local:    "Minnie",
			expected: "Minnie",
		},

		{
			Option:   "name",
			Type:     "string",
			Help:     "Test env override",
			Default:  "Donald",
			Global:   "Daisy",
			User:     "Mickey",
			Local:    "Minnie",
			ENV:      "Batman",
			expected: "Batman",
		},

		{
			Option:   "name",
			Type:     "string",
			Help:     "Test args override",
			Default:  "Donald",
			Global:   "Daisy",
			User:     "Mickey",
			Local:    "Minnie",
			ENV:      "Batman",
			Arg:      "Superman",
			expected: "Superman",
		},
		{
			Option:   "age",
			Type:     "int32",
			Help:     "Test int32",
			Default:  int32(2),
			Local:    "45",
			expected: int32(45),
		},
		{
			Option:   "height",
			Type:     "float32",
			Help:     "Test float32",
			Default:  float32(1.85),
			Local:    "1.65",
			expected: float32(1.65),
		},
		{
			Option:   "male",
			Type:     "bool",
			Help:     "Test bool",
			Default:  true,
			Local:    "false",
			expected: false,
		},
		{
			Option:   "xmas",
			Type:     "date",
			Help:     "Test date",
			Local:    "2014-12-24",
			expected: time.Date(2014, 12, 24, 0, 0, 0, 0, time.UTC),
		},
		{
			Option:   "noon",
			Type:     "time",
			Help:     "Test time",
			Local:    "12:00:00",
			expected: time.Date(0, 1, 1, 12, 0, 0, 0, time.UTC),
		},
		{
			Option:   "xmasnoon",
			Type:     "datetime",
			Help:     "Test datetime",
			Local:    "2014-12-24 12:00:00",
			expected: time.Date(2014, 12, 24, 12, 0, 0, 0, time.UTC),
		},
		{
			Option:   "friends",
			Type:     "json",
			Help:     "Test json",
			Local:    `["Ben", "Bob", "Jil"]`,
			expected: `["Ben", "Bob", "Jil"]`,
		},
	}

	err := withTempConfig(func() {
		for _, test := range tests {
			cfg, er := New("testapp", "0.1", "a testapp")
			if er != nil {
				t.Errorf(er.Error())
			}
			setters := []func(*Option){}

			if test.Default != nil {
				setters = append(setters, Default(test.Default))
			}

			if test.Required {
				setters = append(setters, Required)
			}

			if string(test.Shortflag) != "" {
				setters = append(setters, Shortflag(test.Shortflag))
			}

			cfg.MustNewOption(test.Option, test.Type, test.Help, setters)
			cfg.Reset()

			if test.Global != "" {
				if err := cfg.Set(test.Option, test.Global, GLOBAL_DIRS); err != nil {
					t.Fatal(err)
				}
			}

			if err := cfg.SaveToGlobals(); err != nil {
				t.Fatal(err)
			}

			cfg.Reset()

			if test.User != "" {
				if err := cfg.Set(test.Option, test.User, USER_DIR); err != nil {
					t.Fatal(err)
				}
			}

			if err := cfg.SaveToUser(); err != nil {
				t.Fatal(err)
			}

			cfg.Reset()

			if test.Local != "" {
				if err := cfg.Set(test.Option, test.Local, WORKING_DIR); err != nil {
					t.Fatal(err)
				}
			}

			if err := cfg.SaveToLocal(); err != nil {
				t.Fatal(err)
			}

			cfg.Reset()

			if test.ENV != "" {
				ENV = []string{"TESTAPP_CONFIG_" + strings.ToUpper(test.Option) + "=" + test.ENV}
			} else {
				ENV = []string{}
			}

			if test.Arg != "" {
				ARGS = []string{"--" + test.Option + "=" + test.Arg}
			} else {
				ARGS = []string{}
			}

			if err := cfg.Load(true); err != nil {
				t.Fatal(err)
			}

			got := cfg.values[test.Option]

			if got != test.expected {
				t.Errorf("%#v faied: test[%s] = %s, expected %s", test.Help, test.Option, got, test.expected)
			}
		}
	})

	if err != nil {
		t.Fatal(err)
	}

}
