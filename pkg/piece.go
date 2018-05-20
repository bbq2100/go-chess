package pkg

type piece struct {
	position  *squarePosition
	pieceType string
}

type color int

const (
	white color = iota
	black
)

const (
	bPawn   = "p"
	bKnight = "n"
	bBishop = "b"
	bRook   = "r"
	bQueen  = "q"
	bKing   = "k"
	wPawn   = "P"
	wKnight = "N"
	wBishop = "B"
	wRook   = "R"
	wQueen  = "Q"
	wKing   = "K"
)
