package pkg

type Piece struct {
	Position *SquarePosition
	Color    Color
	Type     PieceType
}

type Color int

const (
	White Color = iota
	Black
)

type PieceType int

const (
	Pawn   PieceType = iota
	Knight
	Bishop
	Rook
	Queen
	King
)
