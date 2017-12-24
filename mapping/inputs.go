package mapping

type Inputs struct {
	ValuesByName map[string]string
}

func NewInputsFromFlags(valuesByName map[string]string) (inputs *Inputs) {
	inputs = &Inputs{
		ValuesByName: valuesByName,
	}
	return
}

func (inputs *Inputs) ValueForFormName(name string) (value string) {
	return inputs.ValuesByName[name]
}
