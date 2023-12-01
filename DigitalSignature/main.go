package main

import (
	"DigitalSignature/RSA"
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func TestDigSign() {
	reader := bufio.NewReader(os.Stdin)
	var bytes []byte

	var length int
	fmt.Print("- Enter key length (128/256/512/1024 hex bytes) for GenKey: ")
	fmt.Scanln(&length)

	d, e, n := RSA.GenKey(length)
	fmt.Println("- Private Key: "+d.Text(16)+" ; "+n.Text(16), "\n- Public Key: "+e.Text(16)+" ; "+n.Text(16)) // "\nn:", n.Text(16))

	var mess1 string
	fmt.Print("- Enter the message: ")
	bytes, _, _ = reader.ReadLine()
	mess1 = string(bytes)

	var sk1, pk1, p2 string
	fmt.Print("- Enter 1 part of your Private Key(d): ")
	fmt.Scanln(&sk1)

	fmt.Print("- Enter 1 part of Public Key(e): ")
	fmt.Scanln(&pk1)

	fmt.Print("- Enter 2 part of keys(n): ")
	fmt.Scanln(&p2)

	e_pk1, _ := new(big.Int).SetString(pk1, 16)
	n_p2, _ := new(big.Int).SetString(p2, 16)
	d_sk1, _ := new(big.Int).SetString(sk1, 16)

	c := RSA.DigitSign(mess1, e_pk1, n_p2)
	fmt.Println("- Your Digital Signature:", c.Text(16))

	var ds string
	fmt.Print("- Enter your Digital Signature: ")
	fmt.Scanln(&ds)

	var mess2 string
	fmt.Print("- Enter the message: ")
	bytes, _, _ = reader.ReadLine()
	mess2 = string(bytes)

	cInt, _ := new(big.Int).SetString(ds, 16)

	dec := RSA.Verification(mess2, cInt, d_sk1, n_p2)
	fmt.Println("- Verification of DS is", dec, "!")
}

func main() {
	TestDigSign()
}
