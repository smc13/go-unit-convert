package definitions

const LENGTH_MEASURE = "Length"

var metricLengthSystem = System{
	Name:    "Metric",
	Measure: LENGTH_MEASURE,
	Ratio:   3.28084,
}

var Nanometer = Unit{
	System: &metricLengthSystem,
	Name: UnitName{
		Singular: "Nanometer",
		Plural:   "Nanometers",
	},
	Abbr:     "nm",
	ToAnchor: 1e-9,
}

var Micrometer = Unit{
	System: &metricLengthSystem,
	Name: UnitName{
		Singular: "Micrometer",
		Plural:   "Micrometers",
	},
	Abbr:     "μm",
	ToAnchor: 1e-6,
}

var Millimeter = Unit{
	System: &metricLengthSystem,
	Name: UnitName{
		Singular: "Millimeter",
		Plural:   "Millimeters",
	},
	Abbr:     "mm",
	ToAnchor: 1e-3,
}

var Centimeter = Unit{
	System: &metricLengthSystem,
	Name: UnitName{
		Singular: "Centimeter",
		Plural:   "Centimeters",
	},
	Abbr:     "cm",
	ToAnchor: 1e-2,
}

var Meter = Unit{
	System: &metricLengthSystem,
	Name: UnitName{
		Singular: "Meter",
		Plural:   "Meters",
	},
	Abbr:     "m",
	ToAnchor: 1,
}

var Kilometer = Unit{
	System: &metricLengthSystem,
	Name: UnitName{
		Singular: "Kilometer",
		Plural:   "Kilometers",
	},
	Abbr:     "km",
	ToAnchor: 1e3,
}

var MetricLengths = UnitSet{
	Nanometer,
	Micrometer,
	Millimeter,
	Centimeter,
	Meter,
	Kilometer,
}

// imperial units

var imperialLengthSystem = System{
	Name:    "Imperial",
	Measure: LENGTH_MEASURE,
	Ratio:   1 / 3.28084,
}

var Inch = Unit{
	System: &imperialLengthSystem,
	Name: UnitName{
		Singular: "Inch",
		Plural:   "Inches",
	},
	Abbr:     "in",
	ToAnchor: 1 / 12,
}

var Yard = Unit{
	System: &imperialLengthSystem,
	Name: UnitName{
		Singular: "Yard",
		Plural:   "Yards",
	},
	Abbr:     "yd",
	ToAnchor: 3,
}

var UsSurveyFoot = Unit{
	System: &imperialLengthSystem,
	Name: UnitName{
		Singular: "US Survey Foot",
		Plural:   "US Survey Feet",
	},
	Abbr:     "ft-us",
	ToAnchor: 1.000002,
}

var Foot = Unit{
	System: &imperialLengthSystem,
	Name: UnitName{
		Singular: "Foot",
		Plural:   "Feet",
	},
	Abbr:     "ft",
	ToAnchor: 1,
}

var Fathom = Unit{
	System: &imperialLengthSystem,
	Name: UnitName{
		Singular: "Fathom",
		Plural:   "Fathoms",
	},
	Abbr:     "ftm",
	ToAnchor: 6,
}

var Mile = Unit{
	System: &imperialLengthSystem,
	Name: UnitName{
		Singular: "Mile",
		Plural:   "Miles",
	},
	Abbr:     "mi",
	ToAnchor: 5280,
}

var ImperialLengths = UnitSet{
	Inch,
	Yard,
	UsSurveyFoot,
	Foot,
	Fathom,
	Mile,
}
