package pkg

type stack struct {
	datas []interface{}
}

func newStack() *stack {
	return &stack{}
}

func (s *stack) add(v interface{}) {
	s.datas = append(s.datas, v)
}

func (s *stack) pop() interface{} {
	l := len(s.datas)
	switch l {
	case 0:
		return nil
	default:
		i := s.datas[l-1 : l]
		s.datas = s.datas[:l-1]
		return i[0]
	}
}

func (s *stack) size() int {
	return len(s.datas)
}
