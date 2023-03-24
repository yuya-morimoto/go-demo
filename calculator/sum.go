package calculator

import "fmt"

// アッパーキャメルの場合は外部から参照できる
var offset float64 = 1
var Offset float64 = 1

func Sum(a float64, b float64) float64 {
	fmt.Println("multiply: ", multiply(a, b))
	fmt.Println("Multiply: ", Multiply(a, b))
	return a + b + offset
}
