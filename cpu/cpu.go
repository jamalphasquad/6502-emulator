package cpu

type CPU struct {
	A  uint8  // accumulator
	X  uint8  // index register X
	Y  uint8  // index register Y
	PC uint16 //program counter

	//actually 8 bit not 16, hardware hardcodes 0x01 infront of it so we are using 16bit to make it easier to work with
	SP uint16 //-> stack is located at 0x0100-0x01FF (2^8 = 256 bytes), so SP is an offset from 0x0100

	P uint8 //processor status

	// 64kb of addressable memory (16-bit address bus) 2^16 = 65536
	Memory [65536]uint8
}

// P (Processor Status) register flags:
var (
	PFLAG_CARRY     uint8 = 1 << 0
	PFLAG_ZERO      uint8 = 1 << 1
	PFLAG_INTERRUPT uint8 = 1 << 2
	PFLAG_DECIMAL   uint8 = 1 << 3
	PFLAG_BREAK     uint8 = 1 << 4
	//bit 5 is unused and always set to 1
	PFLAG_UNUSED   uint8 = 1 << 5
	PFLAG_OVERFLOW uint8 = 1 << 6
	PFLAG_NEGATIVE uint8 = 1 << 7
)

// 0000-00FF: Zero Page -> first 256 bytes of memory, often used for fast access to variables and pointers
// 0100-01FF: Stack -> stack pointer (SP) is an offset from 0x0100
// 0200-07FF: RAM	-> general purpose
// 0800-FFFF: ROM (cartridge in case of NES) -> read-only memory, typically contains game code and data

//Vectors: ($FFFA–$FFFF: 16-bit addresses stored in little-endian format):
// FFFA-FFFB: NMI vector -> address to jump to on non-maskable interrupt
// FFFC-FFFD: Reset vector -> address to jump to on reset (computer also reads from here on boot)
// FFFE-FFFF: IRQ/BRK vector -> address to jump to on interrupt or break

func (c *CPU) Reset() {
	c.PC = uint16(0xFFFC) // set program counter to reset vector address

	//stack pointer is 8bit, it cant store 0x0100, so we set it to 0 and treat it as an offset from 0x0100
	c.SP = 0x0100 // set stack pointer to top of stack

	//reset all flags in processor status register
	c.P = 0
	//using bitfield operations like this:
	//c.P |= PFLAG_DECIMAL // set decimal mode flag of P (processor status) on reset

	c.A, c.X, c.Y = 0, 0, 0 // clear registers

	//reset memory to 0:
	for i := range c.Memory {
		c.Memory[i] = 0
	}

}

// TODO: stub for now (executing n number of instructions)
func (c *CPU) ExectuteCode(numInstructions uint32) {

	for i := uint32(0); i < numInstructions; i++ {
		instruction := c.fetchInstruction()
	}
}

func (c *CPU) fetchInstruction() uint8 {
	instr := c.Memory[c.PC] // fetch instruction at current program counter
	println("fetching instruction: ", instr)
	c.PC++ // increment program counter to point to next instruction
	return instr
}
