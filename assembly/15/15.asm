assume cs:code

stack segment
    db 128 dup (0)
stack ends

code segment
    start:  mov ax, stack ;构建栈
            mov ss, ax
            mov sp, 128
            
            push cs
            pop ds
            
            mov ax, 0
            mov es, ax
            
            mov si, offset int9
            mov di, 204h
            mov cx, offset int9end - offset int9 ;安装中断例程到内存空间
            cld ;df标志位置0
            rep movsb ;安装中断例程 
            
            push es: [9 * 4] ;保存原地址
            pop es: [200h]
            push es: [9 * 4 + 2]
            pop es: [202h]
            
            cli ;禁止中断，否则键盘中断例程的地址会错误
            mov word ptr es: [9 * 4], 204h ;写入新地址
            mov word ptr es: [9 * 4 + 2], 0
            sti ;放开中断
            
            mov ax, 4c00h
            int 21h
            
    int9:   push ax
            push bx
            push cx
            push es
    
            in al, 60h ;从端口地址空间写入
            
            pushf
            call word ptr cs:[200h] ;执行 int 9 原中断例程
            
            cmp al, 9eh ;A的断码是9eh
            jne int9ret ;一旦有断码就往下执行，否则就返回
            
            mov ax, 0b800h
            mov es, ax
            mov bx, 0
            mov cx, 2000
    s:      mov byte ptr es:[bx], 'A'
            add bx, 2
            loop s
            
    int9ret:pop es
            pop cx
            pop bx
            pop ax
            iret
    
    int9end:nop
    
code ends

end start
