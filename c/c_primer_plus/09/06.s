	.section	__TEXT,__text,regular,pure_instructions
	.build_version macos, 10, 14
	.globl	_main                   ## -- Begin function main
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
	movl	$1, %edi
	movl	$0, -4(%rbp)
	callq	_up_and_down
	xorl	%eax, %eax
	addq	$16, %rsp
	popq	%rbp
	retq
	.cfi_endproc
                                        ## -- End function
	.globl	_up_and_down            ## -- Begin function up_and_down
	.p2align	4, 0x90
_up_and_down:                           ## @up_and_down
	.cfi_startproc
## %bb.0:
	pushq	%rbp
	.cfi_def_cfa_offset 16
	.cfi_offset %rbp, -16
	movq	%rsp, %rbp
	.cfi_def_cfa_register %rbp
	subq	$16, %rsp
	leaq	L_.str(%rip), %rax
	leaq	-4(%rbp), %rdx
	movl	%edi, -4(%rbp)
	movl	-4(%rbp), %esi
	movq	%rax, %rdi
	movb	$0, %al
	callq	_printf
	cmpl	$4, -4(%rbp)
	movl	%eax, -8(%rbp)          ## 4-byte Spill
	jge	LBB1_2
## %bb.1:
	movl	-4(%rbp), %eax
	addl	$1, %eax
	movl	%eax, %edi
	callq	_up_and_down
LBB1_2:
	leaq	L_.str.1(%rip), %rdi
	leaq	-4(%rbp), %rdx
	movl	-4(%rbp), %esi
	movb	$0, %al
	callq	_printf
	movl	%eax, -12(%rbp)         ## 4-byte Spill
	addq	$16, %rsp
	popq	%rbp
	retq
	.cfi_endproc
                                        ## -- End function
	.section	__TEXT,__cstring,cstring_literals
L_.str:                                 ## @.str
	.asciz	"Level %d: n location %p\n"

L_.str.1:                               ## @.str.1
	.asciz	"LEVEL %d: n location %p\n"


.subsections_via_symbols
