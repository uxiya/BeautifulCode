;hello asm
;linux汇编hello world!
section .data
		msg db "hello world!"
		len equ $ - msg
section .text
		global _start
_start:
		mov edx, len
		mov ecx, msg
		mov ebx, 1
		mov eax, 4
		init 0x80

		mov ebx, 0
		mov eax, 1
		init 0x80
