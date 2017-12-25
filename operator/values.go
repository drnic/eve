package operator

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// LoadValues loads values from Target file or STDIN
func (out *Output) LoadValues() (err error) {
	data, err := ioutil.ReadFile(out.TargetPath)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &out.Operators)
	return
}
