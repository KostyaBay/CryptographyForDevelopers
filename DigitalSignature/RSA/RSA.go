package RSA

import (
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

func GenKey(keyLength int) (sk1, pk1, p2 *big.Int) {
	var p, q, n, d, e, fEuler *big.Int

	var bigOne = big.NewInt(1)

	for true {
		p, _ = rand.Prime(rand.Reader, keyLength)
		q, _ = rand.Prime(rand.Reader, keyLength)

		fEuler = new(big.Int).Mul(new(big.Int).Sub(p, bigOne), new(big.Int).Sub(q, bigOne))
		e, _ = rand.Int(rand.Reader, fEuler)

		if e.Cmp(fEuler) == -1 && e.Cmp(bigOne) == 1 && new(big.Int).GCD(nil, nil, e, fEuler).String() == "1" {
			break
		}
	}

	n = new(big.Int).Mul(p, q)
	d = new(big.Int).ModInverse(e, fEuler)

	return d, e, n
}

func DigitSign(mess string, e, n *big.Int) *big.Int {
	hash := Hashing(mess)
	return new(big.Int).Exp(hash, e, n)
}

func Verification(mess string, c, d, n *big.Int) bool {
	hashMess := new(big.Int).Exp(c, d, n)
	hash := Hashing(mess)
	if hashMess.String() == hash.String() {
		return true
	} else {
		return false
	}
}

func Hashing(mess string) *big.Int {
	hashMess := sha256.Sum256([]byte(mess))
	return new(big.Int).SetBytes(hashMess[:])
}
