package RSA

import (
	"crypto/rand"
	"math/big"
)

func GenKey() {
	var n, g big.Int
	var bigOne = big.NewInt(1)
	//var bigNegOne = big.NewInt(-1)
	var fEuler = big.NewInt(0)
	//var g = big.NewInt(0)

	p, _ := rand.Prime(rand.Reader, 512)
	q, _ := rand.Prime(rand.Reader, 512)

	n.Mul(p, q)

	fEuler = new(big.Int).Mul(new(big.Int).Sub(p, bigOne), new(big.Int).Sub(q, bigOne))
	e, _ := rand.Int(rand.Reader, fEuler)
	if e.Cmp(fEuler) == -1 && e.Cmp(1) == 1 && g.GCD(e, fEuler) == 1 {
		E := e
	}

}

func Encrypt() {}

func Decrypt() {}

func ShareKey() {}
