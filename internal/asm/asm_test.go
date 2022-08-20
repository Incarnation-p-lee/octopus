package asm

import (
	"github.com/Incarnation-p-lee/cachalot/pkg/assert"
	"testing"
)

func TestGetOpcode(t *testing.T) {
	testOpcodes := []struct {
		binary         uint32
		expectedOpcode uint32
	}{
		{
			uint32(0x0),
			uint32(0x0),
		},
		{
			uint32(0x1),
			uint32(0x1),
		},
		{
			uint32(0b1111111),
			uint32(0b1111111),
		},
		{
			uint32(0b11111111),
			uint32(0b1111111),
		},
	}

	for _, v := range testOpcodes {
		instr := &RISCVInstruction{
			EncodedBinary: v.binary,
		}
		opcode := instr.getOpcode()

		assert.IsEqual(t, v.expectedOpcode, opcode, "instruction opcode should be same")
	}
}

func TestGetFunct3(t *testing.T) {
	testFunct3s := []struct {
		binary         uint32
		expectedFunct3 uint32
	}{
		{
			uint32(0x0),
			uint32(0x0),
		},
		{
			uint32(0x1fff),
			uint32(0x1),
		},
		{
			uint32(0x7fff),
			uint32(0x7),
		},
		{
			uint32(0xffff),
			uint32(0x7),
		},
	}

	for _, v := range testFunct3s {
		instr := &RISCVInstruction{
			EncodedBinary: v.binary,
		}
		funct3 := instr.getFunct3()

		assert.IsEqual(t, v.expectedFunct3, funct3, "instruction funct3 should be same")
	}
}

func TestGetFunct7(t *testing.T) {
	testFunct7s := []struct {
		binary         uint32
		expectedFunct7 uint32
	}{
		{
			uint32(0x0),
			uint32(0x0),
		},
		{
			uint32(0x02000000),
			uint32(0x1),
		},
		{
			uint32(0xfe000000),
			uint32(0x7f),
		},
		{
			uint32(0xffffffff),
			uint32(0x7f),
		},
	}

	for _, v := range testFunct7s {
		instr := &RISCVInstruction{
			EncodedBinary: v.binary,
		}
		funct7 := instr.getFunct7()

		assert.IsEqual(t, v.expectedFunct7, funct7, "instruction funct7 should be same")
	}
}

func TestGetRd(t *testing.T) {
	testRds := []struct {
		instruction RISCVInstruction
		expectedRd  int
	}{
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0x80),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeRegister},
			},
			expectedRd: 0x1,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0x80),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeLoad},
			},
			expectedRd: 0x1,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0x80),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeLongImmediate},
			},
			expectedRd: 0x1,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0xffff),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeUnconditionalJump},
			},
			expectedRd: 0x1f,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0xffff),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeStore},
			},
			expectedRd: -1,
		},
	}

	for _, v := range testRds {
		rd := v.instruction.getRd()

		assert.IsEqual(t, v.expectedRd, rd, "instruction rd should be same")
	}
}

func TestGetRs1(t *testing.T) {
	testRs1s := []struct {
		instruction RISCVInstruction
		expectedRs1 int
	}{
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0xffff),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeRegister},
			},
			expectedRs1: 0x1,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0xfffff),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeLoad},
			},
			expectedRs1: 0x1f,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0xfffff),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeStore},
			},
			expectedRs1: 0x1f,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0x1fffff),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeConditionalJump},
			},
			expectedRs1: 0x1f,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0xffffff),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeLongImmediate},
			},
			expectedRs1: -1,
		},
	}

	for _, v := range testRs1s {
		rs1 := v.instruction.getRs1()

		assert.IsEqual(t, v.expectedRs1, rs1, "instruction rs1 should be same")
	}
}

func TestGetRs2(t *testing.T) {
	testRs2s := []struct {
		instruction RISCVInstruction
		expectedRs2 int
	}{
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0x1fffff),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeRegister},
			},
			expectedRs2: 0x1,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0x1ffffff),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeStore},
			},
			expectedRs2: 0x1f,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0x3ffffff),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeConditionalJump},
			},
			expectedRs2: 0x1f,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0xffffff),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeLongImmediate},
			},
			expectedRs2: -1,
		},
	}

	for _, v := range testRs2s {
		rs2 := v.instruction.getRs2()

		assert.IsEqual(t, v.expectedRs2, rs2, "instruction rs2 should be same")
	}
}

func TestGetImm(t *testing.T) {
	testImms := []struct {
		instruction RISCVInstruction
		expectedImm int
	}{
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0xffff0000),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeLoad},
			},
			expectedImm: 0xfff,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0xfffff0ff),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeStore},
			},
			expectedImm: 0b111111100001,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0x8ffffff0),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeConditionalJump},
			},
			expectedImm: 0b1100011111110,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0xfaf0ffff),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeLongImmediate},
			},
			expectedImm: 0xfaf0f000,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0xf9f32fff),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeUnconditionalJump},
			},
			expectedImm: 0b100110010111110011110,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0xffffff),
				Metadata:      RISCVInstructionMetadata{InstructionType: ISATypeRegister},
			},
			expectedImm: -1,
		},
	}

	for _, v := range testImms {
		imm := v.instruction.getImm()

		assert.IsEqual(t, v.expectedImm, imm, "instruction imm should be same")
	}
}
