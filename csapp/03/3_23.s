	.file	"3_23.c"
	.text
	.globl	_dw_loop
	.def	_dw_loop;	.scl	2;	.type	32;	.endef
_dw_loop:
LFB0:
	.cfi_startproc
	movl	4(%esp), %eax
	movl	%eax, %ecx
	imull	%eax, %ecx
	leal	(%eax,%eax), %edx
L2:
	leal	1(%ecx,%eax), %eax
	subl	$1, %edx
	testl	%edx, %edx
	jg	L2
	rep ret
	.cfi_endproc
LFE0:
	.ident	"GCC: (MinGW.org GCC-6.3.0-1) 6.3.0"
