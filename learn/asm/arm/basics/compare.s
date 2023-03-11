.EQU SYS_WRITE, 64
.EQU SYS_EXIT, 93

.global _start

_start:
  MOV X0, #5
  MOV X1, #4
  ADD X2, X0, X1
  CMP X2, #10
  BEQ equals_ten
  // if the value in register 2 does not equal 10, branch to not_equals_ten
  B not_equals_ten

not_equals_ten:
  // Write "The value is not equal to 10" to standard out
  MOV X0, #1
  LDR X1, =not_equal_message
  MOV X2, #28
  BL write_message
  B exit_program

equals_ten:
  // Write "The value equals 10" to standard out
  MOV X0, #1
  LDR X1, =equal_message
  MOV X2, #24
  BL write_message
  B exit_program

write_message:
  // Write the message to standard out
  MOV X8, SYS_WRITE
  SVC #0
  RET

exit_program:
  // Exit the program
  MOV X0, #0
  MOV X8, SYS_EXIT
  SVC #0

.section .rodata
  equal_message:
    .ascii "The value is equal to 10\n"
  not_equal_message:
    .ascii "The value is not equal to 10\n"

