package ch13

import (
	"fmt"
	"testing"
)

/**
go 中多态的支持
*/

/**
接口变量
、
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

func writeFisrtProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

/**
多肽测试
*/
func TestPolymorphism(t *testing.T) {

	goP := new(GoProgrammer)

	javaP := new(JavaProgrammer)
	writeFisrtProgram(goP)
	writeFisrtProgram(javaP)
}
