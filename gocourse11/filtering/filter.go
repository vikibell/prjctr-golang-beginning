package filtering

type CleanerLevel int

const (
	No CleanerLevel = iota
	Low
	Middle
	High
)

type Filter struct {
	cleanLevel    CleanerLevel
	absorber      string
	waterImprover string
}

func (f Filter) CleanLevel() CleanerLevel {
	return f.cleanLevel
}

func (f Filter) Absorber() string {
	return f.absorber
}

func (f Filter) WaterImprover() string {
	return f.waterImprover
}

type option struct {
	cleanLevel    CleanerLevel
	absorber      string
	waterImprover string
}

type Option func(opt *option)

func WithCleanLevel(cleanLevel CleanerLevel) Option {
	return func(o *option) {
		o.cleanLevel = cleanLevel
	}
}

func WithAbsorber(absorber string) Option {
	return func(o *option) {
		o.absorber = absorber
	}
}

func WithWaterImprover(waterImprover string) Option {
	return func(o *option) {
		o.waterImprover = waterImprover
	}
}

func NewFilter(options ...Option) Filter {
	o := &option{}
	for _, opt := range options {
		opt(o)
	}
	return Filter{
		cleanLevel:    o.cleanLevel,
		absorber:      o.absorber,
		waterImprover: o.waterImprover,
	}
}
