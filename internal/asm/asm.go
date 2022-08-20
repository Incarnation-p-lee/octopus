package asm

// RISCVInstructionMetadata indicates the structed instruction of RISCV metadata.
type RISCVInstructionMetadata struct {
	Opcode           int
	InstructionClass string
	InstructionType  string
	InstructionName  string
}

// RISCVInstruction indicates the structed instruction of RISCV.
type RISCVInstruction struct {
	EncodedBinary uint32
	Metadata      RISCVInstructionMetadata
}

// RISCVInstructionEntity indicates the nested instruction entity of RISCV metadata.
type RISCVInstructionEntity struct {
	Metadata         RISCVInstructionMetadata
	Funct3ToMetadata map[int]RISCVInstructionMetadata
	Funct7ToMetadata map[int]RISCVInstructionMetadata
}

const (
	// ISAClassRV32I indicates RISCV RV32I.
	ISAClassRV32I = "RV32I"

	// ISATypeRegister indicates RISCV register format.
	ISATypeRegister = "Register(R)"
	// ISATypeLoad indicates RISCV load and short imm format.
	ISATypeLoad = "Load(I)"
	// ISATypeStore indicates RISCV store format.
	ISATypeStore = "Store(S)"
	// ISATypeConditionalJump indicates RISCV conditional jump format.
	ISATypeConditionalJump = "ConditionalJump(B)"
	// ISATypeLongImmediate indicates RISCV long imm format.
	ISATypeLongImmediate = "LongImmediate(U)"
	// ISATypeUnconditionalJump indicates RISCV unconditional jump format.
	ISATypeUnconditionalJump = "UnconditionalJump(J)"

	unknownInstructionName  = "unknown instruction name"
	invalidInstructionValue = -1

	opcodeOffset = 0
	opcodeMask   = uint32(0x7f) << opcodeOffset

	funct3Offset = 12
	funct3Mask   = uint32(0x7) << funct3Offset

	funct7Offset = 25
	funct7Mask   = uint32(0x7f) << funct7Offset

	rdOffset = 7
	rdMask   = uint32(0x1f) << rdOffset

	rs1Offset = 15
	rs1Mask   = uint32(0x1f) << rs1Offset

	rs2Offset = 20
	rs2Mask   = uint32(0x1f) << rs2Offset
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

func (instr *RISCVInstruction) getRd() int {
	switch instr.Metadata.InstructionType {
	case ISATypeRegister, ISATypeLoad, ISATypeLongImmediate, ISATypeUnconditionalJump:
		return int(getBitMask(instr.EncodedBinary, rdMask, rdOffset))
	}

	return invalidInstructionValue
}

func (instr *RISCVInstruction) getRs1() int {
	switch instr.Metadata.InstructionType {
	case ISATypeRegister, ISATypeLoad, ISATypeStore, ISATypeConditionalJump:
		return int(getBitMask(instr.EncodedBinary, rs1Mask, rs1Offset))
	}

	return invalidInstructionValue
}

func (instr *RISCVInstruction) getRs2() int {
	switch instr.Metadata.InstructionType {
	case ISATypeRegister, ISATypeStore, ISATypeConditionalJump:
		return int(getBitMask(instr.EncodedBinary, rs2Mask, rs2Offset))
	}

	return invalidInstructionValue
}
