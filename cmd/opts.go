package cmd

// BrokerOpts describes subset of flags/options for selecting target service broker API
type BrokerOpts struct {
}

// EveOpts describes the flags/options for the CLI
type EveOpts struct {
	Version bool              `short:"v" long:"version" description:"Show version"`
	Debug   bool              `long:"debug" description:"Show some debug info along the way"`
	Inputs  map[string]string `short:"i" long:"inputs" description:"Mapping of form inputs to values"`
	Mapping string            `short:"m" long:"mapping" description:"Path to file describing mapping of form inputs to operator file result"`
	Target  string            `short:"t" long:"target" description:"Path to Operator file to create"`

	Convert ConvertOpts `command:"convert" description:"Convert form values to Operator file"`
}

// Opts carries all the user provided options (from flags or env vars)
var Opts EveOpts
