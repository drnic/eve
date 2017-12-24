package cmd

import "fmt"

// ConvertOpts represents the 'convert' command
type ConvertOpts struct {
}

// Execute is callback from go-flags.Commander interface
func (c ConvertOpts) Execute(_ []string) (err error) {
	fmt.Printf("%#v\n", Opts)
	return nil
}
