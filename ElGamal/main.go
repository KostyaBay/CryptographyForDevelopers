package main

import (
	"ElGamal/Encryption"
	"ElGamal/Signature"
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func Test_KeyGen(length int) (p, g, y, x *big.Int) {
	println("*** KEY GENERATION TEST ***")
	p, g, y, x = Signature.GenKey(length)
	fmt.Println("- Private Key(x): "+x.Text(16), "\n- Public Key(p;g;y): "+p.Text(16)+" ; "+g.Text(16)+" ; "+y.Text(16))
	return p, g, y, x
}

func Test_DigitSign(p, g, y, x *big.Int) {
	println("*** DIGIT SIGNATURE TEST ***")
	reader := bufio.NewReader(os.Stdin)
	var bytes []byte

	var mess1 string
	fmt.Print("- Enter the message: ")
	bytes, _, _ = reader.ReadLine()
	mess1 = string(bytes)

	r, s := Signature.DigitSign(mess1, p, g, x)
	fmt.Println("- Your Digital Signature(r;s):", r.Text(16)+" ; "+s.Text(16))

	var mess2 string
	fmt.Print("- Enter the message: ")
	bytes, _, _ = reader.ReadLine()
	mess2 = string(bytes)

	dec := Signature.Verification(mess2, p, g, y, r, s)
	fmt.Println("- Verification of DS is", dec, "!")
}

func Test_Encrypt(p, g, y, x *big.Int) {
	println("*** ENCRYPT/DECRYPT TEST ***")
	reader := bufio.NewReader(os.Stdin)
	var bytes []byte

	var mess string
	fmt.Print("- Enter the message: ")
	bytes, _, _ = reader.ReadLine()
	mess = string(bytes)

	a, b := Encryption.Encrypt(mess, p, g, y)
	fmt.Println("- Encrypt your message(a;b): " + a.Text(16) + " ; " + b.Text(16))

	M := Encryption.Decrypt(a, b, p, x)
	fmt.Println("- Decrypt your message:", M)
}

func main() {
	var length int
	fmt.Print("- Enter key length (from 2048 to 4096) for GenKey: ")
	fmt.Scanln(&length)

	p, g, y, x := Test_KeyGen(length)

	Test_DigitSign(p, g, y, x)
	Test_Encrypt(p, g, y, x)
}
