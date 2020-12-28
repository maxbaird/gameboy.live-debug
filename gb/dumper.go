package gb

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func buildRegisterState(core *Core) string {
	A := core.CPU.Registers.A
	B := core.CPU.Registers.B
	C := core.CPU.Registers.C
	D := core.CPU.Registers.D
	E := core.CPU.Registers.E
	F := core.CPU.Registers.F
	HL := core.CPU.Registers.HL
	PC := core.CPU.Registers.PC
	SP := core.CPU.Registers.SP

	zero := core.CPU.Flags.Zero
	sub := core.CPU.Flags.Sub
	halfCarry := core.CPU.Flags.HalfCarry
	carry := core.CPU.Flags.Carry

	var str strings.Builder

	str.WriteString(fmt.Sprintf("A: %v\nB: %v\nC: %v\nD: %v\nE: %v\nF: %v\nHL: %v\n", A, B, C, D, E, F, HL))
	str.WriteString("===========\n")
	str.WriteString(fmt.Sprintf("Z: %v\nN: %v\nH: %v\nC: %v\n", zero, sub, halfCarry, carry))
	str.WriteString("===========\n")
	str.WriteString(fmt.Sprintf("sp: %v\npc: %v", SP, PC))
	str.WriteString("\n")

	return str.String()
}

func buildMemoryState(core *Core) string {
	var str strings.Builder

	for idx, val := range core.Memory.MainMemory {
		str.WriteString(fmt.Sprintf("0x%02X: %v\n", idx, val))
	}

	return str.String()
}

func writeToFile(filename string, data string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, data)
	if err != nil {
		return err
	}
	return file.Sync()
}

func dumpState(core *Core, iteration int, opcode int, extendedOpcode int, cycles int) {
	str := ""

	str += fmt.Sprintf("Iteration: %d\n", iteration)
	str += fmt.Sprintf("Opcode: 0x%02X\n", opcode)
	str += fmt.Sprintf("Extended Opcode: 0x%02X\n", extendedOpcode)
	str += fmt.Sprintf("Cycles: %d\n", cycles)
	str += "===========\n"
	str += buildRegisterState(core)
	str += "===========\n"
	str += buildMemoryState(core)
	fname := fmt.Sprintf("%s%d%s", "/tmp/", iteration, ".gblive")

	fmt.Println("Started writing ", fname)

	err := writeToFile(fname, str)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Finished writing", fname)
}
