package mode

import "fmt"

// API hello interface
type API interface {
	Say(name string) string
}

// NewAPI API 接口的工厂函数
func NewAPI(t int) API {
	if t == 1 {
		return &HiAPI{}
	}
	return &HelloAPI{}
}

// HiAPI API 实现一
type HiAPI struct{}

// Say say hi
func (*HiAPI) Say(name string) string {
	return fmt.Sprintf("Hi %s", name)
}

// HelloAPI API 实现二
type HelloAPI struct{}

// Say hello
func (*HelloAPI) Say(name string) string {
	return fmt.Sprintf("Hello %s", name)
}
