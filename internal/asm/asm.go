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

	typeIImm11to0Offset = 20
	typeIImm11to0Mask = uint32(0xfff) << typeIImm11to0Offset

	typeSImm4To0Offset = 7
	typeSImm4To0Mask = uint32(0x1f) << typeSImm4To0Offset
	typeSImm11To5Offset = 20
	typeSImm11To5Mask = uint32(0x7f) << typeSImm11To5Offset

	typeBImm4To1Offset = 8
	typeBImm4To1Mask = uint32(0xf) << typeBImm4To1Offset
	typeBImm10To5Offset = 25
	typeBImm10To5Mask = uint32(0x3f) << typeBImm10To5Offset
	typeBImm11Offset = 7
	typeBImm11Mask = uint32(0x1) << typeBImm11Offset
	typeBImm12Offset = 31
	typeBImm12Mask = uint32(0x1) << typeBImm12Offset

	typeUImm32to12Offset = 12
	typeUImm32to12Mask = uint32(0xfffff) << typeUImm32to12Offset

	typeJImm10To1Offset = 21
	typeJImm10To1Mask = uint32(0x3ff) << typeJImm10To1Offset
	typeJImm11Offset = 20
	typeJImm11Mask = uint32(0x1) << typeJImm11Offset
	typeJImm19To12Offset = 12
	typeJImm19To12Mask = uint32(0xff) << typeJImm19To12Offset
	typeJImm20Offset = 31
	typeJImm20Mask = uint32(0x1) << typeJImm20Offset
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

func (instr *RISCVInstruction) getImm() int {
	switch instr.Metadata.InstructionType {
	case ISATypeLoad:
		return int(instr.getITypeImm())
	case ISATypeStore:
		return int(instr.getSTypeImm())
	case ISATypeConditionalJump:
		return int(instr.getBTypeImm())
	case ISATypeLongImmediate:
		return int(instr.getUTypeImm())
	case ISATypeUnconditionalJump:
		return int(instr.getJTypeImm())
	}

	return invalidInstructionValue
}

func (instr *RISCVInstruction) getITypeImm() uint32 {
	return getBitMask(instr.EncodedBinary, typeIImm11to0Mask, typeIImm11to0Offset)
}

func (instr *RISCVInstruction) getSTypeImm() uint32 {
	imm4To0 := getBitMask(instr.EncodedBinary, typeSImm4To0Mask, typeSImm4To0Offset)
	imm11To5 := getBitMask(instr.EncodedBinary, typeSImm11To5Mask, typeSImm11To5Offset)

	return (imm11To5 << 5) + imm4To0
}

func (instr *RISCVInstruction) getBTypeImm() uint32 {
	imm4To1 := getBitMask(instr.EncodedBinary, typeBImm4To1Mask, typeBImm4To1Offset)
	imm10To5 := getBitMask(instr.EncodedBinary, typeBImm10To5Mask, typeBImm10To5Offset)
	imm11 := getBitMask(instr.EncodedBinary, typeBImm11Mask, typeBImm11Offset)
	imm12 := getBitMask(instr.EncodedBinary, typeBImm12Mask, typeBImm12Offset)

	return (imm12 << 12) + (imm11 << 11) + (imm10To5 << 5) + (imm4To1 << 1)
}

func (instr *RISCVInstruction) getUTypeImm() uint32 {
	imm32to12 := getBitMask(instr.EncodedBinary, typeUImm32to12Mask, typeUImm32to12Offset)

	return imm32to12 << 12
}

func (instr *RISCVInstruction) getJTypeImm() uint32 {
	imm10To1 := getBitMask(instr.EncodedBinary, typeJImm10To1Mask, typeJImm10To1Offset)
	imm11 := getBitMask(instr.EncodedBinary, typeJImm11Mask, typeJImm11Offset)
	imm19To12 := getBitMask(instr.EncodedBinary, typeJImm19To12Mask, typeJImm19To12Offset)
	imm20 := getBitMask(instr.EncodedBinary, typeJImm20Mask, typeJImm20Offset)

	return (imm20 << 20) + (imm19To12 << 12) + (imm11 << 11) + (imm10To1 << 1)
}
