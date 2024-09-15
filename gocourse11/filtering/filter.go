package filtering

type cleanerLevel int

const (
	No cleanerLevel = iota
	Low
	Middle
	High
)

type Filter struct {
	cleanLevel    cleanerLevel
	absorber      string
	waterImprover string
}

type option struct {
	cleanLevel    cleanerLevel
	absorber      string
	waterImprover string
}

type Option func(opt *option)

func WithCleanLevel(cleanLevel cleanerLevel) Option {
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

func GetFilter(cleanerLevel cleanerLevel) Filter {
	if cleanerLevel == Low {
		return NewFilter(
			WithCleanLevel(Low),
			WithAbsorber("sand"),
			WithWaterImprover("t2w"),
		)
	}

	if cleanerLevel == Middle {
		return NewFilter(
			WithCleanLevel(Middle),
			WithAbsorber("coal"),
			WithWaterImprover("cn2"),
		)
	}

	if cleanerLevel == High {
		return NewFilter(
			WithCleanLevel(High),
			WithAbsorber("vibranium"),
			WithWaterImprover("yy78"),
		)
	}

	return NewFilter(WithCleanLevel(No))
}
