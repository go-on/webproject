package config

import "errors"

func (c *Config) SetGlobalOptions(options map[string]string) error {
	c.Reset()
	if err := c.setMap(options); err != nil {
		return err
	}
	return c.SaveToGlobals()
}

func (c *Config) SetUserOptions(options map[string]string) error {
	c.Reset()
	if err := c.setMap(options); err != nil {
		return err
	}
	return c.SaveToUser()
}

func (c *Config) SetLocalOptions(options map[string]string) error {
	c.Reset()
	if err := c.setMap(options); err != nil {
		return err
	}
	return c.SaveToLocal()
}

// SaveToGlobals saves the given config values to a global config file
// don't save secrets inside the global config, since it is readable for everyone
// A new global config is written with 0644. The config is saved inside the first
// directory of GLOBAL_DIRS
func (c *Config) SaveToGlobals() error {
	if GLOBAL_DIRS == "" {
		return errors.New("GLOBAL_DIRS not set")
	}
	return c.WriteConfigFile(c.FirstGlobalsFile(), 0644)
}

// SaveToUser saves all values to the user config file
// creating missing directories
// A new config is written with 0640, ro readable for user group and writeable for the user
func (c *Config) SaveToUser() error {
	if USER_DIR == "" {
		return errors.New("USER_DIR not set")
	}
	return c.WriteConfigFile(c.UserFile(), 0640)
}

// SaveToLocal saves all values to the local config file
// A new config is written with 0640, ro readable for user group and writeable for the user
func (c *Config) SaveToLocal() error {
	if WORKING_DIR == "" {
		return errors.New("WORKING_DIR not set")
	}
	return c.WriteConfigFile(c.LocalFile(), 0640)
}
