package main

import (
	"fmt"
	"time"
)

// dump slice length, capacity, and contents
func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}
	dump("dwarfs", dwarfs)
	for i := 0; i < 50; i++ {
		dwarfs = append(dwarfs, "Foobar")
		dump("dwarfs", dwarfs)
		time.Sleep(1 * time.Second)
	}
}
