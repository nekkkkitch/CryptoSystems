package protocol

import (
	"fmt"
	"testing"
)

func TestTrade(t *testing.T) {
	var p Protocol
	p.SetPublic(5, 23)
	a := 4
	b := 2
	A := p.GetBig(a)
	B := p.GetBig(b)
	Z1 := p.CalcZ(A, b)
	Z2 := p.CalcZ(B, a)
	fmt.Println(Z1, Z2)
}
