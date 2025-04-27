package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	width  = 80
	height = 15
)

type Universe [][]bool

func NewUniverse() Universe {
	a := make([][]bool, height)
	for i := range a {
		a[i] = make([]bool, width)
	}
	return Universe(a)
}

func (u Universe) Show() {
	for y := range u {
		row := u[y]
		fmt.Printf("|")
		for x := range row {
			// fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
			if row[x] {
				fmt.Printf("o|")
			} else {
				fmt.Printf(" |")
			}
		}
		fmt.Printf("|")
		fmt.Printf("\n")
	}
}

func (u Universe) Seed() {
	for y := range u {
		for x := range u[y] {
			p := rand.Intn(100)
			if p <= 25 {
				u[y][x] = true
			}
		}
	}
}

func (u Universe) IsAlive(x int, y int) bool {
	x = (x + width) % width
	y = (y + height) % height
	return u[y][x]
}

func (u Universe) NumberAliveNeighbors(x int, y int) int {
	n := 0
	for y_off := -1; y_off <= 1; y_off++ {
		for x_off := -1; x_off <= 1; x_off++ {
			if (y_off != 0 && x_off != 0) && u.IsAlive(x+x_off, y+y_off) {
				n++
			}
		}
	}
	return n
}

func (u Universe) Next(x, y int) bool {
	n := u.NumberAliveNeighbors(x, y)
	// first if self is dead then if 3 neighbors are alive become alive
	if !u.IsAlive(x, y) && (n == 3) {
		return true
	} else if u.IsAlive(x, y) && ((n == 3) || (n == 2)) {
		return true
	}
	return false
}

func (u Universe) Set(x, y int, b bool) {
	u[y][x] = b
}

func step(a, b Universe) {

	// iterate over the entire universe
	for y := range a {
		for x := range a[y] {
			nextState := a.Next(x, y)
			b.Set(x, y, nextState)
		}
	}
}

func main() {
	fmt.Println("fin")
	a := NewUniverse()
	a.Seed()
	b := NewUniverse()
	b.Seed()
	n_steps := 10
	for i := 0; i < n_steps; i++ {
		// fmt.Print("\x0c")
		step(a, b)
		a.Show()
		fmt.Print("\n\n\n\n")
		time.Sleep(5 * time.Second)
		a, b = b, a
	}

	fmt.Println(a.NumberAliveNeighbors(0, 0))
}
