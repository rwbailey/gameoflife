package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	a := randomState(188, 42)
	//a := gliderStart(188, 42)

	printBoard(a)

	//a = calculateNextBoardState(a)
	//printBoard(a)

	for {
		time.Sleep(500 * time.Millisecond)
		a = calculateNextBoardState(a)
		printBoard(a)
	}
}

func gliderStart(w, h int) (a [][]int) {
	var b []int

	for j := 0; j < h; j++ {

		for i := 0; i < w; i++ {
			b = append(b, 0)
		}

		a = append(a, b)
		b = []int{}
	}

	a[0][0] = 1
	a[1][1] = 1
	a[1][2] = 1
	a[2][0] = 1
	a[2][1] = 1


	return a
}

func calculateNextBoardState(b [][]int) [][]int {

	u := make ([][]int, len(b))
	for e := range b {
		u[e] = make([]int, len(b[e]))
		copy(u[e], b[e])
	}

	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[i]); j++ {
			calculateNextCellState(i, j, b, u)
		}
	}

	return u
}

func calculateNextCellState(y, x int, b [][]int, u [][]int) {

	N := 0

	for j := y-1; j <= y+1; j++ {
		for i := x-1; i <= x+1; i++ {
			if j == -1 || i == -1 {
				continue
			}
			if j == len(b) || i == len(b[0]) {
				continue
			}
			if j == y && i == x {
				continue
			}
			if b[j][i] == 1 {
				N++
			}
		}
	}

	if b[y][x] == 1 {
		if N < 2 {
			u[y][x] = 0
		} else if N > 3 {
			u[y][x] = 0
		}
	} else if b[y][x] == 0 {
		if N == 3 {
			u[y][x] = 1
		}
	}
}

func printBoard(b [][]int) {
	fmt.Println()
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
			//b = append(b, r.Intn(2))

			rn := r.Intn(100)
			if rn > 50 {
				b = append(b, 1)
			} else {
				b = append(b, 0)
			}

		}

		a = append(a, b)
		b = []int{}
	}

	return a
}
