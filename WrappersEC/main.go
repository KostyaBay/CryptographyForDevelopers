package main

import (
	"crypto/elliptic"
	"fmt"
	"math/big"
)

type ECPoint struct {
	X *big.Int
	Y *big.Int
}

func BasePointGGet() (point ECPoint) {
	elliptic.GenerateKey()
	return
} //G-generator receiving

func ECPointGen(x, y *big.Int) (point ECPoint) {
	return
} //ECPoint creation
func IsOnCurveCheck(a ECPoint) (c bool) {
	return
} //DOES P ∈ CURVE?
func AddECPoints(a, b ECPoint) (c ECPoint) {
	return
} //P + Q
func DoubleECPoints(a ECPoint) (c ECPoint) {
	return
} //2P
func ScalarMult(k big.Int, a ECPoint) (c ECPoint) {
	return
} //k * P
func ECPointToString(point ECPoint) (s string) {
	return
} //Serialize point
func StringToECPoint(s string) (point ECPoint) {
	return
} //Deserialize point
func PrintECPoint(point ECPoint) {
	return
} //Print point

func main() {
	fmt.Println("Hello, World!")
}
