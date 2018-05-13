package pkg

import (
	"testing"
	"reflect"
)

func TestInitBoardConfiguration(t *testing.T) {
	b := NewBoard()
	expectedBoard := Board{MoveCountTotal: 0, WhosTurn: White, Pieces: initPiecesList(), Squares:initSquares()}

	if !reflect.DeepEqual(b, expectedBoard) {
		t.Fail()
	}
}
