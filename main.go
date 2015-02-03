// execif v0.0.1
//
// (c) Ground Six 2015
//
// @package execif
// @version 0.1.1
//
// @author Harry Lawrence <http://github.com/hazbo>
//
// License: MIT
//
// For the full copyright and license information, please view the LICENSE
// file that was distributed with this source code.

package main

import (
    "os"
    "strings"
    "github.com/asink/cli"
    "github.com/asink/libasink"
    "github.com/asink/inotify"
)

func main() {
    app := cli.NewApp()

    app.Name    = appName
    app.Version = version
    app.Usage   = usage
    app.Author  = author
    app.Email   = email

    app.Commands = []cli.Command{
        {
            Name: "run",
            Usage: "[file or dir to check]",
            SkipFlagParsing: true,
            Action: func(c *cli.Context) {
                initExecif()
            },
        },
    }
    app.Run(os.Args)
}

// Start the execif program by creating a command and
// starting the watcher.
func initExecif() {
    command  := ""
    args     := []string{}
    location := os.Args[2]

    for i := 3; i < len(os.Args); i++ {
        if i == 3 {
            command = os.Args[i]
        } else {
            if os.Args[i] != "" {
                args = append(args, string(os.Args[i]))
            }
        }
    }
    asinkcommand := asink.NewCommand(command)
    asinkcommand.Args = args

    startWatcher(getFileToWaitFor(location), location, asinkcommand)
}

// Starts the watcher on the specified directory
// then execs the asink command once the new
// file or directory is found.
func startWatcher(directory string, location string, command asink.Command) {
    watcher, err := inotify.NewWatcher()
    if err != nil {
        panic(err)
    }
    err = watcher.Watch(directory)
    if err != nil {
        panic(err)
    }
    for {
        select {
        case ev := <-watcher.Event:
            if ev.Name == location {
                command.Exec()
                os.Exit(0)
            }
        case err := <-watcher.Error:
            println("error:", err)
        }
    }
}

// Gets the file or directory the user will be waiting
// for, based on the first arg passed to execif
func getFileToWaitFor(location string) string {
    parts := strings.Split(location, "/")
    wait  := ""
    for _,v := range parts {
        wait = v
    }
    return strings.Replace(location, "/" + wait, "", -1)
}
