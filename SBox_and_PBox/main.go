package main

import (
	"strings"
)

// Реалізація алгоритму P-блоку (пряме та зворотне перетворення) на прикладі вхідних даних розміром 8 бітів
func PBox(text uint8) uint8 {
	var ptext uint8                             //ініціалізуємо змінну для результату прямого перетворення P-блоку
	posChange := []int8{2, 4, 1, 7, 5, 3, 0, 6} //ініціалізуємо масив ключів(перестановку) для нашого перетворення (key = 2 4 1 7 5 3 0 6)
	for i, v := range posChange {
		tempTobyte := text >> i            //робимо проміжну дію - зсув вправо на i бітів, щоб обрати біт, з яким будемо працювати
		tempCheck := tempTobyte & uint8(1) //застосовуємо побітову операцію & до нашого tempTobyte з 1 таким чином, щоб залишився лише виключно наш обраний біт
		ptext = ptext + tempCheck<<v       //переставимо обраний біт на v позицій вліво (на нову позицію за ключом) та додамо в нашу змінну-результату переставлений біт
	}
	return ptext
}

func PBoxInv(ptext uint8) uint8 {
	var InvText uint8                           //ініціалізуємо змінну для результату зворотнього перетворення P-блоку
	posChange := []int8{2, 4, 1, 7, 5, 3, 0, 6} //ініціалізуємо масив ключів(перестановку) для нашого зворотнього перетворення (key = 2 4 1 7 5 3 0 6)
	for i, v := range posChange {
		tempRetPBox := ptext >> v            //робимо проміжну дію - зсув вправо на v бітів, щоб обрати біт, з яким будемо працювати
		tempTobyte := tempRetPBox & uint8(1) //застосовуємо побітову операцію & до нашого tempRetPBox з 1 таким чином, щоб залишився лише виключно наш обраний біт
		InvText = InvText + tempTobyte<<i    //переставимо обраний біт на i позицій вліво (на оригінальну його позицію) та додамо в нашу змінну-результату переставлений біт
	}
	return InvText
}

// Реалізація алгоритму S-блоку (пряме та зворотне перетворення) на прикладі вхідних даних розміром 8 бітів
func SBox(text uint8) uint8 {
	var stext uint8 //ініціалізуємо змінну для результату прямого перетворення S-блоку
	//ініціалізуємо таблицю констант для реалізації S-блоку (згенерована самостійно)
	tableChange := []uint8{0b0110, 0b0101, 0b1100, 0b1010, 0b0001, 0b1110, 0b0111, 0b1001, 0b1011, 0b0000, 0b0011, 0b1101, 0b1000, 0b1111, 0b0100, 0b0010}
	text_p1 := text & uint8(15)      //побітова операція & з 1, щоб обрати 1-шу тетраду (молодші 4-біти)
	text_p2 := text >> 4             //побітовий зсув вправо на 4, щоб обрати 2-гу тетраду (старші 4-біти)
	text_p1 = tableChange[text_p1]   //оберемо для 1-ої тетради заміну з таблиці констант за індексом
	text_p2 = tableChange[text_p2]   //оберемо для 2-ої тетради заміну з таблиці констант за індексом
	stext = (text_p2 << 4) + text_p1 //переставимо старші біти на своє місце(зсув вліво на 4) та додамо молодші біти
	return stext
}

// Функція обчислення зворотної таблиці констант
func Inverse(table [16]uint8) [16]uint8 {
	var tableInverse [16]uint8 //ініціалізуємо масив (таблицю констант для зворотнього перетворення)
	for i, v := range table {
		tableInverse[v] = uint8(i) //обраховуємо обернені значення до констант (const->indexes, indexes->const)
	}
	return tableInverse
}

func SBoxInv(ptext uint8) uint8 {
	var stextInv uint8 //ініціалізуємо змінну для результату зворотнього перетворення S-блоку
	//ініціалізуємо таблицю констант для реалізації S-блоку (згенерована самостійно)
	tableChange := [16]uint8{0b0110, 0b0101, 0b1100, 0b1010, 0b0001, 0b1110, 0b0111, 0b1001, 0b1011, 0b0000, 0b0011, 0b1101, 0b1000, 0b1111, 0b0100, 0b0010}
	tableInv := Inverse(tableChange)    //знайдемо таблицю констант для зворотнього перетворення
	text_p1 := ptext & uint8(15)        //побітова операція & з 1, щоб обрати 1-шу тетраду (молодші 4-біти)
	text_p2 := ptext >> 4               //побітовий зсув вправо на 4, щоб обрати 2-гу тетраду (старші 4-біти)
	text_p1 = tableInv[text_p1]         //оберемо для 1-ої тетради заміну зі знайденої таблиці констант за індексом
	text_p2 = tableInv[text_p2]         //оберемо для 2-ої тетради заміну зі знайденої таблиці констант за індексом
	stextInv = (text_p2 << 4) + text_p1 //переставимо старші біти на своє місце(зсув вліво на 4) та додамо молодші біти
	return stextInv
}

// Реалізація алгоритму P-блоку (пряме та зворотне перетворення) на прикладі вхідних даних > 8 бітів
// Вирішив залишити в якості додаткових методів, хоч і не підходять до вимог нашого завдання
func PBox_no8bit(text string) string {
	textArr := strings.Split(text, "")
	ptext := make([]string, len(textArr))
	posChange := []int{2, 4, 1, 7, 5, 3, 0, 6} //key
	for i, v := range posChange {
		ptext[v] = textArr[i]
	}
	return strings.Join(ptext, "")
}

func PBox_no8bit_Inv(ptext string) string {
	textArr := strings.Split(ptext, "")
	InvText := make([]string, len(textArr))
	posChange := []int{2, 4, 1, 7, 5, 3, 0, 6} //key
	for i, v := range posChange {
		InvText[i] = textArr[v]
	}
	return strings.Join(InvText, "")
}

func main() {
	// Закоментуваний виклик додаткових методів, які не задовільняють умовам завдання
	//tstring := "decision"
	//tstring_p := PBox_no8bit(tstring)
	//println("String PBox: " + tstring_p)
	//println("String PboxInv: " + PBox_no8bit_Inv(tstring_p) + "\n")

	// Ініціалізація вхідного тексту, яке буде використовуватись для перетворень:
	text := uint8(129) //10000001

	//Виклик методів алгоритму P-блоку
	text_p := PBox(text)
	print("Pbox: ")
	println(text_p)
	print("PboxInv: ")
	println(PBoxInv(text_p))

	//Виклик методів алгоритму S-блоку
	text_s := SBox(text)
	print("\nSbox: ")
	println(text_s)
	print("SboxInv: ")
	println(SBoxInv(text_s))

}
