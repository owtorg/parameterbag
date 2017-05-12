package parameterbag

type Bag interface {
	Get(string) string
	Set(string, string)
	Has(string) bool
	ListParameters()
}

type params map[string]string

type ParameterBag struct {
	params
}

func NewParameterBag() *ParameterBag {
	bag := make(map[string]string)
	return &ParameterBag{bag}

}
func NewParameterBagFromMap(params params) *ParameterBag {
	return &ParameterBag{params: params}
}

func (b *ParameterBag) Get(name string) string {
	return b.params[name]
}

func (b *ParameterBag) Set(name string, value string) {
	b.params[name] = value
}

func (b *ParameterBag) Has(name string) bool {
	return !(b.params[name] == "")
}

func (b *ParameterBag) ListParameters() []string {
	keys := make([]string, len(b.params))
	i := 0
	for k := range b.params {
		keys[i] = k
		i++
	}
	return keys
}
