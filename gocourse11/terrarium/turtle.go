package terrarium

type TurtleAquariumBuilder struct {
	aquarium *Aquarium
}

func NewTurtleAquariumBuilder() *TurtleAquariumBuilder {
	return &TurtleAquariumBuilder{aquarium: &Aquarium{}}
}

func (ta *TurtleAquariumBuilder) SetSize(size int) Builder {
	ta.aquarium.size = size
	return ta
}

func (ta *TurtleAquariumBuilder) SetAnimal(animalName string) Builder {
	ta.aquarium.animalName = "Turtle " + animalName
	return ta
}

func (ta *TurtleAquariumBuilder) SetSaltLevel(level int) Builder {
	ta.aquarium.saltLevel = level
	return ta
}

func (ta *TurtleAquariumBuilder) Build() *Aquarium {
	return ta.aquarium
}
