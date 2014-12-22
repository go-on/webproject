package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func (c *Config) Load(withArgs bool) error {
	// clear old values
	c.Reset()

	// fmt.Printf("ARGS: %#v\n", ARGS)

	// first load defaults
	c.LoadDefaults()

	// then overwrite with globals, return any error
	if err := c.LoadGlobals(); err != nil {
		return err
	}

	// then overwrite with user, return any error
	if err := c.LoadUser(); err != nil {
		return err
	}

	// then overwrite with locals, return any error
	if err := c.LoadLocals(); err != nil {
		return err
	}

	// then overwrite with env, return any error
	if err := c.MergeEnv(); err != nil {
		return err
	}

	if withArgs {

		if len(ARGS) > 0 {
			// fmt.Println("we are in subcommand " + ARGS[0])
			if sub, has := c.commands[strings.ToLower(ARGS[0])]; has {
				// fmt.Println("we are in subcommand " + ARGS[0])
				c.activeCommand = sub
				if len(ARGS) == 1 {
					ARGS = []string{}
				} else {
					ARGS = ARGS[1:]
				}

				sub.LoadDefaults()

				// then overwrite with env, return any error
				if err := sub.MergeEnv(); err != nil {
					return err
				}

				merged1, err1 := c.mergeArgs(true, ARGS)
				if err1 != nil {
					return err1
				}

				// then overwrite with args
				merged2, err2 := sub.mergeArgs(true, ARGS)
				if err2 != nil {
					return err2
				}

				// fmt.Printf("merged1: %#v\nmerged2: %#v\n", merged1, merged2)

				for _, arg := range ARGS {
					key := arg
					if idx := strings.Index(arg, "="); idx != -1 {
						key = arg[:idx]
					}

					if !merged1[key] && !merged2[key] {
						return UnknownOptionError{c.version, arg}
					}
				}
				return nil

				//return sub.Load(helpIntro)
			}
		}
	}

	if withArgs {

		// then overwrite with args
		return c.MergeArgs()
	}
	return nil
}

// LoadUser loads the user specific config file
func (c *Config) LoadUser() error {
	err, found := c.LoadFile(c.UserFile())
	if found {
		return err
	}
	return nil
}

// LoadLocals merges config inside a .config subdir in the local directory
func (c *Config) LoadLocals() error {
	// fmt.Println("loading locals from " + c.LocalFile())
	err, found := c.LoadFile(c.LocalFile())
	if found {
		return err
	}
	return nil
}

// LoadGlobals loads the first config file for the app it could find inside
// the GLOBAL_DIRS and returns an error if the config could not be merged properly
// If no config file could be found, no error is returned.
func (c *Config) LoadGlobals() error {
	for _, dir := range splitGlobals() {
		err, found := c.LoadFile(filepath.Join(dir, c.appName(), c.appName()+CONFIG_EXT))
		if found {
			return err
		}
	}
	return nil
}

func (c *Config) LoadDefaults() {
	for k, spec := range c.spec {
		if spec.Default != nil {
			c.values[k] = spec.Default
			c.locations[k] = append(c.locations[k], fmt.Sprintf("%v", spec.Default))
		}
	}
}

// LoadFile merges the config from the given file and returns any error happening during the merge
// If the file could not be opened (does not exist), no error is returned
// TODO maybe an error should be returned, if the file exists, but could not be opened because
// of missing access rights
func (c *Config) LoadFile(path string) (err error, found bool) {
	//fmt.Printf("before from slash: %#v\n",path)
	path = filepath.FromSlash(path)
	file, err0 := os.Open(path)
	if err0 != nil {
		//fmt.Printf("missing file: %#v: %s\n",path, err0)
		return nil, false
	}
	found = true
	defer file.Close()
	//fmt.Printf("merging: %#v\n",path)
	err1 := c.Merge(file, path)
	if err1 != nil {
		err = fmt.Errorf("can't merge file %s: %s", file.Name(), err1.Error())
	}
	return
}

// Load loads the config values in the following order where
// each loader overwrittes corresponding config keys that have been defined
/*
	defaults
	global config
	user config
	local config
	env config
	args config
*/
// in the args config any wrong syntax or values result in writing the error to StdErr and
// exiting the program. also if --config_spec is set the spec is directly written to the
// StdOut and the program is exiting. If --help is set, the help message is printed with the
// the help  messages for the config options
func (c *Config) Run() error {
	return c.Load(true)
}
