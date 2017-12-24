package main

import (
	"fmt"
	"os"

	"github.com/jessevdk/go-flags"
	evecmd "github.com/starkandwayne/eve/cmd"
)

var Version = ""

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-v" || os.Args[1] == "--version" {
			if Version == "" {
				fmt.Printf("eve (development)\n")
			} else {
				fmt.Printf("eve v%s\n", Version)
			}
			os.Exit(0)
		}
	}

	parser := flags.NewParser(&evecmd.Opts, flags.Default)

	if len(os.Args) == 1 {
		_, err := parser.ParseArgs([]string{"--help"})
		if err != nil {
			os.Exit(1)
		}
	} else {
		_, err := parser.Parse()
		if err != nil {
			os.Exit(1)
		}
	}
}
