package main

import "cpu6502/cpu"

func main() {
	println("...6502 CPU Emulator starting...")
	cpu := cpu.CPU{}
	cpu.Reset()
}
