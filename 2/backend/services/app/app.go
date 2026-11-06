package app

import (
	"context"
	"strconv"
	"strings"
)

type RSA interface {
	ChangeParams(P, Q int) (int, int)
	Decipher(text []int32) string
	Cipher(text string) []int32
}

// App struct
type App struct {
	r   RSA
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp(r RSA) *App {
	app := &App{}
	app.r = r
	return app
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) Cipher(text string) []int32 {
	toRet := a.r.Cipher(text)
	return toRet
}

func (a *App) Decipher(cipher string) string {
	sliced := strings.Split(cipher, ",")
	toSend := make([]int32, 0, len(sliced))
	for i := range sliced {
		num, _ := strconv.Atoi(sliced[i])
		toSend = append(toSend, int32(num))
	}
	return a.r.Decipher(toSend)
}

type KeysToRet struct {
	Pu int `json:"pu"`
	Pr int `json:"pr"`
}

func (a *App) ChangeParams(P, Q int) KeysToRet {
	pu, pr := a.r.ChangeParams(P, Q)
	return KeysToRet{Pu: pu, Pr: pr}
}
