package main

import (
	"fmt"
)

type tempConvFn func(float64) float64

func celsiusToFahrenheit(c float64) float64 {
	return c*9.0/5.0 + 32
}

func fahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5.0 / 9.0
}

func drawTable(tempConv tempConvFn, t float64) {
	fmt.Printf("|\t%.2f\t|\t%.2f\t|\n", t, tempConv(t))
}

func main() {
	fmt.Println("=================================")
	fmt.Println("|\t摄氏度\t|\t华氏度\t|")
	fmt.Println("=================================")
	for i := 0; i <= 140/5; i++ {
		drawTable(celsiusToFahrenheit, float64(-40+i*5))
	}
	fmt.Println("=================================\n")
	fmt.Println("=================================")
	fmt.Println("|\t华氏度\t|\t摄氏度\t|")
	fmt.Println("=================================")
	for i := 0; i <= 140/5; i++ {
		drawTable(fahrenheitToCelsius, float64(-40+i*5))
	}
	fmt.Println("=================================")
}
