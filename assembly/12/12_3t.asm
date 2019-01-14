assume cs:code

code segment
    start: mov ax,cs
           mov ds,ax
	   mov si,offset do0
	   mov ax,0
	   mov es,ax
	   mov di,200h ;设置目的地址
	   mov cx,offset do0end- offset do0 ;设置cx的传输长度
	   cld ;传输方向正向
	   rep movsb ;循环es:di<-ds:si

	   ;设置中断向量
	   mov ax,0
	   mov es,ax
	   mov word ptr es:[0*4],200h
	   mov word ptr es:[0*4+2],0

	   mov ax,4c00h
	   int 21h
      
      do0: jmp short do0start
           db "overflow!"
 
           do0start: mov ax,cs
	             mov ds,ax
		     mov si,202h ;指向字符串overflow!的位置

		     mov ax,2000h
		     mov es,ax
		     mov di,0 ;指向显示字符串的目的地址

		     mov cx,9
		  s: mov al,[si]
		     mov es:[di],al
		     inc si
		     inc di
		     loop s

		     mov ax,4c00h
		     int 21h
             do0end: nop
code ends
end start


