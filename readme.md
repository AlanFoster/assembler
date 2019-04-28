# Assembler

A basic assembler for the Hack symbolic assembly language written in Go.

## What is it?

Following the through the book `The Elements of Computing Systems`, this project is part of a larger suite of tools:

       +------------+     +-------------+     +------------+
       |    Jack    |     |   Virtual   |     |            |
       |  Compiler  +---->+   Machine   +---->+  Assembler |
       |            |     |  Translator |     |            |
       +------------+     +-------------+     +------------+

### Jack Compiler

The [Jack Compiler](https://github.com/AlanFoster/jackcompiler) converts the high level Jack language into an
intermediate representation which can be ran on a platform agnostic virtual machine. The syntax of Jack is similar to
Java. The virtual machine instructions output by this compiler are stack based, and is modeled after the Java Virtual
Machine (JVM).

This project is written with [ANTLR](https://www.antlr.org/) and Python.

### Hack Virtual Machine Translator

The Jack Compiler outputs virtual machine code. These virtual machine instructions can be compiled using the
[hack virtual machine translator](https://github.com/AlanFoster/hackvirtualmachine). At a high level this tool converts
virtual machine instructions into symbolic assembly commands which can then be passed to an assembler.

This project is written with [ANTLR](https://www.antlr.org/) and Python.

### Hack Assembler

The [Hack Assembler](https://github.com/AlanFoster/hackassembler) takes the symbolic representation of assembly
commands, and converts these instructions into its binary representation using the Hack Assembler. This can then be
loaded on to the Hack platform's ROM and executed.

This project is written with Go.

## Project links

This project is part of a larger suite of tools:

- [Jack Compiler](https://github.com/AlanFoster/jackcompiler) converts the high level Jack language into an
  intermediate representation which can be ran on a platform agnostic virtual machine. Written with [Antlr](https://www.antlr.org/)
  and Python.
- [Hack Virtual Machine Translator](https://github.com/AlanFoster/hackvirtualmachine) - A virtual machine translator
  for the hack assembly language written with [Antlr](https://www.antlr.org/) and Python.
- [Hack Assembler](https://github.com/AlanFoster/hackassembler) - A basic assembler for the Hack symbolic assembly
  language written in Go.

## Running

> go run main.go --entry-file ./your-file.asm --output-file ./your-file.hack

## Language

This assembler converts the symbolic assembly commands into its binary representation.
The width of each instruction is 16 bits, and the source machine only supports two instructions.

### A Instruction

The 'A' instruction can be used to hold either a data value, or be used as a reference to a memory location.
The 'A' Instruction is represented in assembly as:

```
@1337
```

This will load the value 1337 into the A register. The binary representation for this is:

```
0 0000010100111001
^ ^^^^^^^^^^^^^^^^
|        |
|        +- 15-bit representation of 1337
|
+-------- Representation for the 'A' Instruction
```

The assembler also supports symbolic variables, and will be allocated the next free word in memory:

```
@myVariable     // Create a new variable
M = 0           // Assign the value 0
```

### C Instruction

The 'C' Instruction is more complex, and can perform Addition, Subtraction, and conditional Jumps.
Represented in its symbolic representation:

```
D=D-M;JLT
```

This will assign to the D register, the value of D - the current Memory value referenced by register A.
If the resulting value is less than 0, it will jump.

The binary representation for this:

```
111 1010011 010 100
^^^ ^^^^^^^ ^^^ ^^^
|      |     |   |
|      |     |   +- Representation for Jumping to the instruction addressed by the A register
|      |     |      If the result of D-M is Less Than 0
|      |     |
|      |     +- Representation to store result into the D register
|      |
|      +- Representation for the command D-M
|
+------ Representation for the 'C' instruction
```

### L Instruction

The assembler also supports a pseudo instruction that is reserved for labels. Labels are useful as they can be
used to to jump specific program locations:

```
(LOOP)          // Mark the next intruction as being called 'LOOP'
    @LOOP       // Load the ROM location of LOOP in to our register
    0; JMP      // Unconditionally Jump. i.e. Infinitely loop.
```

The L Instruction does not emit any binary, instead the assembler will inline the ROM locations of the target
instruction

## Implementation

At a high level the implementation is:

```
lexer -> Parser -> Generator
```

This approach isn't specifically needed to implement an assembler, but the intermediate representation definitely
made the code generation nicer.

Labels are handled via a first initial pass of the assembly code, and the mapping of labels to ROM locations is stored
in a symbol table. After the ROM location of labels has been identified, the a secondary pass is used to generate
the real binary representation of our symbol assembly code.

## Example

### Input File

Given an assembly program that adds numbers:

```
// Computes R2 = max(R0, R1)  (R0,R1,R2 refer to RAM[0],RAM[1],RAM[2])
   @R0
   D=M              // D = first number
   @R1
   D=D-M            // D = first number - second number
   @OUTPUT_FIRST
   D;JGT            // if D>0 (first is greater) goto output_first
   @R1
   D=M              // D = second number
   @OUTPUT_D
   0;JMP            // goto output_d
(OUTPUT_FIRST)
   @R0
   D=M              // D = first number
(OUTPUT_D)
   @R2
   M=D              // M[2] = D (greatest number)
(INFINITE_LOOP)
   @INFINITE_LOOP
   0;JMP            // infinite loop
```

## Output File

A string representing the binary representation of the input:

```
0000000000000000
1111110000010000
0000000000000001
1111010011010000
0000000000001010
1110001100000001
0000000000000001
1111110000010000
0000000000001100
1110101010000111
0000000000000000
1111110000010000
0000000000000010
1110001100001000
0000000000001110
1110101010000111
```
