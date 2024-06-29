package util

import (
	"fmt"
	"testing"
)

func TestUUID001(t *testing.T) {
	a := NewUUID()
	fmt.Printf("%s\n ", a)
	x := GetTime(a)
	fmt.Printf("%s\n ", x)
}
