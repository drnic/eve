package mapping

type Inputs struct {
	ValuesByName map[string]string
}

func NewInputs() (inputs *Inputs) {
	inputs = &Inputs{
		ValuesByName: map[string]string{},
	}
	return
}

func NewInputsFromFlags(valuesByName map[string]string) (inputs *Inputs) {
	inputs = &Inputs{
		ValuesByName: valuesByName,
	}
	return
}

func (inputs *Inputs) ValueForFormName(name string) (value interface{}) {
	return inputs.ValuesByName[name]
}
