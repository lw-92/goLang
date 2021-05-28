package ch13

import (
	"fmt"
	"testing"
)

/**
go 中如何支持继承
*/

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Print("...")
}
func (p *Pet) SpeakTo(host string) {
	p.Speak()
	fmt.Println(" ", host)
}

/**
dog 如何继承Pet,扩展
*/
//type Dog struct {
//	p *Pet
//}
/**
匿名嵌套类型
*/
type Dog struct {
	Pet
}

func TestDog(t *testing.T) {
	d := new(Dog)
	d.SpeakTo("222")
}
