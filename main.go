package main

import "fmt"

type board struct {
	width int
	height int
}

func main() {

	b := new(board)
	b.deadState(6,5)
	fmt.Println(b)
}

func (b *board) deadState(w, h int ) {
	b.width = w
	b.height = h
}

