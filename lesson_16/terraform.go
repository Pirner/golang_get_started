package main

import (
	"fmt"
)

type planet string

func (p planet) terraform() planet {
	return planet("New " + p)
}

func main() {
	planets := []planet{
		"Mercury", "Venus", "Earth", "Mars",
		"Jupiter", "Saturn", "Uranus", "Neptune",
	}
	fmt.Println(planets)
	fmt.Println(planets[0].terraform())
	planets[0] = planets[0].terraform()
	fmt.Println(planets)
}
