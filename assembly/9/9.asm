assume cs:code

data segment
    db 'welcome to masm!'
data ends

code segment
    start: mov ax,data
           mov ds,ax
	   mov cx,16
	   mov bx,1824
	   mov si,0
	   mov ax,0B800h
	   mov es,ax
	
	s: mov al,[si]
	   mov es:[bx+si],al
	   inc bx
	   inc si
	   loop s

	   mov ax,4c00h
	   int 21h

code ends

end start
