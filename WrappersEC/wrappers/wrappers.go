package wrappers

import (
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

var curve elliptic.Curve = elliptic.P256()

type ECPoint struct {
	X *big.Int
	Y *big.Int
}

// G-generator receiving
func BasePointGGet() (p ECPoint) {
	_, p.X, p.Y, _ = elliptic.GenerateKey(curve, rand.Reader) // _, point.X, point.Y, _ = private key, X, Y, error; де X,Y - координати точки на еліптичній криві
	return p
}

// ECPoint creation
func ECPointGen(x, y *big.Int) (p ECPoint) {
	p.X = x
	p.Y = y
	return p
}

// DOES P ∈ CURVE?
func IsOnCurveCheck(p ECPoint) (c bool) {
	return curve.IsOnCurve(p.X, p.Y)
}

// P + Q
func AddECPoints(a, b ECPoint) (c ECPoint) {
	c.X, c.Y = curve.Add(a.X, a.Y, b.X, b.Y)
	return c
}

// 2P
func DoubleECPoints(a ECPoint) (c ECPoint) {
	c.X, c.Y = curve.Double(a.X, a.Y)
	return c
}

// k * P
func ScalarMult(k *big.Int, a ECPoint) (c ECPoint) {
	c.X, c.Y = curve.ScalarMult(a.X, a.Y, k.Bytes())
	return c
}

// Serialize point
func ECPointToString(p ECPoint) (s string) {
	return p.X.String() + " " + p.Y.String()
}

// Deserialize point
func StringToECPoint(s string) (p ECPoint) {
	p.X = &big.Int{}
	p.Y = &big.Int{}
	str := strings.Split(s, " ")
	p.X.SetString(str[0], 0)
	p.Y.SetString(str[1], 10)
	return p
}

// Print point
func PrintECPoint(p ECPoint) {
	fmt.Println("X: " + p.X.String() + "\nY: " + p.Y.String())
}
