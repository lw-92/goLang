package ch7

import "testing"

/**
map的相关的测试
*/
func TestInit(t *testing.T) {
	m := map[string]int{"one": 1, "two": 2}
	t.Log(m, len(m))
	m2 := make(map[string]int, 10) // 这里len 还是0，因为没有放值
	t.Log(m2, len(m2))
	m3 := map[string]int{}
	m3["four"] = 4
	t.Log(m3, len(m3))

}
