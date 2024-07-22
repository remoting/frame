package cpuid

import (
	"fmt"
	"testing"
)

func TestMachine001(t *testing.T) {
	a, _ := GenerateMachineCode()
	fmt.Println(a)
}
