package cipher

import (
	"fmt"
	"testing"
)

func TestRSA(t *testing.T) {
	r := RSA{}
	r.ChangeParams(103, 53)
	fmt.Println(r)
	cipheredText := r.Cipher("text")
	decipheredText := r.Decipher(cipheredText)
	fmt.Println(cipheredText)
	fmt.Println(decipheredText)
}
