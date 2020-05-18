package main

import (
	"runtime"
	"time"

	"github.com/micmonay/keybd_event"
)

//===================================================================
// These functions should be compatible with Windows, Linux (and Mac?).
// I have only tested it on Windows though.
//===================================================================

func volumeUp() error {
	return keypress(keybd_event.VK_VOLUME_UP)
}

func volumeDown() error {
	return keypress(keybd_event.VK_VOLUME_DOWN)
}

func silence() error {
	return keypress(keybd_event.VK_VOLUME_MUTE)
}

func nextSong() error {
	return keypress(keybd_event.VK_MEDIA_NEXT_TRACK)
}

func prevSong() error {
	return keypress(keybd_event.VK_MEDIA_PREV_TRACK)
}

func pauseSong() error {
	return keypress(keybd_event.VK_MEDIA_PLAY_PAUSE)
}

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
