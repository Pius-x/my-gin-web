package main_test

import (
	"fmt"
	"testing"
)

func TestGo(tt *testing.T) {
	a := [3]string{"A", "B", "C"}
	b := a[1:3]
	c := a[1:3]

	fmt.Println("before a ", a)
	fmt.Println("before b ", b)
	fmt.Println("before c ", c)
	c = append(c, "2", "2", "2", "2")
	c[1] = "F"
	fmt.Println("after a ", a)
	fmt.Println("after b ", b)
	fmt.Println("after c ", c)
}
