package main

import (
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

func keypress(keys ...int) error {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return err
	}

	// For linux, it is very important wait 2 seconds
	// (???) check the keybd_event repository
	if runtime.GOOS == "linux" {
		time.Sleep(2 * time.Second)
	}

	kb.SetKeys(keys...)
	return kb.Launching()
}
