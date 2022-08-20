package asm

import (
	"fmt"
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
				Metadata:      RISCVInstructionMetadata{InstructionFormat: ISAFormatRegister},
			},
			expectedRd: 0x1,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0x80),
				Metadata:      RISCVInstructionMetadata{InstructionFormat: ISAFormatLoad},
			},
			expectedRd: 0x1,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0x80),
				Metadata:      RISCVInstructionMetadata{InstructionFormat: ISAFormatLongImmediate},
			},
			expectedRd: 0x1,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0xffff),
				Metadata:      RISCVInstructionMetadata{InstructionFormat: ISAFormatUnconditionalJump},
			},
			expectedRd: 0x1f,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0xffff),
				Metadata:      RISCVInstructionMetadata{InstructionFormat: ISAFormatStore},
			},
			expectedRd: -1,
		},
	}

	for _, v := range testRds {
		rd := v.instruction.getRd()
		fmt.Printf("rd %v from instr %+v\n", rd, v.instruction)

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
				Metadata:      RISCVInstructionMetadata{InstructionFormat: ISAFormatRegister},
			},
			expectedRs1: 0x1,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0xfffff),
				Metadata:      RISCVInstructionMetadata{InstructionFormat: ISAFormatLoad},
			},
			expectedRs1: 0x1f,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0xfffff),
				Metadata:      RISCVInstructionMetadata{InstructionFormat: ISAFormatStore},
			},
			expectedRs1: 0x1f,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0x1fffff),
				Metadata:      RISCVInstructionMetadata{InstructionFormat: ISAFormatConditionalJump},
			},
			expectedRs1: 0x1f,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0xffffff),
				Metadata:      RISCVInstructionMetadata{InstructionFormat: ISAFormatLongImmediate},
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
				Metadata:      RISCVInstructionMetadata{InstructionFormat: ISAFormatRegister},
			},
			expectedRs2: 0x1,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0x1ffffff),
				Metadata:      RISCVInstructionMetadata{InstructionFormat: ISAFormatStore},
			},
			expectedRs2: 0x1f,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0x3ffffff),
				Metadata:      RISCVInstructionMetadata{InstructionFormat: ISAFormatConditionalJump},
			},
			expectedRs2: 0x1f,
		},
		{
			instruction: RISCVInstruction{
				EncodedBinary: uint32(0xffffff),
				Metadata:      RISCVInstructionMetadata{InstructionFormat: ISAFormatLongImmediate},
			},
			expectedRs2: -1,
		},
	}

	for _, v := range testRs2s {
		rs2 := v.instruction.getRs2()

		assert.IsEqual(t, v.expectedRs2, rs2, "instruction rs2 should be same")
	}
}
