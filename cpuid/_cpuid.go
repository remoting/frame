/*
#include <stdio.h>
#include <cpuid.h>

int get_cpuid() {
    unsigned int eax, ebx, ecx, edx;
    __get_cpuid(0, &eax, &ebx, &ecx, &edx);
    return ebx;
}
*/
import "C"
import (
	"cpuid/cpuid"
	"fmt"
)

func main() {
	fmt.Println(C.get_cpuid())

	a, b, c, d := cpuid.CPUID(0)
	fmt.Println(a, b, c, d)
}
