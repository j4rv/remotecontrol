package main

import (
	"os/exec"
	"strconv"
)

//===================================================
// These functions are only compatible with Windows.
//===================================================

func shutdownInSecs(seconds int) error {
	_ = abortShutdown() // We don't care if this fails
	c := exec.Command("shutdown", "-s", "-t", strconv.Itoa(seconds))
	return c.Run()
}

func abortShutdown() error {
	c := exec.Command("shutdown", "-a")
	return c.Run()
}
