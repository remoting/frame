package cpuid

import (
	"fmt"
	"testing"
)

func TestMachine001(t *testing.T) {
	//7a84c3127b999e8243a508eb3a10c732
	a, _ := GenerateMachineCode()
	fmt.Println(a)
}
