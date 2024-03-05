package _11_runtime

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestCpuNums(t *testing.T) {
	t.Log(getCpuNums())
}

func TestName(t *testing.T) {
	fmt.Println(utf8.RuneCountInString("abc"))
	fmt.Println(utf8.RuneCountInString("abc诶比谁"))
}

func TestDiv(t *testing.T) {
	n := int(179)
	n1 := n / 100
	n2 := n / 100.0
	fmt.Println(n1, n2)
}
