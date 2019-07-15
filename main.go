package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := randomState(100, 100)
	printBoard(a)
}


func printBoard(b [][]int) {

	fmt.Print(string('|'))
	for x := 1; x <= len(b[0]); x++ {
		fmt.Print(string('-'))
	}
	fmt.Println(string('|'))

	for y := 0; y < len(b); y++ {
		fmt.Print(string('|'))
		for x := 0; x < len(b[y]); x++ {
			if b[y][x] == 1 {
				fmt.Print(string('#'))
			} else {
				fmt.Print(string(' '))
			}
		}
		fmt.Println(string('|'))
	}

	fmt.Print(string('|'))
	for x := 1; x <= len(b[0]); x++ {
		fmt.Print(string('-'))
	}
	fmt.Println(string('|'))
}

func randomState(w, h int) (a [][]int) {
	var b []int


	for j := 0; j < h; j++ {

		for i := 0; i < w; i++ {

			s := rand.NewSource(time.Now().UnixNano())
			r := rand.New(s)
			b = append(b, r.Intn(2))
		}

		a = append(a, b)
		b = []int{}
	}

	return a
}
