package classics

type Status struct {
	num []int
	m   map[string]string
}

type IFileStatus interface {
	lock()
}

func (s *Status) GetNum(index int) int {
	return s.num[index]
}
