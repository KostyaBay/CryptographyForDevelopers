package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

func SetRandom(bits int) (*big.Int, error) {
	return rand.Prime(rand.Reader, bits)
}

func main() {

	//k*(d*G) = d*(k*G)

	G := BasePointGGet()
	k, _ := SetRandom(256)
	d, _ := SetRandom(256)

	H1 := ScalarMult(d, G)
	H2 := ScalarMult(k, H1)

	H3 := ScalarMult(k, G)
	H4 := ScalarMult(d, H3)

	fmt.Println(H2) // print H2
	fmt.Println(H4) // print H4

	result := strings.EqualFold(H2.X.String(), H4.X.String()) && strings.EqualFold(H2.Y.String(), H4.Y.String())
	if result == true {
		fmt.Println("True: H2 is equal H4!")
	} else {
		fmt.Println("False: H2 not equal H4!")
	}
}
