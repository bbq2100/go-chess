package pkg

import (
	"testing"
)

const initialPosition = "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w"

func TestInitBoardConfiguration(t *testing.T) {
	b := newBoard()
	expectedBoardAsFan := initialPosition

	if !(b.asFen() == expectedBoardAsFan) {
		t.Fail()
	}
}

func TestMakeMove(t *testing.T) {
	expectedBoard := "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b"
	b := newBoard()
	b.makeMove(&move{"e2", "e4"})

	if b.moveCountTotal != 1 {
		t.Fail()
	}

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

func TestUndoMove(t *testing.T) {
	expectedBoardAfterMove := "rnbqkbnr/pppppppp/8/8/4P3/8/PPPP1PPP/RNBQKBNR b"
	b := newBoard()
	b.makeMove(&move{"e2", "e4"})

	if !(b.asFen() == expectedBoardAfterMove) {
		t.Fail()
	}

	expectedBoardAfterUndoMove := initialPosition
	b.undoMove()

	if !(b.asFen() == expectedBoardAfterUndoMove) {
		t.Fail()
	}
}

func TestUndoMoveOnEmptyHistory(t *testing.T) {
	b := newBoard()
	b.makeMove(&move{"e2", "e4"})
	b.undoMove()
	e := b.undoMove()
	if e != nil {
		t.Fail()
	}
}
