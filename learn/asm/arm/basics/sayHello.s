.global _start

.section .text

_start:
  MOV X0, 1
  LDR X1, =message
  MOV X2, 12
  MOV X8, 64
  SVC 0
  MOV X0, 0
  MOV X8, 93
  SVC 0

.section .rodata
  message:
    .ascii "Hello World\n"

