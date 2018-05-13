package pkg

type SquarePosition struct {
	File, Rank int
}

type PieceOnSquare struct {
	Position *SquarePosition
	Piece    *Piece
}

func (p PieceOnSquare) IsSquareOccupied() bool {
	return p.Piece != nil
}

type Square struct {
	PieceOnSquare *SquarePosition
}
