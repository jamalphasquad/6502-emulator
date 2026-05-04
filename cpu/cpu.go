package cpu

const (
	FlagC uint8 = 1 << 0 // Carry
	FlagZ uint8 = 1 << 1 // Zero
	FlagI uint8 = 1 << 2 // IRQ Disable
	FlagD uint8 = 1 << 3 // Decimal  (present on chip, ignored in most emulators)
	FlagB uint8 = 1 << 4 // Break    (set when BRK pushes P onto stack)
	FlagU uint8 = 1 << 5 // Unused   (always reads 1 on the real chip)
	FlagV uint8 = 1 << 6 // Overflow
	FlagN uint8 = 1 << 7 // Negative
)

type CPU struct {
	// Registers
	A  uint8  // Accumulator       — arithmetic/logic results land here
	X  uint8  // Index register X  — loop counters, address offsets
	Y  uint8  // Index register Y  — loop counters, address offsets
	PC uint16 // Program Counter   — address of next instruction to fetch
	SP uint8  // Stack Pointer     — offset into page $01 (stack memory: $0100–$01FF)
	P  uint8  // Processor Status  — 8 individual flag bits (see constants above)

	// 64 KB flat address space
	Memory [65536]uint8

	// Total cycles consumed so far (useful for timing / debugging)
	Cycles uint64
}

// New returns a CPU with sensible power-on defaults.
// Call Reset() afterwards to load the reset vector into PC.
func New() *CPU {
	c := &CPU{}
	// The unused flag (bit 5) is always 1 on the real chip.
	c.P = FlagU | FlagI
	// Stack pointer starts at $FD after a real reset sequence
	// (the reset sequence performs 3 phantom stack decrements from $FF).
	c.SP = 0xFD
	return c
}
