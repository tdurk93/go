package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	//z	float64 = 1
	var z float64 = 100;
	for delta := 100.0; delta > 0.00001; {
		var newZ float64 = z - (z*z - float64(x))/(2*z)
		//fmt.Println(z, newZ);
		delta = math.Abs(newZ - z);
		z = newZ;
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println("actual value:", math.Sqrt(2))
}

