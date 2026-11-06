package protocol

import "log/slog"

type Protocol struct {
	Modulus int
	Base    int
}

func (p *Protocol) SetPublic(b, m int) {
	p.Modulus = m
	p.Base = b
}

func (p *Protocol) GetBig(small int) int {
	return modPow(p.Base, small, p.Modulus)
}

func (p *Protocol) CalcZ(big, small int) int {
	z := modPow(big, small, p.Modulus)
	slog.Info("z", "z", z, "big", big, "small", small)
	return z
}

func modPow(num, pow, mod int) int {
	mul := num
	for range pow - 1 {
		num = (num * mul) % mod
	}
	return num
}
