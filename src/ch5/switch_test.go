package ch5

import (
	"fmt"
	"runtime"
	"testing"
)

/**
case 不需要break,
 和java 一样，不限制类型
*/
func TestSwitch(t *testing.T) {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	default:
		fmt.Print("default")
	}
}
