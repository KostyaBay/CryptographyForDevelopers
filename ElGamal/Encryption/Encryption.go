package Encryption

import (
	"crypto/rand"
	"math/big"
)

var bigOne = big.NewInt(1)

func Encrypt(mess string, p, g, y *big.Int) (c1, c2 *big.Int) {
	var a, b, k *big.Int
	k, _ = rand.Int(rand.Reader, new(big.Int).Sub(p, bigOne))

	if k.Cmp(new(big.Int).Sub(p, bigOne)) == -1 && k.Cmp(bigOne) == 1 {
		a = new(big.Int).Exp(g, k, p)
		b = new(big.Int).Mod(new(big.Int).Mul(new(big.Int).Exp(y, k, p), new(big.Int).SetBytes([]byte(mess))), p)
	}

	return a, b
}

func Decrypt(a, b, p, x *big.Int) string {
	M := new(big.Int).Mod(new(big.Int).Mul(b, new(big.Int).ModInverse(new(big.Int).Exp(a, x, p), p)), p)
	//M := new(big.Int).Mod(new(big.Int).Mul(b, new(big.Int).Exp(a, new(big.Int).Sub(new(big.Int).Sub(p, bigOne), x), p)), p)
	return string(M.Bytes())
}
