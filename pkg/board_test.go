package pkg

import (
	"testing"
)

func TestInitBoardConfiguration(t *testing.T) {
	b := newBoard()
	expectedBoardAsFan := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w"

	if !(b.asFen() == expectedBoardAsFan) {
		t.Fail()
	}
}

func TestMakeMove(t *testing.T) {
	expectedBoard := "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b"
	b := newBoard()
	b.makeMove(&move{"e2", "e4"})

	if !(b.asFen() == expectedBoard) {
		t.Fail()
	}
}

func TestIllegalMove(t *testing.T) {
	b := newBoard()
	e := b.makeMove(&move{"e0", "e4"})

	if e == nil {
		t.Fail()
	}
}
