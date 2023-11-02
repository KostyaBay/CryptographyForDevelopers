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

// Реалізація монобітного тесту
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

// Реалізація тесту максимальної довжини серії
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
				curCount = 1
			}
		}
	}

	if maxCount <= 36 {
		return true
	}
	return false
}

// Реалізація тесту Поккера
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

// Реалізація тесту довжин серій
func RunsTest(BitChain []uint32) bool {
	pointerBit := (BitChain[0] >> 31) & uint32(1)
	curCount := 0
	maxCount := 0
	var arrSeries [2][6]int
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
				if curCount > 6 {
					curCount = 6
				}
				arrSeries[pointerBit][curCount-1] += 1
				pointerBit = tempValBit
				curCount = 1
			}
		}
	}

	if curCount > maxCount {
		maxCount = curCount
	}

	if curCount > 6 {
		curCount = 6
	}
	arrSeries[pointerBit][curCount-1] += 1

	compare_intervals := [6][2]int{{2267, 2733}, {1079, 1421}, {502, 748}, {223, 402}, {90, 223}, {90, 223}} // [min][max] {{min, max}}

	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			if (arrSeries[i][j] < compare_intervals[i][1]) && (arrSeries[i][j] > compare_intervals[i][0]) {
				return true
			}
		}
	}

	return false
}

func main() {
	for {
		BitChain := GenRandomBitChain()

		fmt.Print("BitChain: ")
		fmt.Println(BitChain)

		fmt.Println("MonobitTest:", MonobitTest(BitChain), "|", "LongRunsTest:", LongRunsTest(BitChain), "|", "PokerTest:", PokerTest(BitChain), "|", "RunsTest:", RunsTest(BitChain))

		if MonobitTest(BitChain) && LongRunsTest(BitChain) && PokerTest(BitChain) && RunsTest(BitChain) {
			fmt.Println("This sequence of 20,000 bits is random.")
			break
		} else {
			fmt.Println("This sequence of 20,000 bits is not random.")
		}
	}
}
