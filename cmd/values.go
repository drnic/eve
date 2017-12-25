package cmd

import (
	"fmt"
	"os"

	"github.com/starkandwayne/eve/mapping"
)

// ValuesOpts represents the 'values' command
type ValuesOpts struct {
}

// Execute is callback from go-flags.Commander interface
func (c ValuesOpts) Execute(_ []string) (err error) {
	if Opts.Debug {
		fmt.Fprintf(os.Stderr, "Options: %#v\n", Opts)
	}
	set, err := mapping.NewMappingSet(Opts.Mapping)
	if err != nil {
		return
	}
	if Opts.Debug {
		fmt.Fprintf(os.Stderr, "Mapping Set: %#v\n", set)
	}

	if Opts.Target == "" {
		fmt.Println("{}")
	} else {
		fmt.Println("{}")
	}

	return nil
}
