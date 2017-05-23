package parameterbag

import "github.com/pkg/errors"

//Factory creates a new parameter bag
//TODO - should this take an interface?
type Constructor func() Bag

//Factory is a generic factory for Parameter bags
type Factory struct {
	constructors map[string]Constructor
}

//AddFactory allows you to register a factory type which can be called later via Create
func (f *Factory) AddFactory(name string, factory Constructor) error {
	if factory == nil {
		return errors.Errorf("Bag factory %s does not exist.", name)
	}
	_, registered := f.constructors[name]
	if registered {
		return errors.Errorf("Bag factory %s is already registered.", name)
	}
	f.constructors[name] = factory
	return nil
}

//Create instantiates an named Bag interface
func (f *Factory) Create(id string) (Bag, error) {
	bagFactory, ok := f.constructors[id]
	if !ok {
		return nil, errors.New("Invalid Bag Type.")
	}
	return bagFactory(), nil
}
