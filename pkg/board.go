package pkg

import (
	"bytes"
	"errors"
	"strconv"
	"strings"
)

type board struct {
	pieces         []*piece
	moveCountTotal int
	whosTurn       color
	history        *stack
}

func newBoard() board {
	return board{moveCountTotal: 0, whosTurn: white, pieces: initBoard(), history: newStack()}
}

func initBoard() []*piece {
	ps := make([]*piece, 0)
	pieceFiles := []int{1, 8}
	kingsRank := 5
	queensRank := 4
	bishopsRank := []int{3, 6}
	knightsRank := []int{2, 7}
	rooksRank := []int{1, 8}

	// Place kings to first position since they're needed all the time; fast access is key here
	ps = append(ps, createPiece(square{1, kingsRank}, king, white))
	ps = append(ps, createPiece(square{8, kingsRank}, king, black))

	for _, file := range pieceFiles {
		color := whichColor(file)

		ps = append(ps, createPiece(square{file, queensRank}, queen, color))

		for _, r := range bishopsRank {
			ps = append(ps, createPiece(square{file, r}, bishop, color))
		}

		for _, r := range knightsRank {
			ps = append(ps, createPiece(square{file, r}, knight, color))
		}

		for _, r := range rooksRank {
			ps = append(ps, createPiece(square{file, r}, rook, color))
		}

	}

	for _, f := range []int{2, 7} {
		for r := 1; r < 9; r++ {
			color := whichColor(f)
			ps = append(ps, createPiece(square{f, r}, pawn, color))
		}
	}

	return ps
}

func whichColor(file int) color { // color mapping for the board initial position
	switch file {
	case 1:
		return white
	case 2:
		return white
	case 7:
		return black
	case 8:
		return black
	default:
		panic("d'oh")
	}
}

func createPiece(s square, pieceType string, color color) *piece {
	return &piece{s, pieceType, color, true, false}
}

func (b *board) asFen() string {
	var buffer bytes.Buffer
	squares := [8][8]*piece{}

	for _, p := range b.pieces {
		file := p.position.file
		rank := p.position.rank
		squares[file-1][rank-1] = p
	}

	// Piece positions
	for file := 7; file >= 0; file-- {
		emptySquares := 0
		for rank := 0; rank < 8; rank++ {
			p := squares[file][rank]
			if p == nil || p.pieceType == "" {
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
			switch p.color {
			case white:
				buffer.WriteString(strings.ToUpper(p.pieceType))
			case black:
				buffer.WriteString(p.pieceType)
			}
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

func (b *board) makeMove(m *move) error {
	from, to, e := m.toSquare()

	if e != nil {
		return e
	}

	found := false

	for _, p := range b.pieces {
		if p.position == from {
			found = true
			p.position = to
			b.moveCountTotal += 1
			b.history.add(undoMove{from: from, to: to})
			switch b.whosTurn {
			case white:
				b.whosTurn = black
			case black:
				b.whosTurn = white
			}
			break
		}
	}

	if !found {
		return errors.New("'from piece' is not existing")
	}

	return nil
}

func (b *board) undoMove() error {
	history := b.history
	if history.size() > 0 {
		move1 := history.pop()
		m, ok := move1.(undoMove)

		if !ok {
			panic("Type cast error")
		}

		found := false

		for _, p := range b.pieces {
			if p.position == m.to {
				found = true
				p.position = m.from
				b.moveCountTotal -= 1
				switch b.whosTurn {
				case white:
					b.whosTurn = black
				case black:
					b.whosTurn = white
				}
				break
			}
		}

		if !found {
			return errors.New("'from piece' is not existing")
		}
	}
	return nil
}
