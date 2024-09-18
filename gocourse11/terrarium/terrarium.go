package terrarium

import "github.com/vikibell/prjctr-golang-beginning/gocourse11/filter"

type Terrarium struct {
	aquariums []Aquarium
	filter    filter.Filter
}

func (t *Terrarium) Filter() filter.Filter {
	return t.filter
}

func (t *Terrarium) SetFilter(filter filter.Filter) {
	t.filter = filter
}

func (t *Terrarium) Aquariums() []Aquarium {
	return t.aquariums
}

func (t *Terrarium) AddAquarium(a Aquarium) {
	t.aquariums = append(t.aquariums, a)
}

func (t *Terrarium) SetAquariums(aquariums []Aquarium) {
	t.aquariums = aquariums
}

func NewTerrarium() *Terrarium {
	return &Terrarium{
		aquariums: make([]Aquarium, 0),
		filter:    filter.Create(filter.No),
	}
}

func (t *Terrarium) CalculatePollutionLevel() int {
	pollutionLevel := 0
	for _, aquarium := range t.aquariums {
		pollutionLevel += aquarium.PollutionLevel()
	}

	return pollutionLevel
}
