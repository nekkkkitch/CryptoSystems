package cipher

import (
	op "2/backend/services/operations"
	"log/slog"
	"math/rand"
)

type RSA struct {
	Public  int
	Private int
	N       int
}

func (r *RSA) ChangeParams(P, Q int) (int, int) {
	r.N = P * Q
	eul := (P - 1) * (Q - 1)
	slog.Info("eul", "eul", eul)
	for i := 2; i < eul; i++ {
		if op.IsRelativlyPrime(i, eul) && op.Inverse(i, eul) > 0 {
			r.Public = i
			if rand.Intn(2) == 1 {
				break
			}
		}
	}
	r.Private = op.Inverse(r.Public, eul)
	return r.Public, r.Private
}

func (r *RSA) Cipher(text string) []int32 {
	msg := []rune(text)
	slog.Info("Message before cipher", "msg", msg)
	for i := range msg {
		msg[i] = rune(op.ModPow(int(msg[i]), r.Public, r.N))
	}
	slog.Info("Message after cipher", "msg", msg)
	return msg
}

func (r *RSA) Decipher(text []int32) string {
	msg := []rune(text)

	slog.Info("Message before decipher", "msg", msg)
	for i := range msg {
		msg[i] = rune(op.ModPow(int(msg[i]), r.Private, r.N))
	}
	slog.Info("Message after decipher", "msg", msg)
	return string(msg)
}
