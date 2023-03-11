.global _start

.section .text

_start:
  MOV X1, #1
  MOV X2, #2
  ADD X0, X1, X2
  MOV X8, 93
  SVC 0

