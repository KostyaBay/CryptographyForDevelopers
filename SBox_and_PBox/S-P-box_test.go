package main

import (
	"testing"
)

func TestPBox(t *testing.T) {
	result := PBox(uint8(200))
	if PBoxInv(result) != uint8(200) {
		t.Errorf("Expected 200, but got %d", result)
	}
}

func TestSBox(t *testing.T) {
	result := SBox(uint8(200))
	if SBoxInv(result) != uint8(200) {
		t.Errorf("Expected 200, but got %d", result)
	}
}

func TestPBox_2(t *testing.T) {
	result := PBox(uint8(129))
	if PBoxInv(result) != uint8(129) {
		t.Errorf("Expected 129, but got %d", result)
	}
}

func TestSBox_2(t *testing.T) {
	result := SBox(uint8(129))
	if SBoxInv(result) != uint8(129) {
		t.Errorf("Expected 129, but got %d", result)
	}
}
