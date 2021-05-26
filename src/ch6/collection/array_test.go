package collection

import "testing"

/**
数组的申明
*/
func TestInit(t *testing.T) {
	var arr [3]int
	t.Log(arr[0], arr[1])

	arr1 := [4]int{1, 2, 3, 4}
	t.Log(len(arr1))
	arr3 := [...]int{1, 2}
	t.Log(arr3)
	// 多为数组
	arr4 := [2][2]int{{1, 2}, {3, 4}}
	t.Log(arr4)

}

/**
数组的便利
*/
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1, 2}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}
	for idex, _ := range arr3 {
		t.Log(arr3[idex])
	}
}
