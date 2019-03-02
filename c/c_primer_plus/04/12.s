	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 10, 14
	.section	__TEXT,__literal8,8byte_literals
	.p2align	3               ## -- Begin function main
LCPI0_0:
	.quad	4613937818241073152     ## double 3
	.section	__TEXT,__literal4,4byte_literals
	.p2align	2
LCPI0_1:
	.long	1077936128              ## float 3
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
	subq	$48, %rsp
	leaq	L_.str(%rip), %rdi
	movsd	LCPI0_0(%rip), %xmm0    ## xmm0 = mem[0],zero
	movss	LCPI0_1(%rip), %xmm1    ## xmm1 = mem[0],zero,zero,zero
	movl	$0, -4(%rbp)
	movss	%xmm1, -8(%rbp)
	movsd	%xmm0, -16(%rbp)
	movq	$200000000, -24(%rbp)   ## imm = 0xBEBC200
	movq	$1234567890, -32(%rbp)  ## imm = 0x499602D2
	movss	-8(%rbp), %xmm0         ## xmm0 = mem[0],zero,zero,zero
	cvtss2sd	%xmm0, %xmm0
	movsd	-16(%rbp), %xmm1        ## xmm1 = mem[0],zero
	movq	-24(%rbp), %rsi
	movq	-32(%rbp), %rdx
	movb	$2, %al
	callq	_printf
	leaq	L_.str.1(%rip), %rdi
	movq	-24(%rbp), %rsi
	movq	-32(%rbp), %rdx
	movl	%eax, -36(%rbp)         ## 4-byte Spill
	movb	$0, %al
	callq	_printf
	leaq	L_.str.2(%rip), %rdi
	movss	-8(%rbp), %xmm0         ## xmm0 = mem[0],zero,zero,zero
	cvtss2sd	%xmm0, %xmm0
	movsd	-16(%rbp), %xmm1        ## xmm1 = mem[0],zero
	movq	-24(%rbp), %rsi
	movq	-32(%rbp), %rdx
	movl	%eax, -40(%rbp)         ## 4-byte Spill
	movb	$2, %al
	callq	_printf
	xorl	%ecx, %ecx
	movl	%eax, -44(%rbp)         ## 4-byte Spill
	movl	%ecx, %eax
	addq	$48, %rsp
	popq	%rbp
	retq
	.cfi_endproc
                                        ## -- End function
	.section	__TEXT,__cstring,cstring_literals
L_.str:                                 ## @.str
	.asciz	"%.le %.le %.le %.le\n"

L_.str.1:                               ## @.str.1
	.asciz	"%ld %ld\n"

L_.str.2:                               ## @.str.2
	.asciz	"%ld %ld %ld %ld\n"


.subsections_via_symbols
