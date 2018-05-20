package pkg

type squarePosition struct {
	File, Rank int
}

func (s square) isSquareOccupied() bool {
	return s.piece != nil
}

type square struct {
	position squarePosition
	piece    *piece
}
