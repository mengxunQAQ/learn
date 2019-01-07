assume cs:code, ds:data

data segment
    s db 9, 8, 7, 4, 2, 0
data ends

code segment
    start:  mov ax, 0b800h
            mov es, ax
            mov di, 160 * 12
            mov ax, data
            mov ds, ax
            mov si, 0
            mov cx, 6
    
    print:  mov al, s[si]
            out 70h, al
            in al, 71h ;通过端口拉取芯片里的时间
            call number
            cmp si, 2 ;决定分隔符是/还是:
            jb slash ;小于则跳转
            je space ;等于则跳转
            cmp si, 5
            jb colon ;小于则跳转
    next:   inc si
            loop print
            
            mov ax, 4c00h
            int 21h
            
    number: push cx
            mov ah, al
            mov cl, 4
            shr ah, cl ;十位 个位分离
            and al, 00001111b  
            add ah, 30h ;转化ascii码
            add al, 30h
            mov byte ptr es:[di], ah ;移存到内存
            mov byte ptr es:[di + 2], al
            add di, 4 ;偏移地址增加
            pop cx
            ret
    
    slash:  mov byte ptr es:[di], '\'
            add di, 2
            jmp next
            
    colon:  mov byte ptr es:[di], ':'
            add di, 2
            jmp next
    
    space:  mov byte ptr es:[di], ' '
            add di, 2
            jmp next

code ends

end start
