package filter

func Create(cleanLevel CleanLevel) Filter {
	switch cleanLevel {
	case Low:
		return NewFilter(
			WithCleanLevel(cleanLevel),
			WithAbsorber("sand"),
			WithWaterImprover("t2w"),
		)
	case Middle:
		return NewFilter(
			WithCleanLevel(cleanLevel),
			WithAbsorber("coal"),
			WithWaterImprover("cn2"),
		)
	case High:
		return NewFilter(
			WithCleanLevel(cleanLevel),
			WithAbsorber("vibranium"),
			WithWaterImprover("yy78"),
		)
	default:
		return NewFilter(WithCleanLevel(No))
	}
}

func Select(pollutionLevel int) Filter {
	switch {
	case pollutionLevel >= 100 && pollutionLevel <= 500:
		return Create(Low)
	case pollutionLevel > 500 && pollutionLevel < 1000:
		return Create(Middle)
	case pollutionLevel >= 1000:
		return Create(High)
	default:
		return Create(No)
	}
}
