package ch12

import (
	"fmt"
	"testing"
	"time"
)

/**\
接口变量
*/
type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (input *GoProgrammer) WriteHelloWorld() string {
	return "fmt.println(\"hello world\")"
}

type JavaProgrammer struct {
}

func (input *JavaProgrammer) WriteHelloWorld() string {
	return "sout(hello word)"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
	p = new(JavaProgrammer)
	t.Log(p.WriteHelloWorld())
}

// 自定义类型

/**
计算某一个函数的执行时间
*/
func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		now := time.Now()
		i := inner(n)
		fmt.Println("time spend:", time.Since(now).Seconds())
		return i
	}
}
