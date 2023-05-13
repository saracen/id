#include "textflag.h"

TEXT libc_gethostuuid_trampoline<>(SB), NOSPLIT, $0-0
	JMP libc_gethostuuid(SB)

GLOBL ·libc_gethostuuid_trampoline_addr(SB), RODATA, $8
DATA ·libc_gethostuuid_trampoline_addr(SB)/8, $libc_gethostuuid_trampoline<>(SB)
