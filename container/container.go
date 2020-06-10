package container

type Container struct {
	Baggages []Baggage
}

func New() Container {
	return Container{}
}

type Baggage struct {
	Name    string
	Content string
}

func (container *Container) Append(baggage Baggage) {
	container.Baggages = append(container.Baggages, baggage)
}

func (contaner Container) Find(name string) Baggage {
	for _, baggage := range contaner.Baggages {
		if baggage.Name == name {
			return baggage
		}
	}

	return Baggage{}
}
