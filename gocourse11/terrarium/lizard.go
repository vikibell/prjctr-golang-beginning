package terrarium

type LizardAquariumBuilder struct {
	aquarium *Aquarium
}

func NewLizardAquariumBuilder() *LizardAquariumBuilder {
	return &LizardAquariumBuilder{aquarium: &Aquarium{}}
}

func (la *LizardAquariumBuilder) SetSize(size int) Builder {
	la.aquarium.size = size
	return la
}

func (la *LizardAquariumBuilder) SetAnimal(animalName string) Builder {
	la.aquarium.animalName = "Lizard " + animalName
	return la
}

func (la *LizardAquariumBuilder) SetSaltLevel(level int) Builder {
	la.aquarium.saltLevel = level
	return la
}

func (la *LizardAquariumBuilder) SetPollutionLevel(pLevel int) Builder {
	la.aquarium.pollutionLevel = pLevel
	return la
}

func (la *LizardAquariumBuilder) Build() *Aquarium {
	return la.aquarium
}
