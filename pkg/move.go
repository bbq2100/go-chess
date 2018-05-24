package pkg

import "strconv"

type move struct {
	from string
	to   string
}

var (
	files = map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
		"f": 6,
		"g": 7,
		"h": 8,
	}
)

func (m *move) toSquare() (from, to square, err error) {
	do := func(s string) (file int, rank int, err error) {
		f, e := strconv.Atoi(s[1:2])
		if e != nil {
			return 0, 0, e
		}
		return f, files[s[0:1]], nil
	}
	fileFrom, rankFrom, e := do(m.from)
	fileTo, rankTo, e := do(m.to)

	if e != nil {
		return square{}, square{}, e
	}

	return square{fileFrom, rankFrom}, square{fileTo, rankTo}, nil
}
