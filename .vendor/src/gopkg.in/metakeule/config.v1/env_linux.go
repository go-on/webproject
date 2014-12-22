// +build linux

// set USER_DIR, GLOBAL_DIRS and WORKING_DIR based on the XDG Base Directory Specification
// see http://standards.freedesktop.org/basedir-spec/basedir-spec-latest.html

/*
/etc/project1/
/home/user/.config/project1/
*/

package config

import (
	"os"
	"path/filepath"
	"strings"
)

func setUserDir() {
	xdg_config_home := os.Getenv("XDG_CONFIG_HOME")
	if xdg_config_home == "" {
		home := os.Getenv("HOME")
		if home == "" {
			home = filepath.Join("/home", os.Getenv("USER"))
		}
		xdg_config_home = filepath.Join(home, ".config")
	}
	USER_DIR = xdg_config_home
}

func setGlobalDir() {
	xdg_config_dirs := os.Getenv("XDG_CONFIG_DIRS")
	if xdg_config_dirs == "" {
		xdg_config_dirs = "/etc"
	}
	GLOBAL_DIRS = xdg_config_dirs
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
