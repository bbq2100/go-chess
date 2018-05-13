package pkg

type Board struct {
	Pieces         []*Piece
	MoveCountTotal int
	WhosTurn       Color
	Squares        [8][8]*Square
	PositionWhiteKing *SquarePosition
	PositionBlackKing *SquarePosition
}

func NewBoard() Board {
	return Board{MoveCountTotal: 0, WhosTurn: White, Pieces: initPiecesList(), Squares:initSquares()}
}

func initSquares() [8][8]*Square {
	return [8][8]*Square{}
}

func initPiecesList() []*Piece {
	return []*Piece{}
}
