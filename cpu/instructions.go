package cpu

import "log"

type Instruction struct {
	Opcode  uint8 //one byte opcodes in 6502
	Name    string
	Length  uint16 //number of bytes in instruction (opcode + operands)
	Cycles  uint8  //number of clock cycles the instruction takes to execute
	Address uint16
}

// list of 6502 instructions (not exhaustive, just a few examples)
var instructionSet = map[uint8]Instruction{
	0xA9: {Opcode: 0xA9, Name: "LDA", Length: 2, Cycles: 2}, // Load Accumulator with Immediate
	0xA5: {Opcode: 0xA5, Name: "LDA", Length: 2, Cycles: 3}, // Load Accumulator from Zero Page
	0xAD: {Opcode: 0xAD, Name: "LDA", Length: 3, Cycles: 4}, // Load Accumulator from Absolute
	0x85: {Opcode: 0x85, Name: "STA", Length: 2, Cycles: 3}, // Store Accumulator in Zero Page
	0x8D: {Opcode: 0x8D, Name: "STA", Length: 3, Cycles: 4}, // Store Accumulator in Absolute
}

var sampleProgram = []uint8{
	0xA9, 0x01, // LDA #$01 -> load the value 1 into the accumulator
	// 0x85, 0x10, // STA $10
}

func loadSampleProgram(c *CPU) {
	for i, b := range sampleProgram {
		c.Memory[0xFFFC+uint16(i)] = b // load program into memory starting at address 0xFFFC
	}
}

func executeInstruction(c *CPU, instr Instruction) {

	switch instr.Opcode {
	case 0xA9: // LDA Immediate\
		log.Println("Running LDA command")
		value := c.Memory[c.PC+1]                   // operand is the byte immediately following the opcode
		c.A = value                                 //load value into accumulator
		c.PC += instructionSet[instr.Opcode].Length // move program counter to next instruction
	}
}
