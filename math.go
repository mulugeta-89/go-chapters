package main

import (
	"fmt"
	"math"
)
var print = fmt.Println
func main(){
	print("square root of 0", math.Sqrt(8))
	print("absolute value of -8", math.Abs(-8))
	print("ceiling of 9.1", math.Ceil(9.1))
	print("floor of 9.9", math.Floor(9.9))
	print("rounding 34.3", math.Round(34.3))
	print("cube root of 8", math.Cbrt(8))
	print("max between 5 and 6", math.Max(5,6))
	print("min between 5 and 6", math.Min(5,6))
	print("2 the power of 5", math.Pow(2,5))


}