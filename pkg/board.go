package pkg

import (
	"bytes"
	"strconv"
)

type board struct {
	pieces            []*piece
	moveCountTotal    int
	whosTurn          color
	squares           [8][8]*square
	positionWhiteKing *squarePosition
	positionBlackKing *squarePosition
}

func newBoard() board {
	squares, pieces := initBoard()
	return board{moveCountTotal: 0, whosTurn: white, pieces: pieces, squares: squares}
}

func initBoard() (squares [8][8]*square, pieces []*piece) {
	emptySquare := ""

	ss := [8][8]*square{}
	ps := make([]*piece, 0)

	for i, e := range [][]struct {
		file      int
		rank      int
		pieceType string
	}{
		{{1, 1, wRook}, {1, 2, wKnight}, {1, 3, wBishop}, {1, 4, wQueen}, {1, 5, wKing}, {1, 6, wBishop}, {1, 7, wKnight}, {1, 8, wRook}},
		{{2, 1, wPawn}, {2, 2, wPawn}, {2, 3, wPawn}, {2, 4, wPawn}, {2, 5, wPawn}, {2, 6, wPawn}, {2, 7, wPawn}, {2, 8, wPawn}},
		{{3, 1, emptySquare}, {3, 2, emptySquare}, {3, 3, emptySquare}, {3, 4, emptySquare}, {3, 5, emptySquare}, {3, 6, emptySquare}, {3, 7, emptySquare}, {3, 8, emptySquare}},
		{{4, 1, emptySquare}, {4, 2, emptySquare}, {4, 3, emptySquare}, {4, 4, emptySquare}, {4, 5, emptySquare}, {4, 6, emptySquare}, {4, 7, emptySquare}, {4, 8, emptySquare}},
		{{5, 1, emptySquare}, {5, 2, emptySquare}, {5, 3, emptySquare}, {5, 4, emptySquare}, {5, 5, emptySquare}, {5, 6, emptySquare}, {5, 7, emptySquare}, {5, 8, emptySquare}},
		{{6, 1, emptySquare}, {6, 2, emptySquare}, {6, 3, emptySquare}, {6, 4, emptySquare}, {6, 5, emptySquare}, {6, 6, emptySquare}, {6, 7, emptySquare}, {6, 8, emptySquare}},
		{{7, 1, bPawn}, {7, 2, bPawn}, {7, 3, bPawn}, {7, 4, bPawn}, {7, 5, bPawn}, {7, 6, bPawn}, {7, 7, bPawn}, {7, 8, bPawn}},
		{{8, 1, bRook}, {8, 2, bKnight}, {8, 3, bBishop}, {8, 4, bQueen}, {8, 5, bKing}, {8, 6, bBishop}, {8, 7, bKnight}, {8, 8, bRook}},
	} {
		for ii := 0; ii < 8; ii++ {
			efs := e[ii]
			pos := squarePosition{efs.file, efs.rank}
			ffs := &piece{&pos, e[ii].pieceType}
			ss[i][ii] = &square{pos, ffs}
			ps = append(ps, ffs)
		}
	}

	return ss, ps
}

func (b board) asFen() string {
	var buffer bytes.Buffer

	// Piece positions
	for file := 7; file >= 0; file-- {
		emptySquares := 0
		for rank := 0; rank < 8; rank++ {
			p := b.squares[file][rank]
			if p == nil || p.piece.pieceType == "" {
				emptySquares++
				if rank == 7 {
					buffer.WriteString(strconv.Itoa(emptySquares))
					buffer.WriteString("/")
				}
				continue
			}
			if emptySquares > 0 {
				buffer.WriteString(strconv.Itoa(emptySquares))
				emptySquares = 0
			}
			buffer.WriteString(p.piece.pieceType)
			if rank == 7 && file > 0 {
				buffer.WriteString("/")
			}
		}
	}

	buffer.WriteString(" ")

	// Active color
	switch b.whosTurn {
	case white:
		buffer.WriteString("w")
	case black:
		buffer.WriteString("b")
	}

	return buffer.String()
}
