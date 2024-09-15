package terrarium

type Aquarium struct {
	size           int
	animalName     string
	saltLevel      int
	pollutionLevel int
}

func (a Aquarium) GetAnimalName() string {
	return a.animalName
}

func (a Aquarium) GetSaltLevel() int {
	return a.saltLevel
}

func (a Aquarium) GetPollutionLevel() int {
	return a.pollutionLevel
}

func (a Aquarium) GetSize() int {
	return a.size
}

type Builder interface {
	SetSize(size int) Builder
	SetAnimal(animalName string) Builder
	SetSaltLevel(level int) Builder
	Build() *Aquarium
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct(size int, animalName string, level int) *Aquarium {
	return d.builder.
		SetSize(size).
		SetAnimal(animalName).
		SetSaltLevel(level).
		Build()
}
