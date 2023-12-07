package Signature

import (
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

var bigOne = big.NewInt(1)
var bigNul = big.NewInt(0)

func GenKey(keyLength int) (pk1, pk2, pk3, sk *big.Int) {
	var p, g, x, y *big.Int

	if 2048 <= keyLength && keyLength <= 4096 {
		for true {
			p, _ = rand.Prime(rand.Reader, keyLength)
			g, _ = rand.Int(rand.Reader, p)
			x, _ = rand.Int(rand.Reader, new(big.Int).Sub(p, bigOne))

			if x.Cmp(bigOne) == 1 {
				break
			}
		}
	} else {
		return nil, nil, nil, nil
	}

	y = new(big.Int).Exp(g, x, p)

	return p, g, y, x
}

func DigitSign(mess string, p, g, x *big.Int) (ds1, ds2 *big.Int) {
	hash := Hashing(mess)
	var k, r, s *big.Int
	for true {
		k, _ = rand.Int(rand.Reader, new(big.Int).Sub(p, bigOne))

		if k.Cmp(new(big.Int).Sub(p, bigOne)) == -1 && k.Cmp(bigOne) == 1 && new(big.Int).GCD(nil, nil, k, new(big.Int).Sub(p, bigOne)).String() == "1" {
			break
		}
	}

	r = new(big.Int).Exp(g, k, p)
	s = new(big.Int).Mod(new(big.Int).Mul(
		new(big.Int).Sub(
			hash,
			new(big.Int).Mul(x, r)),
		new(big.Int).ModInverse(
			k,
			new(big.Int).Sub(p, bigOne))), new(big.Int).Sub(p, bigOne))

	return r, s
}

func Verification(mess string, p, g, y, r, s *big.Int) bool {
	if r.Cmp(p) == -1 && r.Cmp(bigNul) == 1 && s.Cmp(new(big.Int).Sub(p, bigOne)) == -1 && s.Cmp(bigNul) == 1 {
		hash := Hashing(mess)
		left := new(big.Int).Mod(new(big.Int).Mul(new(big.Int).Exp(y, r, p), new(big.Int).Exp(r, s, p)), p)
		right := new(big.Int).Exp(g, hash, p)
		if left.String() == right.String() {
			return true
		}
	}
	return false
}

func Hashing(mess string) *big.Int {
	hashMess := sha256.Sum256([]byte(mess))
	return new(big.Int).SetBytes(hashMess[:])
}
