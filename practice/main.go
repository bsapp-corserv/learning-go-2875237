package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3
	fmt.Println("Integer Sum: ", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("floatSum Sum: ", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("floatSum ROUNDED: ", floatSum)

	circleRadius := 15.5

	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("circumference: %.2f\n", circumference)

	n := time.Now()
	fmt.Println("the time is current: ", n)

	t := time.Date(2008, time.August, 16, 16, 0, 0, 0, time.UTC)
	fmt.Println("OLD DATE: ", t)
	fmt.Println("OLD DATE 1: ", t.Format(time.ANSIC))

	prasedTime, _ := time.Parse(time.ANSIC, "Sat Aug 16 16:00:00 2008")
	fmt.Printf("the type of parsed time is %T\n", prasedTime)
}
