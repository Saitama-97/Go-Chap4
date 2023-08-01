package utils

import (
	"fmt"
	"testing"
	"unsafe"
)

/**
 * @Time    : 2023.08.01 10:27
 * @File    : integer_test.go
 * @Project : chapter4
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 测试 int 及相关整型所占字节数
 */

func TestInteger(t *testing.T) {
	var a8 = int8(1)
	var a16 = int16(1)
	var a32 = int32(1)
	var a64 = int64(1)
	var a = int(1)
	fmt.Println("a8:", unsafe.Sizeof(a8))
	fmt.Println("a16:", unsafe.Sizeof(a16))
	fmt.Println("a32:", unsafe.Sizeof(a32))
	fmt.Println("a64:", unsafe.Sizeof(a64))
	fmt.Println("a:", unsafe.Sizeof(a))
}
