package mapping

import "strconv"

type Inputs struct {
	ValuesByName map[string]interface{}
}

func NewInputs() (inputs *Inputs) {
	inputs = &Inputs{
		ValuesByName: map[string]interface{}{},
	}
	return
}

func NewInputsFromFlags(valuesByName map[string]string) (inputs *Inputs) {
	inputs = NewInputs()
	for name, value := range valuesByName {
		valueAsInt, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			inputs.ValuesByName[name] = valueAsInt
		} else {
			inputs.ValuesByName[name] = value
		}
	}
	return
}

func (inputs *Inputs) ValueForFormName(name string) (value interface{}) {
	return inputs.ValuesByName[name]
}
