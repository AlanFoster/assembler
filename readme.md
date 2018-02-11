# Assembler

A basic assembler for the Hack symbolic assembly language written in Go.

## Running

> go run main.go --entry-file ./your-file.asm --output-file ./your-file.hack

## Language

This assembler converts the symbolic assembly commands into its binary representation.
The width of each instruction is 16 bits, and the source machine only supports two instructions.

### A Instruction

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

### C Instruction

Whilst a 'C' Instruction represented in its symbolic representation:

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
+-------- Representation for the 'C' instruction

```

## Implementation

At a high level the implementation is:

```
lexer -> Parser -> Generator
```

This approach isn't specifically needed to implement an assembler, but the intermediate representation definitely
made the code generation nicer.

## Example

### Input File

Given an assembly program that adds numbers:

```
@0
D=M
@1
D=D-M
@10
D;JGT
@1
D=M
@12
0;JMP
@0
D=M
@2
M=D
@14
0;JMP
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
