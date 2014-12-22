// +build !linux,!windows,!darwin

package config

// environment for unixy system that are not linux and not darwin, like the BSD family

import (
	"os"
	"path/filepath"
	"strings"
)

func setUserDir() {
	home := os.Getenv("HOME")
	if home == "" {
		home = filepath.Join("/home", os.Getenv("USER"))
	}
	USER_DIR = filepath.Join(home + ".config")
}

func setGlobalDir() {
	GLOBAL_DIRS = "/usr/local/etc"
}

func setWorkingDir() {
	wd, err := os.Getwd()
	if err != nil {
		wd = "."
	}

	WORKING_DIR = wd
}

func splitGlobals() []string {
	return strings.Split(GLOBAL_DIRS, ":")
}

func init() {
	setUserDir()
	setGlobalDir()
	setWorkingDir()
}
