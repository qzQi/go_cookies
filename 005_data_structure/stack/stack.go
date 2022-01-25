package stack

type Stack struct {
	// 结构体导出的话，成员最好也导出。不然可能没法初始化
	Size     int
	ElemData []interface{}
}

func (s *Stack) Empty() bool {
	return s.Size == 0
}

func (s *Stack) S() int {
	return s.Size
}

func (s *Stack) Put(val interface{}) {
	s.Size++
	s.ElemData = append(s.ElemData, val)
}

func (s *Stack) Pop() interface{} {
	s.Size--
	v := s.ElemData[len(s.ElemData)-1]
	s.ElemData = s.ElemData[:len(s.ElemData)-1]
	return v
}
