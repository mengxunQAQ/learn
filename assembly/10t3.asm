assume cs:code

stack segment
    dw 16 dup(0)
stack ends

data segment
    dw 8 dup (0)
data ends

code segment
    start: mov ax,12666
           mov bx,data
	   mov ds,bx
	   mov bx,stack
	   mov ss,bx
	   mov sp,32 
	   mov si,0
	   mov di,0
	   call dtoc

	   mov ax,4c00h
	   int 21h

     dtoc: push bx
           push cx
	   push dx
	   push si
	   push ds
	   push di

           mov bx,10
           mov cx,1

 push_stack:
           jcxz pop_stack
	   mov dx,0
	   div bx
	   mov cl,al
	   mov ch,0
	   add dl,30h
           push dx
	   inc si
	   jmp short push_stack
	   
	   mov cx,si
 pop_stack: 
           pop bx
	   mov [di],bl
           inc di
	   loop pop_stack

       ok: pop di
           pop ds
           pop si
	   pop dx
	   pop cx
	   pop bx
           ret

code ends

end start

