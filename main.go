package main

import (
	"fmt"

	"github.com/smc13/go-unit-convert/definitions"
)

func main() {
	fmt.Println(NewValue(3000, &definitions.Meter).To(&definitions.Kilometer))
	fmt.Println(NewValue(1.8288, &definitions.Meter).ToBest(&definitions.ImperialLengths))
	fmt.Println(NewValue(3000, &definitions.Meter).ToBest(&definitions.ImperialLengths))

	fmt.Println(NewValue(1, &definitions.Meter).To(&definitions.Kilogram))
	fmt.Println(NewValue(1, &definitions.Meter).ToBest(&definitions.MetricMassSystem))
}
