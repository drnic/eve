package operator

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// LoadValues loads values from Target file or STDIN
func (out *Output) LoadValues() (err error) {
	var data []byte
	if out.TargetPath == "" || out.TargetPath == "-" {
		data, err = ioutil.ReadAll(os.Stdin)
	} else {
		data, err = ioutil.ReadFile(out.TargetPath)
	}
	if err != nil {
		return
	}
	err = yaml.Unmarshal(data, &out.Operators)
	return
}
