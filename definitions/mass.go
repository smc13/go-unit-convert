package definitions

const MASS_MEASURE = "Mass"

// metric
var metricMassSystem = System{
	Name:    METRIC_SYSTEM_NAME,
	Measure: MASS_MEASURE,
	Ratio:   1 / 453.592,
}

var Microgram = Unit{
	System: &metricMassSystem,
	Name: UnitName{
		Singular: "Microgram",
		Plural:   "Micrograms",
	},
	Abbr:     "mcg",
	ToAnchor: 1 / 1000000,
}

var Milligram = Unit{
	System: &metricMassSystem,
	Name: UnitName{
		Singular: "Milligram",
		Plural:   "Milligrams",
	},
	Abbr:     "mg",
	ToAnchor: 1 / 1000,
}

var Gram = Unit{
	System: &metricMassSystem,
	Name: UnitName{
		Singular: "Gram",
		Plural:   "Grams",
	},
	Abbr:     "g",
	ToAnchor: 1,
}

var Kilogram = Unit{
	System: &metricMassSystem,
	Name: UnitName{
		Singular: "Kilogram",
		Plural:   "Kilograms",
	},
	Abbr:     "kg",
	ToAnchor: 1000,
}

var Tonne = Unit{
	System: &metricMassSystem,
	Name: UnitName{
		Singular: "Metric Tonne",
		Plural:   "Metric Tonnes",
	},
	Abbr:     "mt",
	ToAnchor: 1000000,
}

var MetricMassSystem = UnitSet{
	Microgram,
	Milligram,
	Gram,
	Kilogram,
	Tonne,
}

// imperial

var imperialMassSystem = System{
	Name:    IMPERIAL_SYSTEM_NAME,
	Measure: MASS_MEASURE,
	Ratio:   453.592,
}

var Ounce = Unit{
	Name: UnitName{
		Singular: "Ounce",
		Plural:   "Ounces",
	},
	Abbr:     "oz",
	ToAnchor: 1 / 16,
}

var Pound = Unit{
	Name: UnitName{
		Singular: "Pound",
		Plural:   "Pounds",
	},
	Abbr:     "lb",
	ToAnchor: 1,
}

var Ton = Unit{
	Name: UnitName{
		Singular: "Ton",
		Plural:   "Tons",
	},
	Abbr:     "t",
	ToAnchor: 2000,
}

var ImperialMassSystem = UnitSet{
	Ounce,
	Pound,
	Ton,
}
