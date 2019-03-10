	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 10, 14
	.section	__TEXT,__literal8,8byte_literals
	.p2align	3               ## -- Begin function main
LCPI0_0:
	.quad	4613937818241073152     ## double 3
LCPI0_1:
	.quad	4617315517961601024     ## double 5
	.section	__TEXT,__text,regular,pure_instructions
	.globl	_main
	.p2align	4, 0x90
_main:                                  ## @main
	.cfi_startproc
## %bb.0:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register %rbp
	subq	$16, %rsp
	movl	$3, %edi
	movl	$0, -4(%rbp)
	movb	$0, %al
	callq	_imax
	leaq	L_.str(%rip), %rdi
	movl	$3, %esi
	movl	$5, %edx
	movl	%eax, %ecx
	movb	$0, %al
	callq	_printf
	movsd	LCPI0_0(%rip), %xmm0    ## xmm0 = mem[0],zero
	movsd	LCPI0_1(%rip), %xmm1    ## xmm1 = mem[0],zero
	movl	%eax, -8(%rbp)          ## 4-byte Spill
	movb	$2, %al
	callq	_imax
	leaq	L_.str(%rip), %rdi
	movl	$3, %esi
	movl	$5, %edx
	movl	%eax, %ecx
	movb	$0, %al
	callq	_printf
	xorl	%ecx, %ecx
	movl	%eax, -12(%rbp)         ## 4-byte Spill
	movl	%ecx, %eax
	addq	$16, %rsp
	popq	%rbp
	retq
	.cfi_endproc
                                        ## -- End function
	.globl	_imax                   ## -- Begin function imax
	.p2align	4, 0x90
_imax:                                  ## @imax
	.cfi_startproc
## %bb.0:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register %rbp
	subq	$16, %rsp
	leaq	L_.str.1(%rip), %rax
	movl	%edi, -4(%rbp)
	movl	%esi, -8(%rbp)
	movl	-4(%rbp), %esi
	movl	-8(%rbp), %edx
	movq	%rax, %rdi
	movb	$0, %al
	callq	_printf
	movl	-4(%rbp), %edx
	cmpl	-8(%rbp), %edx
	movl	%eax, -12(%rbp)         ## 4-byte Spill
	jle	LBB1_2
## %bb.1:
	movl	-4(%rbp), %eax
	movl	%eax, -16(%rbp)         ## 4-byte Spill
	jmp	LBB1_3
LBB1_2:
	movl	-8(%rbp), %eax
	movl	%eax, -16(%rbp)         ## 4-byte Spill
LBB1_3:
	movl	-16(%rbp), %eax         ## 4-byte Reload
	addq	$16, %rsp
	popq	%rbp
	retq
	.cfi_endproc
                                        ## -- End function
	.section	__TEXT,__cstring,cstring_literals
L_.str:                                 ## @.str
	.asciz	"The maximum of %d and %d is %d.\n"

L_.str.1:                               ## @.str.1
	.asciz	"%d %d\n"


.subsections_via_symbols
