package main

import (
	"strconv"
	"fmt"
	"bytes"
	"strings"
)

const (
	BinaryBoardDefault = "0000000000000000000000000000000000000000000000000000000000000000"
)

type long = uint64

var (
	WP, WN, WB, WR, WQ, WK, BP, BN, BB, BR, BQ, BK long = 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
	ChessBoardInit                                      = [][]string{
		{"r", "n", "b", "q", "k", "b", "n", "r"},
		{"p", "p", "p", "p", "p", "p", "p", "p"},
		{"-", "-", "-", "-", "-", "-", "-", "-"},
		{"-", "-", "-", "-", "-", "-", "-", "-"},
		{"-", "-", "-", "-", "-", "-", "-", "-"},
		{"-", "-", "-", "-", "-", "-", "-", "-"},
		{"P", "P", "P", "P", "P", "P", "P", "P"},
		{"R", "N", "B", "Q", "K", "B", "N", "R"},}
)

func initPosition() {
	for i := 0; i < 64; i++ {
		binaryPosition := add1ToPosition(BinaryBoardDefault, i)
		switch ChessBoardInit[i/8][i%8] {
		case "P":
			WP += valueOf(binaryPosition)
		case "N":
			WN += valueOf(binaryPosition)
		case "B":
			WB += valueOf(binaryPosition)
		case "R":
			WR += valueOf(binaryPosition)
		case "Q":
			WQ += valueOf(binaryPosition)
		case "K":
			WK += valueOf(binaryPosition)
		case "p":
			BP += valueOf(binaryPosition)
		case "n":
			BN += valueOf(binaryPosition)
		case "b":
			BB += valueOf(binaryPosition)
		case "r":
			BR += valueOf(binaryPosition)
		case "q":
			BQ += valueOf(binaryPosition)
		case "k":
			BK += valueOf(binaryPosition)
		}
	}
}

func valueOf(position string) (long) {
	v, _ := strconv.ParseUint(position, 2, 64)
	return v
}

func drawBoard() string {
	board := [8][8]string{}

	for id := 0; id < 64; id++ {
		board[id/8][id%8] = "-"
	}

	var i long = 0
	for ; i < 64; i++ {
		switch {
		case ((WP >> i) & 1) == 1:
			board[i/8][i%8] = "P"
		case ((WN >> i) & 1) == 1:
			board[i/8][i%8] = "N"
		case ((WB >> i) & 1) == 1:
			board[i/8][i%8] = "B"
		case ((WR >> i) & 1) == 1:
			board[i/8][i%8] = "R"
		case ((WQ >> i) & 1) == 1:
			board[i/8][i%8] = "Q"
		case ((WK >> i) & 1) == 1:
			board[i/8][i%8] = "K"
		case ((BP >> i) & 1) == 1:
			board[i/8][i%8] = "p"
		case ((BN >> i) & 1) == 1:
			board[i/8][i%8] = "n"
		case ((BB >> i) & 1) == 1:
			board[i/8][i%8] = "b"
		case ((BR >> i) & 1) == 1:
			board[i/8][i%8] = "r"
		case ((BQ >> i) & 1) == 1:
			board[i/8][i%8] = "q"
		case ((BK >> i) & 1) == 1:
			board[i/8][i%8] = "k"
		}
	}
	var buffer bytes.Buffer
	for i := 0; i < 8; i++ {
		ff := board[i]
		buffer.WriteString(strings.Join(ff[:], " ") + "\n")
	}
	return buffer.String()
}

func add1ToPosition(in string, i int) string {
	runes := []rune(in)
	start := string(runes[i+1:])
	end := string(runes[0:i])
	return start + "1" + end
}

func main() {
	initPosition()
	fmt.Print(drawBoard())
}
