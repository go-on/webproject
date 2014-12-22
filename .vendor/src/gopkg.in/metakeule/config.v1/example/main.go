package main

import (
	"fmt"
	"os"

	"gopkg.in/metakeule/config.v1"
)

var (
	cfg = config.MustNew("example", "0.0.1", "example is an example app for config")

	extra  = cfg.NewBool("extra", "extra is just a first \ntest option as a bool    ", config.Shortflag('x'), config.Default(false))
	second = cfg.NewString("second", "second is the second option and a string", config.Shortflag('s'), config.Default("2nd"))

	project = cfg.MustCommand("project", "example project sub command")

	projectName = project.NewString("name", "name of the project")
)

func main() {

	err := cfg.Run()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
	}

	if extra.Get() {
		fmt.Println("extra is true")
	} else {
		if !extra.IsSet() {
			fmt.Println("extra has not been set")
		} else {
			fmt.Println("extra is false")
		}
	}

	fmt.Printf("extra locations: %#v\n", cfg.Locations("extra"))

	fmt.Printf("second is: %#v\n", second.Get())
	fmt.Printf("second locations: %#v\n", cfg.Locations("second"))

	cmd := cfg.ActiveCommand()

	if cmd == nil {
		fmt.Println("no subcommand")
		return
	}
	switch cmd {
	case project:
		fmt.Println("project subcommand")
		fmt.Printf("project name is: %#v\n", projectName.Get())
		fmt.Printf("project locations: %#v\n", project.Locations("name"))
	default:
		panic("must not happen")
	}

	/*
		err := cfg.SaveToLocal()

		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Println("saved to ", cfg.LocalFile())
		}
	*/
}
