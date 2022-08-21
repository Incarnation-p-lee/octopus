package asm

const (
	opcodeLui   = 0b0110111
	opcodeAuipc = 0b0010111
	opcodeJal   = 0b1101111
	opcodeJalr  = 0b1100111

	opcodeB           = 0b1100011
	opcodeBFunct3Beq  = 0b000
	opcodeBFunct3Bne  = 0b001
	opcodeBFunct3Blt  = 0b100
	opcodeBFunct3Bge  = 0b101
	opcodeBFunct3Bltu = 0b110
	opcodeBFunct3Bgeu = 0b111

	opcodeLoad          = 0b0000011
	opcodeLoadFunct3Lb  = 0b000
	opcodeLoadFunct3Lh  = 0b001
	opcodeLoadFunct3Lw  = 0b010
	opcodeLoadFunct3Lbu = 0b100
	opcodeLoadFunct3Lhu = 0b101
)

var opcodeToInstructionEntity = map[int]RISCVInstructionEntity{
	opcodeLui: RISCVInstructionEntity{
		Metadata: RISCVInstructionMetadata{
			Opcode:           opcodeLui,
			InstructionClass: ISAClassRV32I,
			InstructionType:  ISATypeLongImmediate,
			InstructionName:  "lui",
		},
	},
	opcodeAuipc: RISCVInstructionEntity{
		Metadata: RISCVInstructionMetadata{
			Opcode:           opcodeAuipc,
			InstructionClass: ISAClassRV32I,
			InstructionType:  ISATypeLongImmediate,
			InstructionName:  "auipc",
		},
	},
	opcodeJal: RISCVInstructionEntity{
		Metadata: RISCVInstructionMetadata{
			Opcode:           opcodeJal,
			InstructionClass: ISAClassRV32I,
			InstructionType:  ISATypeUnconditionalJump,
			InstructionName:  "jal",
		},
	},
	opcodeJalr: RISCVInstructionEntity{
		Metadata: RISCVInstructionMetadata{
			Opcode:           opcodeJalr,
			InstructionClass: ISAClassRV32I,
			InstructionType:  ISATypeLoad,
			InstructionName:  "jalr",
		},
	},
	opcodeB: RISCVInstructionEntity{
		Metadata: RISCVInstructionMetadata{
			Opcode:           opcodeB,
			InstructionClass: ISAClassRV32I,
			InstructionType:  ISATypeConditionalJump,
			InstructionName:  unknownInstructionName,
		},
		Funct3ToMetadata: map[int]RISCVInstructionMetadata{
			opcodeBFunct3Beq: RISCVInstructionMetadata{
				InstructionName: "beq",
			},
			opcodeBFunct3Bne: RISCVInstructionMetadata{
				InstructionName: "bne",
			},
			opcodeBFunct3Blt: RISCVInstructionMetadata{
				InstructionName: "blt",
			},
			opcodeBFunct3Bge: RISCVInstructionMetadata{
				InstructionName: "bge",
			},
			opcodeBFunct3Bltu: RISCVInstructionMetadata{
				InstructionName: "bltu",
			},
			opcodeBFunct3Bgeu: RISCVInstructionMetadata{
				InstructionName: "bgeu",
			},
		},
	},
	opcodeLoad: RISCVInstructionEntity{
		Metadata: RISCVInstructionMetadata{
			Opcode:           opcodeLoad,
			InstructionClass: ISAClassRV32I,
			InstructionType:  ISATypeLoad,
			InstructionName:  unknownInstructionName,
		},
		Funct3ToMetadata: map[int]RISCVInstructionMetadata{
			opcodeLoadFunct3Lb: RISCVInstructionMetadata{
				InstructionName: "lb",
			},
			opcodeLoadFunct3Lh: RISCVInstructionMetadata{
				InstructionName: "lh",
			},
			opcodeLoadFunct3Lw: RISCVInstructionMetadata{
				InstructionName: "lw",
			},
			opcodeLoadFunct3Lbu: RISCVInstructionMetadata{
				InstructionName: "lbu",
			},
			opcodeLoadFunct3Lhu: RISCVInstructionMetadata{
				InstructionName: "lhu",
			},
		},
	},
}
