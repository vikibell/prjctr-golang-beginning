package terrarium

type Aquarium struct {
	size           int
	animalName     string
	saltLevel      int
	pollutionLevel int
}

func (a Aquarium) AnimalName() string {
	return a.animalName
}

func (a Aquarium) SaltLevel() int {
	return a.saltLevel
}

func (a Aquarium) PollutionLevel() int {
	return a.pollutionLevel
}

func (a Aquarium) Size() int {
	return a.size
}

type Builder interface {
	SetSize(size int) Builder
	SetAnimal(animalName string) Builder
	SetSaltLevel(level int) Builder
	SetPollutionLevel(pLevel int) Builder
	Build() *Aquarium
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct(size int, animalName string, level int, pLevel int) *Aquarium {
	return d.builder.
		SetSize(size).
		SetAnimal(animalName).
		SetSaltLevel(level).
		SetPollutionLevel(pLevel).
		Build()
}
