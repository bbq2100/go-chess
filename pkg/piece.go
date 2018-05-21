package pkg

type piece struct {
	position   square
	pieceType  string
	color      color
	canCastle  bool
	isCaptured bool
}

type color int

const (
	white color = iota
	black
)

const (
	pawn   = "p"
	knight = "n"
	bishop = "b"
	rook   = "r"
	queen  = "q"
	king   = "k"
)
