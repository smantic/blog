package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Printf("1\t = %064b\n", math.Float64bits(1.0))

	var tencents float64 = 0.1

	b := math.Float64bits(tencents)
	//fmt.Printf("0.1\t = %064b\n", math.Float32bits(float32(tencents)))
	fmt.Printf("0.1\t = %064b\n", b)
	//fmt.Printf("0.1\t = %d\n", math.Float64frombits(b))

	fmt.Printf("0.1 + (0.1 * 10^10) = %f\n", tencents+(tencents*math.Pow(10, 10)))
}
