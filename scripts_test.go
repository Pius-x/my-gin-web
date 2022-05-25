package main_test

import (
	"fmt"
	"testing"
)

func TestGo(t *testing.T) {

	m := make(map[string]int)
	v := m["key1"]
	fmt.Printf("c %+v \n", v)
}
