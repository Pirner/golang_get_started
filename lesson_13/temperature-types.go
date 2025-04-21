package main

import (
	"fmt"
)

type celsius float64
type kelvin float64

// convert kelvin to celsius
func kelvinToCelsius(k kelvin) celsius {
	return celsius(k - 273.15)
}
func celsiusToKelvin(c celsius) kelvin {
	return kelvin(c + 273.15)
}
func (k kelvin) celsius() celsius {
	return celsius(k - 273.15)
}
func (c celsius) kelvin() kelvin {
	return kelvin(c + 273.15)
}

func main() {
	var k kelvin = 294.0
	c := k.celsius()
	fmt.Println(c)
	var surface_sun celsius = 127
	k_surface_sun := surface_sun.kelvin()
	fmt.Println(k_surface_sun)
}
