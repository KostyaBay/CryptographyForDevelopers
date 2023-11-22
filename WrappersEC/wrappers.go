package main

import (
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

var curve elliptic.Curve = elliptic.P256()

type ECPoint struct {
	X *big.Int
	Y *big.Int
}

// G-generator receiving
func BasePointGGet() (point ECPoint) {
	_, point.X, point.Y, _ = elliptic.GenerateKey(curve, rand.Reader) // _, point.X, point.Y, _ = private key, X, Y, error; де X,Y - координати точки на еліптичній криві
	return point
}

// ECPoint creation
func ECPointGen(x, y *big.Int) (point ECPoint) {
	return
}
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
