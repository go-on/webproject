package config

import (
	"os"
	"path/filepath"
)

var (
	USER_DIR    string
	GLOBAL_DIRS string // colon separated list to look for
	WORKING_DIR string
	CONFIG_EXT  = ".conf"
	ENV         []string
	ARGS        []string
)

func init() {
	ENV = os.Environ()
	ARGS = os.Args[1:]
}

// globalsFile returns the global config file path for the given dir
func (c *Config) globalsFile(dir string) string {
	return filepath.Join(dir, c.appName(), c.appName()+CONFIG_EXT)
}

// UserFile returns the user defined config file path
func (c *Config) UserFile() string {
	return filepath.Join(USER_DIR, c.appName(), c.appName()+CONFIG_EXT)
}

// LocalFile returns the local config file (inside the .config subdir of the current working dir)
func (c *Config) LocalFile() string {
	//fmt.Println(WORKING_DIR, ".config", c.appName(), c.appName()+CONFIG_EXT)
	return filepath.Join(WORKING_DIR, ".config", c.appName(), c.appName()+CONFIG_EXT)
}

// GlobalFile returns the path for the global config file in the first global directory
func (c *Config) FirstGlobalsFile() string {
	return c.globalsFile(splitGlobals()[0])
}
