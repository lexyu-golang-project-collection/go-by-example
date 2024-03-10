package main

import (
	"errors"
	"fmt"
	"math"
)

type areaError struct {
	err    string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{
			err:    "radius is negative",
			radius: radius,
		}
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		var areaError *areaError
		if errors.As(err, &areaError) {
			fmt.Printf("Area calculation failed, radius %0.2f is less than zero", areaError.radius)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of rectangle %0.2f", area)
}
