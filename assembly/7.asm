assume cs:codesg, ss:stack, ds:data, es:table

stack segment
dw 8 dup (0)
stack ends

data segment
db '1975', '1976', '1977', '1978', '1979', '1980', '1981', '1982', '1983'
db '1984', '1985', '1986', '1987', '1988', '1989', '1990', '1991', '1992'
db '1993', '1994', '1995'

dd 16, 22, 382, 1356, 2390, 8000, 16000, 24486, 50065, 97479, 140417, 197514
dd 345980, 590827, 803530, 1183000, 1843000, 2759000, 3753000, 4649000, 5937000

dw 3, 7, 9, 13, 28, 38, 130, 220, 476, 778, 1001, 1442, 2258, 2793, 4037, 5635, 8226
dw 11542, 14430, 15257, 17800
data ends

table segment
db 21 dup ('year summ ne ?? ')
table ends

codesg segment

     start: mov ax,data
            mov ds,ax
            mov ax,table
	    mov es,ax
            mov ax,stack
	    mov ss,ax
	    mov sp,16
            mov cx,21
	    
	    mov si,0
	    mov di,0
	    mov bp,0
            mov bx,0
	s0: push ds:[bp+168]
	    push [bx+di+86]
	    push [bx+di+84]
	    push [bx+di+2]
	    push [bx+di]
	    
            pop es:[bx+si]
            pop es:[bx+si+2]
            pop es:[bx+si+5]
            pop es:[bx+si+7]
            pop es:[bx+si+10]
	    mov dx,es:[bx+si+7]
	    mov ax,es:[bx+si+5]
	    div word ptr es:[bx+si+10]
	    mov es:[bx+si+13],ax
            
	    add si,16
	    add di,4
	    add bp,2
	    loop s0

	    mov ax,4c00h
	    int 21h
codesg ends

end start
