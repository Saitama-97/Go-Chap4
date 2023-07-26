package socket

import (
	"reflect"
	"strings"
	"testing"
)

/**
 * @Time    : 2023/7/26 17:05
 * @File    : trim_test.go
 * @Project : chapter4
 * @Author  : Saitama
 * @IDE     : GoLand
 * @Desc    : 测试 strings.Trim() 函数
 */

func Test_Trim(t *testing.T) {
	str := "ahello worldl"
	cutSet := "hel"
	retSlice := strings.TrimLeft(str, cutSet)
	want := "llo world"
	if !reflect.DeepEqual(want, retSlice) {
		t.Errorf("want:%v, retSlice:%v", want, retSlice)
	}
}
