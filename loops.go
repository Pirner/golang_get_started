package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var targetNumber = 25
	for {
		var randomNumber = rand.Intn(100)

		if targetNumber == randomNumber {
			fmt.Printf("target number is %d\n", targetNumber)
			break
		} else {
			if targetNumber > randomNumber {
				fmt.Println(randomNumber, "target is bigger, increase")
			} else {
				fmt.Println(randomNumber, "target is smaller, decrease")
			}
		}
	}
}
