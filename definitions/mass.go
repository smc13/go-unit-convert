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

var MetricMassSystem = SystemSet{
	System: &metricMassSystem,
	Units: []*Unit{
		&Microgram,
		&Milligram,
		&Gram,
		&Kilogram,
		&Tonne,
	},
}
