package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	// "flag"
	// "fmt"
	// "os"
	"os/exec"

	"gopkg.in/metakeule/config.v1"
)

var (
	cfg               = config.MustNew("config", "1.10.0", "a multiplattform and multilanguage configuration tool")
	optionProgram     = cfg.NewString("program", "the program where the options belong to (must be a config compatible program)", config.Required, config.Shortflag('p'))
	optionLocations   = cfg.NewBool("locations", "the locations where the options are currently set", config.Shortflag('l'))
	cfgSet            = cfg.MustCommand("set", "set an option").Skip("locations")
	optionSetKey      = cfgSet.NewString("option", "the option that should be set", config.Required, config.Shortflag('o'))
	optionSetValue    = cfgSet.NewString("value", "the value the option should be set to", config.Required, config.Shortflag('v'))
	optionSetPathType = cfgSet.NewString("type", "the type of the config path where the value should be set. valid values are global,user and local", config.Shortflag('t'), config.Required)
	cfgGet            = cfg.MustCommand("get", "get the current value of an option").Skip("locations")
	optionGetKey      = cfgGet.NewString("option", "the option that should be get, if not set, all options that are set are returned", config.Shortflag('o'))
	cfgPath           = cfg.MustCommand("path", "show the paths for the configuration files").Skip("locations")
	optionPathType    = cfgPath.NewString("type", "the type of the config path. valid values are global,user,local and all", config.Shortflag('t'), config.Default("all"))
)

func GetVersion(cmdpath string) (string, error) {
	cmd := exec.Command(cmdpath, "--version")
	out, err := cmd.Output()
	if err != nil {
		return "", err
	}
	// fmt.Printf("version: %#v\n", string(out))
	v := strings.Split(strings.TrimSpace(string(out)), " ")
	if len(v) != 3 {
		return "", fmt.Errorf("%s --version returns invalid result: %#v", cmdpath, string(out))
	}
	return strings.TrimSpace(v[2]), nil
}

func GetSpec(cmdpath string, c *config.Config) error {
	cmd := exec.Command(cmdpath, "--config-spec")
	out, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("%s does not seem to be compatible with config", cmdpath)
		// return err
	}
	return c.UnmarshalJSON(out)
}

func writeErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		fmt.Fprintln(os.Stdout, " -> run 'config help' to get more help")
		os.Exit(1)
	}
}

func cmdPrintAll() error {
	if optionLocations.IsSet() {
		err := cmdConfig.Load(false)
		if err != nil {
			return fmt.Errorf("Can't load options for command %s: %s", cmd, err.Error())
			// os.Exit(1)
		}
		locations := map[string][]string{}

		cmdConfig.EachValue(func(name string, value interface{}) {
			locations[name] = cmdConfig.Locations(name)
		})

		var b []byte
		b, err = json.Marshal(locations)
		if err != nil {
			return fmt.Errorf("Can't print locations for command %s: %s", cmd, err.Error())
			// os.Exit(1)
		}

		fmt.Fprintln(os.Stdout, string(b))
		// os.Exit(0)
	}
	return nil
}

var cmdConfig *config.Config
var commandPath string
var cmd string

