package collection

import "testing"

/**
切片
内部是一个结构体
 ptr 一个指针，指向内部的一个数组
len 元素的个数
cap 内部数组的容量
*/
func TestSlice(t *testing.T) {
	var s0 []int            //不需要说明长度
	t.Log(len(s0), cap(s0)) //容量和长度都是0
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0)) //容量和长度都是1
	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))
	s2 := make([]int, 3, 5)
	t.Log(len(s2), cap(s2)) //长度和容量
	//t.Log(s2[0],s2[1],s2[2],s2[3])// 只初始化了3个，不能访问下标为3的元素
	t.Log(s2[0], s2[1], s2[2]) //
	s2 = append(s2, 2)         //append 会直接加入到后面 0002
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2], s2[3])

}