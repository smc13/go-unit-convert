package main

import (
	"encoding/json"
	"fmt"

	"github.com/smc13/go-unit-convert/definitions"
)

func NewValue(value float64, unit definitions.Unit) *Value {
	return &Value{
		unit:  unit,
		value: value,
	}
}

type Value struct {
	unit  definitions.Unit
	value float64
}

func (v *Value) Unit() definitions.Unit {
	return v.unit
}

func (v *Value) Value() float64 {
	return v.value
}

func (v *Value) String() string {
	// format the value to a maximum of 4 decimal places (or less if they're all 0)
	return fmt.Sprintf("%.4g%s", v.value, v.unit.Abbr)
}

// marshalJSON is a custom JSON marshaller for Value
func (v *Value) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Value     float64
		Unit      string
		Formatted string
	}{
		Value:     v.value,
		Unit:      v.unit.Name.Singular,
		Formatted: v.String(),
	})
}

func (v *Value) To(unit definitions.Unit) (*Value, error) {
	if v.unit == unit {
		return v, nil
	}

	if v.unit.System.Measure != unit.System.Measure {
		return nil, fmt.Errorf("cannot convert incompatible measures: %s to %s", v.unit.System.Measure, unit.System.Measure)
	}

	result := v.value * v.unit.ToAnchor
	if v.unit.AnchorShift != 0 {
		result -= v.unit.AnchorShift
	}

	// convert between systems
	if v.unit.System != unit.System {
		result = v.unit.System.Transform(result)
	}

	// apply the new unit's anchor shift
	if unit.AnchorShift != 0 {
		result += unit.AnchorShift
	}

	result /= unit.ToAnchor

	return &Value{
		unit:  unit,
		value: result,
	}, nil
}

func (v *Value) ToBest(set definitions.UnitSet) (*Value, error) {
	isNegative := v.value < 0
	cutoff := 1.0
	if isNegative {
		cutoff = -1.0
	}

	var best *Value

	for _, u := range set {
		result, err := v.To(u)

		if err != nil && (isNegative && result.value > cutoff) || (!isNegative && result.value < cutoff) {
			continue
		}

		if best == nil || ((isNegative && result.value <= cutoff && result.value > best.value) || (!isNegative && result.value >= cutoff && result.value < best.value)) {
			best = result
		}
	}

	if best == nil {
		return nil, fmt.Errorf("no units in the set are compatible with %s", v.unit.Name.Singular)
	}

	return best, nil
}
