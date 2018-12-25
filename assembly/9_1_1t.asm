assume cs:code,ds:data

data segment
    dw 0,0
data ends

code segment
    start: mov ax,data
	       mov ds,ax
		   mov bx,0
		   jmp word ptr [bx+18]
code ends
end start
