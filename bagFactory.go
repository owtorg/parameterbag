package parameterbag

import "errors"

//Constructor creates a new parameter bag
type Constructor func() Bag

//Factory is a generic factory for Parameter bags
type Factory struct {
	constructors map[string]Constructor
}

//AddFactory allows you to register a factory type which can be called later via Create
func (f *Factory) AddFactory(name string, factory Constructor) error {
	if factory == nil {
		return errors.New("Bag factory " + name + " does not exist.")
	}
	_, registered := f.constructors[name]
	if registered {
		return errors.New("Bag factory " + name + " is already registered.")
	}
	f.constructors[name] = factory
	return nil
}

//Create instantiates an named Bag interface
func (f *Factory) Create(id string) (Bag, error) {
	bagFactory, ok := f.constructors[id]
	if !ok {
		return nil, errors.New("Invalid bag type")
	}
	return bagFactory(), nil
}
