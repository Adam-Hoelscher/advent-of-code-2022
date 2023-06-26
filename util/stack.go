package util

type ByteStack []byte

func (s ByteStack) Len() int {
	return len(s)
}

func (s ByteStack) Peek() byte {
	return s[len(s)-1]
}

func (s *ByteStack) Pop() byte {
	x := s.Peek()
	*s = (*s)[:len(*s)-1]
	return x
}

func (s *ByteStack) Push(x byte) {
	*s = append(*s, x)
}
