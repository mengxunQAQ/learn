assume cs:codesg

datasg segment
    db "Beginner's All-purpose Symbolic Instruction Code.",0
datasg ends

codesg segment
    begin: mov ax,datasg
           mov ds,ax
	   mov si,0
	   call letterc

	   mov ax,4c00h
	   int 21h

  letterc: push cx
           push si
	   pushf

	   mov ch,0
	   start: mov cl,ds:[si]
	          jcxz ok
		  cmp cl,97
		  jb next
		  cmp cl,122
		  ja next
		  
		  sub cl,20h
		  mov ds:[si],cl
		  next: inc si
	   
	   jmp start

	   ok: popf
	       pop si
	       pop cx
	       ret

codesg ends
end begin
