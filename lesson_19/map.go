package main

import (
	"fmt"
)

func main() {
	temperature := map[string]int{
		"Earth": 15,
		"Mars":  -65,
		"Moon":  32,
	}
	fmt.Println(temperature)

	if moon, ok := temperature["Moon"]; ok {
		fmt.Printf("On average the moon is %vº C.\n", moon)
	} else {
		fmt.Println("Where is the moon?")
	}
}
