package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

// function for input value in SHA-1
func inputValue(inStr []byte) []uint32 {
	var res []uint32
	l := len(inStr) * 8
	rem := l % 512

	if rem < 440 {
		inStr = append(inStr, uint8(128)) //128 = 1000 0000
		for i := 0; i < (448-(rem+8))/8; i++ {
			inStr = append(inStr, 0)
		}
		inStr = binary.BigEndian.AppendUint64(inStr, uint64(l))
	}
	if rem == 440 {
		inStr = append(inStr, uint8(128))
		inStr = binary.BigEndian.AppendUint64(inStr, uint64(l))
	} else if rem > 440 {
		inStr = append(inStr, uint8(128))
		for i := 0; i < (512-(rem+8)+448)/8; i++ {
			inStr = append(inStr, 0)
		}
		inStr = binary.BigEndian.AppendUint64(inStr, uint64(l))
	}

	for i := 0; i < len(inStr)/4; i++ {
		res = append(res, uint32(inStr[i*4])<<24+uint32(inStr[i*4+1])<<16+uint32(inStr[i*4+2])<<8+uint32(inStr[i*4+3]))
	}

	return res
}

// cyclic left shift function
func shiftLeft(block uint32, n int) uint32 {
	for i := 0; i < n; i++ {
		temp1 := block >> 31
		temp2 := block << 1
		block = temp2 + temp1
	}
	return block
}

// function of determining four non-linear operations
func operation_F(m uint32, l uint32, k uint32, i int) uint32 {
	if 0 <= i && i <= 19 {
		return (m & l) | (^m & k)
	}
	if 20 <= i && i <= 39 {
		return m ^ l ^ k
	}
	if 40 <= i && i <= 59 {
		return (m & l) | (m & k) | (l & k)
	}
	if 60 <= i && i <= 79 {
		return m ^ l ^ k
	} else {
		return 0
	}
}

// function of determining four constants
func const_K(i int) uint32 {
	if 0 <= i && i <= 19 {
		return 0x5A827999
	}
	if 20 <= i && i <= 39 {
		return 0x6ED9EBA1
	}
	if 40 <= i && i <= 59 {
		return 0x8F1BBCDC
	}
	if 60 <= i && i <= 79 {
		return 0xCA62C1D6
	} else {
		return 0
	}
}

// own implementation of the SHA-1 hashing algorithm
func SHA_1(inMess []byte) []byte {
	var W [80]uint32

	M := inputValue(inMess)

	var A uint32 = 0x67452301
	var B uint32 = 0xEFCDAB89
	var C uint32 = 0x98BADCFE
	var D uint32 = 0x10325476
	var E uint32 = 0xC3D2E1F0

	for i := 0; i < len(M)/16; i++ {

		for j := 0; j < 80; j++ {
			if j >= 0 && j <= 15 {
				W[j] = M[j]
			}
			if j >= 16 && j <= 79 {
				W[j] = shiftLeft((W[j-3] ^ W[j-8] ^ W[j-14] ^ W[j-16]), 1)
			}
		}

		var a uint32 = A
		var b uint32 = B
		var c uint32 = C
		var d uint32 = D
		var e uint32 = E

		for i := 0; i < 80; i++ {
			temp := shiftLeft(a, 5) + operation_F(b, c, d, i) + e + W[i] + const_K(i)
			e = d
			d = c
			c = shiftLeft(b, 30)
			b = a
			a = temp
		}
		A += a
		B += b
		C += c
		D += d
		E += e
	}

	var resUint = [5]uint32{A, B, C, D, E}

	var resByte = UintArrToByteArr(resUint)

	return resByte
}

// function of converting the hash of the []uint32 into []byte
func UintArrToByteArr(uintArr [5]uint32) []byte {
	byteArr := make([]byte, len(uintArr)*4)

	for i, v := range uintArr {
		index := i * 4
		byteArr[index] = byte(v >> 24)
		byteArr[index+1] = byte(v >> 16)
		byteArr[index+2] = byte(v >> 8)
		byteArr[index+3] = byte(v)
	}

	return byteArr
}

func main() {
	res := SHA_1([]byte("1ABC104FE402BD"))

	fmt.Println("SHA1 to ", "1ABC104FE402BD", " is ", res)

	HashString := hex.EncodeToString(res[:])
	fmt.Println("SHA1 to ", "1ABC104FE402BD", " is ", HashString)
}
