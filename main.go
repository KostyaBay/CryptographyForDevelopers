package main

import (
	"fmt"
	"strconv"
)

type BigNum struct {
	arr []uint64
}

func (BG *BigNum) setHex(hex string) {
	var newHex uint64
	for i := 0; i < len(hex); i += 16 {

		tempRem := len(hex) % 16
		if tempRem != 0 {
			temp := 16 - tempRem
			// Доповнення рядка "0" до необхідної довжини
			for temp > 0 {
				hex = "0" + hex
				temp--
			}
		} else {
			newHex, _ = strconv.ParseUint(hex[i:(i+16)], 16, 64)
			BG.arr = append(BG.arr, newHex)
		}
	}
}

func (BG BigNum) getHex() string {
	var resString string
	for i := 0; i < len(BG.arr); i++ {
		tempString := strconv.FormatUint(BG.arr[i], 16)
		if len(tempString) != 16 {
			tempString = "0" + tempString
		}
		resString += tempString
	}
	return resString
}

func Inverse(num BigNum) BigNum {
	var resInv BigNum
	for i := 0; i < len(num.arr); i++ {
		resInv.arr = append(resInv.arr, ^num.arr[i])
	}
	return resInv
}

func XOR(num1, num2 BigNum) BigNum {
	var resXOR BigNum
	var count_i int
	if len(num1.arr) >= len(num2.arr) {
		count_i = len(num1.arr)
	} else {
		count_i = len(num2.arr)
	}

	for i := 0; i < count_i; i++ {
		resXOR.arr = append(resXOR.arr, num1.arr[i]^num2.arr[i])
	}
	return resXOR
}

func OR(num1, num2 BigNum) BigNum {
	var resOR BigNum
	var count_i int
	if len(num1.arr) >= len(num2.arr) {
		count_i = len(num1.arr)
	} else {
		count_i = len(num2.arr)
	}

	for i := 0; i < count_i; i++ {
		resOR.arr = append(resOR.arr, num1.arr[i]|num2.arr[i])
	}
	return resOR
}

func AND(num1, num2 BigNum) BigNum {
	var resAND BigNum
	var count_i int
	if len(num1.arr) >= len(num2.arr) {
		count_i = len(num1.arr)
	} else {
		count_i = len(num2.arr)
	}

	for i := 0; i < count_i; i++ {
		resAND.arr = append(resAND.arr, num1.arr[i]&num2.arr[i])
	}
	return resAND
}

func ShiftR(num BigNum, n int) BigNum {
	var resShiftR BigNum
	trash_el := n / 64
	tempRem := n % 64
	if trash_el >= len(num.arr) {
		resShiftR.setHex("0")
	}
	num.arr = num.arr[trash_el:]
	resShiftR.arr = append(num.arr, num.arr[0]>>tempRem)
	for i := 1; i < len(num.arr); i++ {
		resShiftR.arr[i-1] += ((^uint64(0) >> (64 - tempRem)) & num.arr[i]) << (64 - tempRem)
		resShiftR.arr = append(num.arr, num.arr[i]>>tempRem)
	}
	return resShiftR
}

func ShiftL(num BigNum, n int) BigNum {
	var resShiftL BigNum
	var SaveBits uint64
	trash_el := n / 64
	tempRem := n % 64

	for i := 0; i < trash_el; i++ {
		resShiftL.setHex("0")
	}

	for i := 0; i < len(num.arr); i++ {
		resShiftL.arr = append(resShiftL.arr, (num.arr[i]<<tempRem)+SaveBits)
		SaveBits = (^uint64((0xffffffffffffffff>>tempRem)&num.arr[i]) >> (64 - tempRem))
	}

	if num.arr[len(num.arr)-1] & ^uint64(0xffffffffffffffff>>tempRem) != 0 {
		resShiftL.arr = append(resShiftL.arr, num.arr[len(num.arr)-1]>>(64-tempRem))
	}

	return resShiftL
}

func main() {
	var num1, num2, resXOR, resInv, resOR, resAND, resShiftR, resShiftL BigNum
	num1.setHex("51BF608414AD5726A3C1BEC098F77B1B54FFB2787F8D528A74C1D7FDE6470EA4") // "123EA0")
	num2.setHex("403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c") //"123EF0")
	fmt.Println("num1: " + num1.getHex())
	fmt.Println("num2: " + num2.getHex())
	resXOR = XOR(num1, num2)
	fmt.Println(" XOR: " + resXOR.getHex())
	resInv = Inverse(num1)
	fmt.Println("Inv1: " + resInv.getHex())
	resOR = OR(num1, num2)
	fmt.Println("  OR: " + resOR.getHex())
	resAND = AND(num1, num2)
	fmt.Println(" AND: " + resAND.getHex())
	resShiftR = ShiftR(num1, 5)
	fmt.Println(" 1>>: " + resShiftR.getHex())
	resShiftL = ShiftL(num1, 5)
	fmt.Println(" 1<<: " + resShiftL.getHex())
}
