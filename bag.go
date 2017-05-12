package parameterbag

//Bag - interface for a parameter bag
type Bag interface {
	//Get a parameter bag value. If value does not exist should return ""
	Get(string) string
	//Set a parameter value
	Set(string, string)
	//Has a parameter or not?
	Has(string) bool
	//List all parameters that exist in this bag
	ListParameters() []string
}

//ParameterBag - concrete definition of a parameter bag type
type ParameterBag struct {
	params map[string]string
}

// NewParameterBag - create a new parameter bag with a blank list of parameters
func NewParameterBag() *ParameterBag {
	bag := make(map[string]string)
	return &ParameterBag{bag}

}

// NewParameterBagFromMap - Create a new parameter bag by passing in an existing map[string]string
func NewParameterBagFromMap(params map[string]string) *ParameterBag {
	return &ParameterBag{params: params}
}

//Get returns the specified parameter
func (b *ParameterBag) Get(name string) string {
	return b.params[name]
}

//Set sets the specified parameter
func (b *ParameterBag) Set(name string, value string) {
	b.params[name] = value
}

//Has this bag got the specified parameter?
func (b *ParameterBag) Has(name string) bool {
	return !(b.params[name] == "")
}

//ListParameters gets a slice of all the parameter keys in this bag
func (b *ParameterBag) ListParameters() []string {
	keys := make([]string, len(b.params))
	i := 0
	for k := range b.params {
		keys[i] = k
		i++
	}
	return keys
}
