package main

import "fmt"

type celsius float64
type fahrenheit float64
type kelvin float64

// implement celsius to kelvin
func (c celsius) kelvin() kelvin { return kelvin(c + 273.15) }

// implement celsius to Fahrenheit
func (c celsius) fahrenheit() fahrenheit { return fahrenheit(c*9.0/5.0 + 32.0) }

// implement fahrenheit to celsius
func (f fahrenheit) celsius() celsius { return celsius(f-32.0) / (9.0 / 5.0) }

// implement fahrenheit to kelvin
func (f fahrenheit) kelvin() kelvin { return kelvin(f.celsius() + 273.15) }

// implement kelvin to celsius
func (k kelvin) celsius() celsius { return celsius(k - 273.15) }

// implement kelvin to fahrenheit
func (k kelvin) fahrenheit() fahrenheit { return fahrenheit(k.celsius().fahrenheit()) }

func drawTable(drawLine func(c celsius)) {
	// TODO draw header with a function given
	// TODO draw each single line with the function given
	for i := -40; i <= 100; i += 5 {
		drawCelsiusKelvinLine(celsius(i))
	}
}

func drawCelsiusKelvinLine(c celsius) {
	fmt.Printf("| %.2f \t\t | %.2f \t\t |\n", c, c.kelvin())
}

func main() {
	drawTable(drawCelsiusKelvinLine)
}
