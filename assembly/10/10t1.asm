assume cs:code,ss:stack

stack segment
    db 8 dup (0) 
stack ends

data segment
    db 'Welcome to masm!',0
data ends

code segment
    start: mov dh,8
           mov dl,3
	   mov cl,2
	   mov ax,data
	   mov ds,ax
	   mov ax,stack
	   mov ss,ax
	   mov sp,16
	   mov si,0
	   call show_str

	   mov ax,4c00h
	   int 21h
 show_str: push ax
           push si
	   push di
	   push es
           push cx
	   push bx
           push dx

	   mov bx,0
           dec dh
	   mov al,80
	   mul dh
	   add bx,ax
	   mov dh,0
	   add bx,dx

	   mov ax,0b800h
           mov si,0
           mov di,0
	   mov es,ax
           mov al,cl
	   
           
     char: mov cl,[si]
           mov ch,0
	   jcxz ok
	   mov ch,al
	   mov es:[bx+di],cx
	   inc si
	   add di,2
           jmp char

       ok: pop dx
           pop bx
           pop cx
           pop es
	   pop di
	   pop si
	   pop ax
           ret


code ends
end start
