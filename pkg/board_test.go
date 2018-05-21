package pkg

import (
	"testing"
)

func TestInitBoardConfiguration(t *testing.T) {
	b := newBoard()
	expectedBoardAsFan := "rnbqkbnr/pppppppp/8/8/8/8/PPPPPPPP/RNBQKBNR w"

	println(b.asFen())

	if !(b.asFen() == expectedBoardAsFan) {
		t.Fail()
	}
}

func TestMakeMove(t *testing.T) {
	newBoard()

}
