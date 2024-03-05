package chapter05

import (
	"fmt"
	"testing"
)

func TestPractice(t *testing.T) {
	com1 := NewComputer(&CoreI7{}, &InterMemory{}, &InterCard{})
	com1.Work()

	fmt.Println("========")

	com2 := NewComputer(&CoreI3{}, &KingstonMemory{}, &NvidiaCard{})
	com2.Work()
}
