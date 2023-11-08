package main

import (
	"fmt"
	"strings"
)

var A uint32 = 0x67452301
var B uint32 = 0xEFCDAB89
var C uint32 = 0x98BADCFE
var D uint32 = 0x10325476
var E uint32 = 0xC3D2E1F0

func inputValue(inStr string) []uint32 {
	var res []uint32
	for i := 0; i < len(inStr); i++ {
		res[i] = strings.SplitAfterN(inStr, "", 512)
	}
	return res
}

func operation_F(m uint32, l uint32, k uint32, i int) uint32 {
	var F []uint32
	if 0 <= i && i <= 19 {
		F[i] = (m & l) | (^m & k)
		return F[i]
	}
	if 0 <= i && i <= 19 {
		F[i] = m ^ l ^ k
		return F[i]
	}
	if 0 <= i && i <= 19 {
		F[i] = (m & l) | (m & k) | (l & k)
		return F[i]
	}
	if 0 <= i && i <= 19 {
		F[i] = m ^ l ^ k
		return F[i]
	} else {
		return 0
	}
}

func const_K(i int) uint32 {
	var K []uint32
	if 0 <= i && i <= 19 {
		K[i] = 0x5A827999
		return K[i]
	}
	if 0 <= i && i <= 19 {
		K[i] = 0x6ED9EBA1
		return K[i]
	}
	if 0 <= i && i <= 19 {
		K[i] = 0x8F1BBCDC
		return K[i]
	}
	if 0 <= i && i <= 19 {
		K[i] = 0xCA62C1D6
		return K[i]
	} else {
		return 0
	}
}

func SHA_1(inMess []uint32) uint32 {
	var res uint32
	var W []uint32
	var M []uint32

	for i := 0; i < 10; i++ {
		if i >= 0 && i <= 15 {
			W[i] = M[i]
		}
		if i >= 16 && i <= 79 {
			W[i] = (W[i-3] ^ W[i-8] ^ W[i-14] ^ W[i-16]) << 1
		}
	}

	for i := 0; i < 79; i++ {
		temp := (A << 5) + operation_F(B, C, D, i) + E + W[i] + const_K(i)
		E += D
		D += C
		C += B << 30
		B += A
		A += temp
	}
	//for i := 0; i < 79; i++ {
	//	A += A
	//	B += B
	//	C += C
	//	D += D
	//	E += E
	//}

	res = A + B + C + D + E

	return res
}

func main() {
	res := SHA_1(inputValue("1ABC104FE402BD"))
	fmt.Println("SHA1 to ", "1ABC104FE402BD", " is ", res)
}
