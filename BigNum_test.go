package main

import (
	"testing"
)

func TestInverse(t *testing.T) {
	var num BigNum
	hex := "51BF608414AD5726A3C1BEC098F77B1B54FFB2787F8D528A74C1D7FDE6470EA4"
	num.setHex(hex)
	result := Inverse(num)
	if result.getHex() != "ae409f7beb52a8d95c3e413f670884e4ab004d878072ad758b3e280219b8f15b" {
		t.Errorf("Expected another number, but got %d", result)
	}
}

func TestXOR(t *testing.T) {
	var num1, num2 BigNum
	hex1 := "51BF608414AD5726A3C1BEC098F77B1B54FFB2787F8D528A74C1D7FDE6470EA4"
	hex2 := "403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c"
	num1.setHex(hex1)
	num2.setHex(hex2)
	result := XOR(num1, num2)
	if result.getHex() != "1182d8299c0ec40ca8bf3f49362e95e4ecedaf82bfd167988972412095b13db8" {
		t.Errorf("Expected another number, but got %d", result)
	}
}

func TestOR(t *testing.T) {
	var num1, num2 BigNum
	hex1 := "51BF608414AD5726A3C1BEC098F77B1B54FFB2787F8D528A74C1D7FDE6470EA4"
	hex2 := "403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c"
	num1.setHex(hex1)
	num2.setHex(hex2)
	result := OR(num1, num2)
	if result.getHex() != "51bff8ad9cafd72eabffbfc9befffffffcffbffaffdd779afdf3d7fdf7f73fbc" {
		t.Errorf("Expected another number, but got %d", result)
	}
}

func TestAND(t *testing.T) {
	var num1, num2 BigNum
	hex1 := "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4"
	hex2 := "403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c"
	num1.setHex(hex1)
	num2.setHex(hex2)
	result := AND(num1, num2)
	if result.getHex() != "403d208400a113220340808088d16a1b10121078400c1002748196dd62460204" {
		t.Errorf("Expected another number, but got %d", result)
	}
}

func TestShiftR(t *testing.T) {
	var num BigNum
	n := 5
	hex := "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4"
	num.setHex(hex)
	result := ShiftR(num, n)
	if result.getHex() != "28dfb0420a56ab9351e0df604c7bbd8daa7fd93c3fc6a9453a60ebfef323875\n" {
		t.Errorf("Expected another number, but got %d", result)
	}
}

func TestShiftL(t *testing.T) {
	var num BigNum
	n := 5
	hex := "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4"
	num.setHex(hex)
	result := ShiftR(num, n)
	if result.getHex() != "a37ec108295aae4d47837d8131eef636a9ff64f0ff1aa514e983affbcc8e1d480" {
		t.Errorf("Expected another number, but got %d", result)
	}
}
