package filter

func Create(cleanLevel CleanLevel) Filter {
	if cleanLevel == Low {
		return NewFilter(
			WithCleanLevel(Low),
			WithAbsorber("sand"),
			WithWaterImprover("t2w"),
		)
	}

	if cleanLevel == Middle {
		return NewFilter(
			WithCleanLevel(Middle),
			WithAbsorber("coal"),
			WithWaterImprover("cn2"),
		)
	}

	if cleanLevel == High {
		return NewFilter(
			WithCleanLevel(High),
			WithAbsorber("vibranium"),
			WithWaterImprover("yy78"),
		)
	}

	return NewFilter(WithCleanLevel(No))
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
