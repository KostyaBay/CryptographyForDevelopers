package main

import (
	"fmt"
	"math/rand"
)

func GenRandomBitChain() []uint32 {
	var BitChain []uint32
	for i := 0; i < 625; i++ {
		BitChain = append(BitChain, rand.Uint32())
	}
	return BitChain
}

func MonobitTest(BitChain []uint32) bool {
	counter := 0
	for j := 0; j < 625; j++ {
		for i := 31; i >= 0; i-- {
			tempValOne := BitChain[j] >> i
			tempValOne = tempValOne & uint32(1)

			if tempValOne == 0 {
				continue
			}
			if tempValOne == 1 {
				counter++
				continue
			}
		}
	}

	if counter > 9654 && counter < 10346 {
		return true
	}
	return false
}

func LongRunsTest(BitChain []uint32) bool {
	pointerBit := uint32(0)
	curCount := 0
	maxCount := 0
	for j := 0; j < 625; j++ {

		for i := 31; i >= 0; i-- {
			tempValBit := BitChain[j] >> i
			tempValBit = tempValBit & uint32(1)

			if tempValBit == pointerBit {
				curCount += 1
			} else {
				pointerBit = tempValBit
				if curCount > maxCount {
					maxCount = curCount
				}
			}
			//
			//if pointerBit == 0 {
			//
			//	if tempValBit == 1 {
			//		if counter1 >= counter1 {
			//			counter1++
			//		}
			//		pointerBit = 1
			//	}
			//	//break
			//
			//}
			//
			//if pointerBit == 1 {
			//
			//	if tempValBit == 0 {
			//		if counter0 >= counter0 {
			//			counter0++
			//		}
			//		pointerBit = 0
			//	}
			//	//break
			//}
		}
	}

	if maxCount <= 36 {
		return true
	}
	return false
}

func PokerTest(BitChain []uint32) bool {
	var counter [16]int
	var X float32

	for j := 0; j < 625; j++ {
		for i := 28; i >= 0; i -= 4 {
			temp1 := BitChain[j] >> i
			temp2 := temp1 & 0b1111
			counter[temp2] += 1
		}
	}

	var sum float32
	for i := 0; i < 16; i++ {
		sum += float32(counter[i] * counter[i])
	}

	X = (16/float32(5000))*sum - float32(5000)

	if X > 1.03 && X < 57.4 {
		return true
	}
	return false
}

func RunsTest(BitChain []uint32) bool {
	return false
}

func main() {
	BitChain := GenRandomBitChain()
	fmt.Print("BitChain: ")
	fmt.Println(BitChain)

	fmt.Print("MonobitTest: ")
	fmt.Println(MonobitTest(BitChain))

	fmt.Print("LongRunsTest: ")
	fmt.Println(LongRunsTest(BitChain))

	fmt.Print("PokerTest: ")
	fmt.Println(PokerTest(BitChain))

	fmt.Print("RunsTest: ")
	fmt.Println(RunsTest(BitChain))
}
