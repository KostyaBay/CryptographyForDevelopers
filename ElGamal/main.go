package main

import (
	"ElGamal/Signature"
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var bytes []byte
	var mess1 string

	fmt.Print("- Enter the message: ")
	bytes, _, _ = reader.ReadLine()
	mess1 = string(bytes)

	var length int
	fmt.Print("- Enter key length (128/256/512/1024 hex bytes) for GenKey: ")
	fmt.Scanln(&length)

	d, e, n := Signature.GenKey(length)
	fmt.Println("- Private Key: "+d.Text(16)+" ; "+n.Text(16), "\n- Public Key: "+e.Text(16)+" ; "+n.Text(16))

	fmt.Println("Message:", mess1)
}
