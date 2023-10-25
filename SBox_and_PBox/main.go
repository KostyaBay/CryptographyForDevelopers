package main

import (
	"strings"
)

// Реалізація алгоритму P-блоку (пряме та зворотне перетворення) на прикладі вхідних даних розміром 8 бітів
func PBox(text uint8) uint8 {
	var ptext uint8
	posChange := []int8{2, 4, 1, 7, 5, 3, 0, 6} //key = 2 4 1 7 5 3 0 6
	for i, v := range posChange {
		tempTobyte := text >> i
		tempCheck := tempTobyte & uint8(1)
		ptext = ptext + tempCheck<<v
	}
	return ptext
}

func PBoxInv(ptext uint8) uint8 {
	var InvText uint8
	posChange := []int8{2, 4, 1, 7, 5, 3, 0, 6}
	for i, v := range posChange {
		tempRetPBox := ptext >> v
		tempTobyte := tempRetPBox & uint8(1)
		InvText = InvText + tempTobyte<<i
	}
	return InvText
}

// Реалізація алгоритму S-блоку (пряме та зворотне перетворення) на прикладі вхідних даних розміром 8 бітів
func SBox(text uint8) uint8 {
	var ptext uint8
	posChange := []int8{2, 4, 1, 7, 5, 3, 0, 6} //key = 2 4 1 7 5 3 0 6
	for i, v := range posChange {
		tempTobyte := text >> i
		tempCheck := tempTobyte & uint8(1)
		ptext = ptext + tempCheck<<v
	}
	return ptext
}

func SBoxInv(ptext uint8) uint8 {
	var InvText uint8
	posChange := []int8{2, 4, 1, 7, 5, 3, 0, 6}
	for i, v := range posChange {
		tempRetPBox := ptext >> v
		tempTobyte := tempRetPBox & uint8(1)
		InvText = InvText + tempTobyte<<i
	}
	return InvText
}

// Реалізація алгоритму P-блоку (пряме та зворотне перетворення) на прикладі вхідних даних > 8 бітів
// Вирішив залишити в якості додаткових методів, хоч і не підходять до вимог нашого завдання
func PBox_no8bit(text string) string {
	textArr := strings.Split(text, "")
	ptext := make([]string, len(textArr))
	posChange := []int{2, 4, 1, 7, 5, 3, 0, 6}
	for i, v := range posChange {
		ptext[v] = textArr[i]
	}
	return strings.Join(ptext, "")
}

func PBox_no8bit_Inv(ptext string) string {
	textArr := strings.Split(ptext, "")
	InvText := make([]string, len(textArr))
	posChange := []int{2, 4, 1, 7, 5, 3, 0, 6}
	for i, v := range posChange {
		InvText[i] = textArr[v]
	}
	return strings.Join(InvText, "")
}

func main() {
	tstring := "decision"
	tstring_p := PBox_no8bit(tstring)
	println("String PBox: " + tstring_p)
	println("String PboxInv: " + PBox_no8bit_Inv(tstring_p) + "\n")

	text := uint8(129) //10000001
	text_p := PBox(text)
	print("Pbox: ")
	println(text_p)
	print("PboxInv: ")
	println(PBoxInv(text_p))

	text_s := SBox(text)
	print("\n Sbox: ")
	println(text_s)
	//print("SboxInv: ")
	//println(SBoxInv(text_s))

}
