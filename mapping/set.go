package mapping

import (
	"io/ioutil"

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
		output.AddOperator(mapping.OperatorPath, value)
	}
}
