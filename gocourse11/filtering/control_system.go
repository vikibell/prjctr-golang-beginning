package filtering

func CreateFilter(cleanerLevel CleanerLevel) Filter {
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

func SelectFilter(pollutionLevel int) Filter {
	switch {
	case pollutionLevel >= 100 && pollutionLevel <= 500:
		return CreateFilter(Low)
	case pollutionLevel > 500 && pollutionLevel < 1000:
		return CreateFilter(Middle)
	case pollutionLevel >= 1000:
		return CreateFilter(High)
	default:
		return CreateFilter(No)
	}
}
