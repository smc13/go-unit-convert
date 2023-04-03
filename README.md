# go-unit-convert

go-unit-convert is a Go package that provides functionality for converting units of measurement. The package provides a simple and intuitive API that can be used to convert values between different units of measurement.

## Installation
To use go-unit-convert, you need to have Go installed on your system. Once you have Go installed, you can install the package using the following command:

```sh
go get github.com/your-username/go-unit-convert
```

## Usage
To use go-unit-convert, you first need to create a new value using the NewValue function. This function takes two arguments: the value you want to convert, and a pointer to the unit definition you want to use. The unit definition specifies the unit of measurement you want to convert from.

Here's an example of how you can use NewValue to create a value representing 3000 meters:

```golang
import (
    "github.com/your-username/go-unit-convert/definitions"
    "github.com/your-username/go-unit-convert"
)

value := unitconvert.NewValue(3000, &definitions.Meter)
```

Once you have created a value, you can use the To method to convert it to another unit of measurement. The `To` method takes a pointer to the unit definition you want to convert to. Here's an example of how you can use To to convert the 3000 meter value to kilometers:

```golang
result, err := value.To(&definitions.Kilometer)
```

The `To` method returns a new value that represents the converted value and the unit its in.

In addition to the `To` method, go-unit-convert also provides a `ToBest` method that converts a value to the "best" unit of measurement in a given category. The `ToBest` method takes a pointer to a `SystemSet` definition, which specifies the category of units to search for the "best" unit.

Here's an example of how you can use ToBest to convert the 1.8288 meter value to the "best" imperial unit:

```golang
result := unitconvert.NewValue(1.8288, &definitions.Meter).ToBest(&definitions.ImperialLengths)
```

The ToBest method returns a new value that represents the converted value in the "best" unit of measurement in the specified category. You can access the converted value using the Value method.
