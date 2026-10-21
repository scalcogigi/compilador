section .data
  format_out: db "%d", 10, 0
  format_in: db "%d", 0
  scan_int: dd 0

section .text
  extern printf
  extern scanf
  extern exit
  global _start

_start:
  push ebp
  mov ebp, esp

  sub esp, 8

  mov dword [ebp-4], 0
  mov eax, 5
  mov [ebp-4], eax
  mov dword [ebp-8], 0
  mov eax, [ebp-4]
  push eax
  mov eax, 0
  mov ecx, eax
  pop eax
  cmp eax, ecx
  setg al
  movzx eax, al
  mov [ebp-8], eax
  mov eax, [ebp-8]
  cmp eax, 0
  je if_false_11
  mov eax, [ebp-4]
  push eax
  push format_out
  call printf
  add esp, 8
if_false_11:

  mov esp, ebp
  pop ebp

  push dword 0
  call exit
