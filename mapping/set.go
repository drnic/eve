package mapping

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/starkandwayne/eve/operator"

	"gopkg.in/yaml.v2"
)

// Set describes a mapping of form values to operator file
//
// Example:
//   mappings:
//   - name: workers-linux-instances
//     path: /instance_groups/name=worker/instances
//   - name: workers-linux-instance-type
//     path: /instance_groups/name=worker/vm_type
type Set struct {
	Mappings []Mapping `yaml:"mappings"`
}

// Mapping of a form name to an operator path replacement
type Mapping struct {
	FormName     string `yaml:"name"`
	OperatorPath string `yaml:"path"`
}

// NewMappingSet creates a set of Mappings from a local file
func NewMappingSet(path string) (set *Set, err error) {
	set = &Set{}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(data, set)
	if err != nil {
		return
	}
	return
}

// GenerateOutput applies Inputs to Set mapping to create operator.Output
func (set *Set) GenerateOutput(inputs *Inputs, output *operator.Output) {
	for _, mapping := range set.Mappings {
		var value interface{} = inputs.ValueForFormName(mapping.FormName)
		fmt.Fprintf(os.Stderr, "%s -> %v\n", mapping.OperatorPath, value)
		output.AddOperator(mapping.OperatorPath, value)
	}
}

// LoadValues pulls out values from Operator file based on mapping set
func (set *Set) LoadValues(output *operator.Output) (values *Inputs, err error) {
	if err = output.LoadValues(); err != nil {
		return
	}
	values = NewInputs()
	for _, mapping := range set.Mappings {
		for _, op := range output.Operators {
			if op.Path == mapping.OperatorPath {
				values.ValuesByName[mapping.FormName] = op.Value.(string)
			}
		}
	}
	return
}
