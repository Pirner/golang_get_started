package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const lightSpeed = 299792
	var distance = rand.Intn(10)
	var age = 30
	fmt.Println(age)
	age++
	fmt.Println(age)
	fmt.Println(149.0 * 0.3783)
	fmt.Println("Another Line for me to print")
	fmt.Println(41 * 365 / 687)
	fmt.Println(distance)
	fmt.Println("My weight on the surface of Mars is", 149.0*0.3783, "lbs, and "+
		"I would be", 41*365.2425/687, "years old.")
	fmt.Printf("My weight on the surface of Mars is %v lbs,", 149.0*0.3783)
	fmt.Printf(" and I would be %v years old.\n", 41*365/687)
}
