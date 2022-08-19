package asm

// RISCVInstruction indicates the structed instruction of RISCV.
type RISCVInstruction struct {
	EncodedBinary uint32
	MetaData      RISCVInstructionMetaData
}

// RISCVInstructionMetaData indicates the structed instruction of RISCV metadata.
type RISCVInstructionMetaData struct {
	Opcode            int
	InstructionClass  string
	InstructionFormat string
	InstructionName   string
	Funct3ToMetaData  map[int]RISCVInstructionMetaData
	Funct7ToMetaData  map[int]RISCVInstructionMetaData
}

const (
	// ISAClassRV32I indicates RISCV RV32I.
	ISAClassRV32I = "RV32I"

	// ISAFormatRegister indicates RISCV register format.
	ISAFormatRegister = "Register(R)"
	// ISAFormatLoad indicates RISCV load and short imm format.
	ISAFormatLoad = "Load(I)"
	// ISAFormatStore indicates RISCV store format.
	ISAFormatStore = "Store(S)"
	// ISAFormatConditionalJump indicates RISCV conditional jump format.
	ISAFormatConditionalJump = "ConditionalJump(B)"
	// ISAFormatLongImmediate indicates RISCV long imm format.
	ISAFormatLongImmediate = "LongImmediate(U)"
	// ISAFormatUnconditionalJump indicates RISCV unconditional jump format.
	ISAFormatUnconditionalJump = "UnconditionalJump(J)"

	unknownInstructionName  = "unknown instruction name"
	invalidInstructionValue = -1

	opcodeOffset = 0
	opcodeMask   = uint32(0x7f) << opcodeOffset

	funct3Offset = 12
	funct3Mask   = uint32(0x7) << funct3Offset

	funct7Offset = 25
	funct7Mask   = uint32(0x7f) << funct7Offset
)

func getBitMask(binary uint32, mask, offset uint32) uint32 {
	return (binary & mask) >> offset
}

func (instr *RISCVInstruction) getOpcode() uint32 {
	return getBitMask(instr.EncodedBinary, opcodeMask, opcodeOffset)
}

func (instr *RISCVInstruction) getFunct3() uint32 {
	return getBitMask(instr.EncodedBinary, funct3Mask, funct3Offset)
}

func (instr *RISCVInstruction) getFunct7() uint32 {
	return getBitMask(instr.EncodedBinary, funct7Mask, funct7Offset)
}
