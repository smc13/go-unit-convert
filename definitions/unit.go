package definitions

type UnitName struct {
	Singular string
	Plural   string
}

type Unit struct {
	Name        UnitName
	Abbr        string
	System      *System
	ToAnchor    float64
	AnchorShift float64
}

type UnitSet []Unit
