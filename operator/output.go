package operator

import "gopkg.in/yaml.v2"

// Output describes the output file containing a set of Operators
type Output struct {
	TargetPath string
	Operators  OpDefinitions
}

type OpDefinitions []OpDefinition

// OpDefinition struct is useful for YAML unmarshaling
type OpDefinition struct {
	Type  string      `yaml:",omitempty"`
	Path  string      `yaml:",omitempty"`
	Value interface{} `yaml:",omitempty"`
}

// NewOperatorOutput is a constructor
func NewOperatorOutput(targetPath string) (out *Output) {
	return &Output{
		TargetPath: targetPath,
	}
}

func (out *Output) AddOperator(path string, value interface{}) {
	op := OpDefinition{
		Type:  "replace",
		Path:  path,
		Value: value,
	}
	out.Operators = append(out.Operators, op)
}

func (out *Output) String() string {
	data, err := yaml.Marshal(out.Operators)
	if err != nil {
		return err.Error()
	}
	return string(data[:])
}
