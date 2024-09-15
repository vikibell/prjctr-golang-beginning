package zone

type Specie int

const (
	Bull Specie = iota
	Horse
	Buffalo
	Elk
)

type Animal struct {
	Name   string
	Specie Specie
}

type Zone interface {
	GetAnimals() []Animal
	AddAnimal(a Animal)
}

type FeedingZone struct {
	Animals []Animal
}

func (fz *FeedingZone) GetAnimals() []Animal {
	return fz.Animals
}

func (fz *FeedingZone) AddAnimal(a Animal) {
	fz.Animals = append(fz.Animals, a)
}
