package definitions

const METRIC_SYSTEM_NAME = "Metric"
const IMPERIAL_SYSTEM_NAME = "Imperial"

type System struct {
	Name  string
	Measure string
	Ratio float64
	TransformFunc func(float64) float64
}

func (s *System) String() string {
	return s.Name
}

func (s *System) Transform(val float64) float64 {
	if s.TransformFunc != nil {
		return s.TransformFunc(val)
	}

	if s.Ratio != 0 {
		return val * s.Ratio
	}

	return val
}