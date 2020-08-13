package mode

import (
	"fmt"
)

/*
简单工厂函数
*/
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

/*
工厂方法
*/

// ConfigParse 定义配置文件解析接口
type ConfigParse interface {
	Parse(path string) map[string]int
}

//JSONParse 解析 json
type JSONParse struct {
}

// Parse 解析json
func (*JSONParse) Parse(path string) map[string]int {
	return map[string]int{"json": 1}
}

// XMLParse 解析 XML
type XMLParse struct {
}

// Parse 解析 xml
func (*XMLParse) Parse(path string) map[string]int {
	return map[string]int{"XML": 2}
}

// ConfigFactory 工厂函数创建接口
type ConfigFactory interface {
	Create(path string) ConfigParse
}

type XMLFactory struct {
}

func (*XMLFactory) Create(path string) XMLParse {
	return XMLParse{}
}

type JSONFactor struct {
}

func (*JSONFactor) Create(path string) JSONParse {
	return JSONParse{}
}

func NewParser(path string) ConfigParse {
	if path == ".json" {
		return JSONFactor{}.Create(path)
	}
	return XMLFactory{}.Create(path)

}

/*
创建者模式
*/

type Car struct {
	Num   int
	Color string
}

type CarBuild interface {
	SetNum(int)
	SetColor(string)
	build() Car
}

type BuildSmallCar struct {
	Num   int
	Color string
}

func (b *BuildSmallCar) SetNum(num int) {
	if num < 10 {
		b.Num = num
	}
}

func (b *BuildSmallCar) SetColor(color string) {
	b.Color = color
}
