package cmd

import (
	"fmt"

	"github.com/starkandwayne/eve/mapping"
	"github.com/starkandwayne/eve/operator"
)

// ConvertOpts represents the 'convert' command
type ConvertOpts struct {
}

// Execute is callback from go-flags.Commander interface
func (c ConvertOpts) Execute(_ []string) (err error) {
	if Opts.Debug {
		fmt.Printf("Options: %#v\n", Opts)
	}
	set, err := mapping.NewMappingSet(Opts.Mapping)
	if err != nil {
		return
	}
	if Opts.Debug {
		fmt.Printf("Mapping Set: %#v\n", set)
	}

	inputs := mapping.NewInputsFromFlags(Opts.Inputs)

	output := operator.NewOperatorOutput(Opts.Target)
	set.GenerateOutput(inputs, output)
	if Opts.Debug {
		fmt.Printf("OperatorOutput: %#v\n", output)
	}

	fmt.Println(output)

	return nil
}