func main() {

	err := cfg.Run()
	writeErr(err)
	cmd = optionProgram.Get()
	commandPath, err = exec.LookPath(cmd)
	writeErr(err)
	var version string
	version, err = GetVersion(commandPath)
	writeErr(err)

	cmdConfig, err = config.New(filepath.Base(cmd), version, "")
	writeErr(err)
	err = GetSpec(commandPath, cmdConfig)
	writeErr(err)

	command := cfg.ActiveCommand()

	if command == nil {
		cmdPrintAll()
		return
	}

	switch command {

	// fmt.Println("no subcommand")
	case cfgGet:
		err := cmdConfig.Load(false)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Can't load config options for program %s: %s", cmd, err.Error())
			os.Exit(1)
		}
		if !optionGetKey.IsSet() {
			var vals = map[string]interface{}{}
			cmdConfig.EachValue(func(name string, value interface{}) {
				vals[name] = value
			})
			var b []byte
			b, err = json.Marshal(vals)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Can't print locations for program %s: %s", cmd, err.Error())
				os.Exit(1)
			}

			fmt.Fprintln(os.Stdout, string(b))
			os.Exit(0)
		} else {
			key := optionGetKey.Get()
			if !cmdConfig.IsOption(key) {
				fmt.Fprintf(os.Stderr, "unknown option %s", key)
				os.Exit(1)
			}

			val := cmdConfig.GetValue(key)
			// cmdConfig.
			fmt.Fprintf(os.Stdout, "%v\n", val)
		}

	case cfgSet:
		key := optionSetKey.Get()
		val := optionSetValue.Get()
		ty := optionSetPathType.Get()
		switch ty {
		case "user":
			if err := cmdConfig.LoadUser(); err != nil {
				fmt.Fprintf(os.Stderr, "Can't load user config file: %s", err.Error())
				os.Exit(1)
			}
			if err := cmdConfig.Set(key, val, cmdConfig.UserFile()); err != nil {
				fmt.Fprintf(os.Stderr, "Can't set option %#v to value %#v: %s", key, val, err.Error())
				os.Exit(1)
			}
			if err := cmdConfig.SaveToUser(); err != nil {
				fmt.Fprintf(os.Stderr, "Can't save user config file: %s", err.Error())
				os.Exit(1)
			}
		case "local":
			if err := cmdConfig.LoadLocals(); err != nil {
				fmt.Fprintf(os.Stderr, "Can't load local config file: %s", err.Error())
				os.Exit(1)
			}
			if err := cmdConfig.Set(key, val, cmdConfig.LocalFile()); err != nil {
				fmt.Fprintf(os.Stderr, "Can't set option %#v to value %#v: %s", key, val, err.Error())
				os.Exit(1)
			}
			if err := cmdConfig.SaveToLocal(); err != nil {
				fmt.Fprintf(os.Stderr, "Can't save local config file: %s", err.Error())
				os.Exit(1)
			}
		case "global":
			if err := cmdConfig.LoadGlobals(); err != nil {
				fmt.Fprintf(os.Stderr, "Can't load global config file: %s", err.Error())
				os.Exit(1)
			}
			if err := cmdConfig.Set(key, val, cmdConfig.FirstGlobalsFile()); err != nil {
				fmt.Fprintf(os.Stderr, "Can't set option %#v to value %#v: %s", key, val, err.Error())
				os.Exit(1)
			}
			if err := cmdConfig.SaveToGlobals(); err != nil {
				fmt.Fprintf(os.Stderr, "Can't save global config file: %s", err.Error())
				os.Exit(1)
			}
		default:
			fmt.Fprintf(os.Stderr, "'%s' is not a valid value for type option. possible values are 'local', 'global' or 'user'", ty)
			os.Exit(1)

		}
	case cfgPath:
		ty := optionPathType.Get()
		switch ty {
		case "user":
			fmt.Fprintln(os.Stdout, cmdConfig.UserFile())
			os.Exit(0)
		case "local":
			fmt.Fprintln(os.Stdout, cmdConfig.LocalFile())
			os.Exit(0)
		case "global":
			fmt.Fprintln(os.Stdout, cmdConfig.FirstGlobalsFile())
			os.Exit(0)
		case "all":
			paths := map[string]string{
				"user":   cmdConfig.UserFile(),
				"local":  cmdConfig.LocalFile(),
				"global": cmdConfig.FirstGlobalsFile(),
			}
			b, err := json.Marshal(paths)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Can't print locations for program %s: %s", cmd, err.Error())
				os.Exit(1)
			}

			fmt.Fprintln(os.Stdout, string(b))
			os.Exit(0)
		default:
			fmt.Fprintf(os.Stderr, "'%s' is not a valid value for type option. possible values are 'local', 'global' or 'user'", ty)
			os.Exit(1)
		}
	// some not allowed subcommand, should already be catched by config.Run
	default:
		panic("must not happen")

	}

}

/*
tool to read and set configurations

keys consist of names that are all uppercase letters separated by underscore _

config [binary] key

returns type: value

supported types are:
bool, int32, float32, string (utf-8), datetime, json

(this reads config)

config [binary] -l key1=value1,key2=value2 // sets the options in the local config file (relative to dir)
config [binary] -u key1=value1,key2=value2 // sets the options in the user config file
config [binary] -g key1=value1,key2=value2 // sets the options in the global config file
config [binary] -c key1=value1,key2=value2 // checks the options for the binary
config [binary] -h key                     // prints help about the key
config [binary] -h                         // prints help about all options
config [binary] -m key1=value1,key2=value2 // merges the options with global/user/local ones and prints the result

each setting of an option is checked for validity of the type.
for json values it is only checked, if it is valid json. additional
checks for the json structure must be done by the binary

values are passed the following way:
boolean values: true|false
int32 values: 34523
float32 values: 4.567
string values: "here the utf-8 string"
datetime values: 2006-01-02T15:04:05Z07:00    (RFC3339)
json values: '{"a": "\'b\'"}'

a binary that is supported by config is supposed to be callable with --config-spec and then return a json encoded hash of the options in the form of
[
	{
		"key": "here_the_key1",
	  "required": true|false,
	  "type": "bool"|"int32"|"float32"|"string"|"datetime"|"json",
	  "description": "...",
	  "default": null|"value"
	},
  {
  	"key": "here_the_key2",
	  required: true|false,
	  type: "bool"|"int32"|"float32"|"string"|"datetime"|"json",
	  description: "...",
	  "default": null|"value"
	}
	[...]
]

config is meant to be used on the plattforms:
- linux
- windows
- mac os x
(maybe Android too)

it builds the configuration options by starting with defaults and merging in the following configurations
(the next overwriting the same previously defined key):

defaults as reported via [binary] --config-spec
plattform-specific global config
plattform-specific user config
local config in directory named .config/[binary] in the current working directory
environmental variables starting with [BINARY]_CONFIG_
given args to the commandline

the binary itself wants to get all options in a single go.
it therefore may run

  config [binary] -args argstring

additionally there is a library for go (and might be created for other languages)
that make it easy to query the final options in a type-safe manner

subcommands are handled as if they were extra binaries with the name
[binary]_[subcommand]: they have separate config files and if a binary name with an  underscore
is passed to config the part after the underscore is considered a subcommand.
The environment variables for subcommands start with [BINARY]_[SUBCOMMAND]_CONFIG_
when a subcommand is called the options/configuration for the binary are also loaded.


*/

/*
func main() {
	flag.Parse()
	fmt.Printf("%#v\n", os.Args)
	fmt.Printf("%#v\n", flag.Args())
}
*/
