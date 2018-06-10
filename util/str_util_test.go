package util

import (
	"testing"
	"fmt"
)

func TestStrUtil(t *testing.T) {
	actual_str := StrUtil("test")
	expected_str := "test calls StrUtil\n"

	if actual_str == expected_str {
		fmt.Println("pass")
	} else {
		t.Errorf("Expected %q, actual %q", expected_str, actual_str)
	}
}
