package terrarium

type SnakeAquariumBuilder struct {
	aquarium *Aquarium
}

func NewSnakeAquariumBuilder() *SnakeAquariumBuilder {
	return &SnakeAquariumBuilder{aquarium: &Aquarium{}}
}

func (sa *SnakeAquariumBuilder) SetSize(size int) Builder {
	sa.aquarium.size = size
	return sa
}

func (sa *SnakeAquariumBuilder) SetAnimal(animalName string) Builder {
	sa.aquarium.animalName = "Snake " + animalName
	return sa
}

func (sa *SnakeAquariumBuilder) SetSaltLevel(level int) Builder {
	sa.aquarium.saltLevel = level
	return sa
}

func (sa *SnakeAquariumBuilder) SetPollutionLevel(pLevel int) Builder {
	sa.aquarium.pollutionLevel = pLevel
	return sa
}

func (sa *SnakeAquariumBuilder) Build() *Aquarium {
	return sa.aquarium
}
