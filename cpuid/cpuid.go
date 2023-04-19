package cpuid

func asmCpuid(op uint32) (eax, ebx, ecx, edx uint32)

func CPUID(op uint32) (eax, ebx, ecx, edx uint32) {
	return asmCpuid(op)
}
