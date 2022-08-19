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
