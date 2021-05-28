package ch12

import "testing"

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
