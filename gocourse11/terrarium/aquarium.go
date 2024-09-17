package terrarium

type Aquarium struct {
	size           int
	animalName     string
	saltLevel      int
	pollutionLevel int
}

func (a Aquarium) PollutionLevel() int {
	return a.pollutionLevel
}

type Builder interface {
	SetSize(size int) Builder
	SetAnimal(animalName string) Builder
	SetSaltLevel(level int) Builder
	SetPollutionLevel(level int) Builder
	Build() *Aquarium
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director {
	return &Director{builder: builder}
}

func (d *Director) Construct(size int, animalName string, saltLevel int, pollutionLevel int) *Aquarium {
	return d.builder.
		SetSize(size).
		SetAnimal(animalName).
		SetSaltLevel(saltLevel).
		SetPollutionLevel(pollutionLevel).
		Build()
}
