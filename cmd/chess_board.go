package main

import "fmt"

type Board struct {
	Pieces []*Piece
}

type Piece struct {
	Position Square
}

type Square struct {
	X, Y int
}

func main() {
	fmt.Println("lol")
}
