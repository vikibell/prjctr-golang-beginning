package terrarium

import "gocourse10/filtering"

type Terrarium struct {
	aquariums []Builder
	filter    filtering.Filter
}

func (t *Terrarium) GetFilter() filtering.Filter {
	return t.filter
}

func (t *Terrarium) SetFilter(filter filtering.Filter) {
	t.filter = filter
}

func (t *Terrarium) addAquarium(a Builder) {
	t.aquariums = append(t.aquariums, a)
}

func NewTerrarium() *Terrarium {
	return &Terrarium{
		aquariums: make([]Builder, 0),
		filter:    filtering.GetFilter(filtering.No),
	}
}
