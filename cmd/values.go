package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/starkandwayne/eve/mapping"
	"github.com/starkandwayne/eve/operator"
)

// ValuesOpts represents the 'values' command
type ValuesOpts struct {
	YAML bool `long:"yaml" description:"Return values in YAML, defaults to JSON"`
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

	ops := operator.NewOperatorOutput(Opts.Target)
	values, err := set.LoadValues(ops)
	if err != nil {
		return
	}
	var data []byte
	if c.YAML {
		data, err = yaml.Marshal(values.ValuesByName)
		if err != nil {
			return
		}
	} else {
		data, err = json.Marshal(values.ValuesByName)
		if err != nil {
			return
		}
	}
	fmt.Println(string(data[:]))

	return nil
}
