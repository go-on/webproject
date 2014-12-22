config
======

[ ![Codeship Status for metakeule/config](https://codeship.io/projects/e39e9b00-5584-0132-0404-0e68a1ef8c1b/status)](https://codeship.io/projects/49328)

[![Build status](https://ci.appveyor.com/api/projects/status/arvt5gn2qrtgmwgl?svg=true)](https://ci.appveyor.com/project/metakeule/config)

cross plattform configuration tool.

Not ready for consumption yet.


Example
-------

```go
package main

import (
    "fmt"

    "github.com/metakeule/config"
)

var (
    cfg = config.MustNew("git", "2.1.3")
    version = cfg.NewBool("version", "prints the version")

    commit = cfg.MustSub("commit")
    
    commitCleanup = commit.NewString(
        "cleanup", 
        "This option determines how...", 
        config.Default("default"),
    )
    
    commitAll = commit.NewBool(
        "all", 
        "Tell the command to automatically...",
        config.Shortflag('a'),
    )
)

func main() {
    cfg.Run("git is a DVCS", nil)

    if version.IsSet() {
        fmt.Println("git version 2.1.3")
    }

    
    switch cfg.CurrentSub() {
    case nil:
        fmt.Println("no subcommand")
    case commit:
        fmt.Println("commit cleanup is: ", commitCleanup.Get())
        fmt.Printf("commit cleanup locations: %#v\n", commit.Locations("cleanup"))
        
        fmt.Println("commit all is: ", commitAll.Get())
        fmt.Printf("commit all locations: %#v\n", commit.Locations("all"))
    default:
        panic("must not happen")
    }
}

```