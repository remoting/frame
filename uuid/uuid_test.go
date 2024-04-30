package uuid

import (
	"fmt"
	"testing"
)

func TestUUID001(t *testing.T) {
	for i := 0; i < 100; i++ {
		a := NewUUID()
		fmt.Printf("%s\n ", a)
		x, _ := GetUUIDTime(a)
		fmt.Printf("%s\n ", x)
	}
}
