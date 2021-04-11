package classics


type Status struct {
	num int[]
	
}


type IFileStatus interface {
	lock()
}

func (s * Status) GetNum(index int) int {
	return s.num[index]
}
