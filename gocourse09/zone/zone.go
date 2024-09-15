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
	animals []Animal
}

func (fz *FeedingZone) GetAnimals() []Animal {
	return fz.animals
}

func (fz *FeedingZone) AddAnimal(a Animal) {
	fz.animals = append(fz.animals, a)
}
