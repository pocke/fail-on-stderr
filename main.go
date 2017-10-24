package main

import (
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	args := os.Args
	c := exec.Command(args[1], args[2:]...)
	c.Stdout = os.Stdout
	c.Stdin = os.Stdin
	stderr := bytes.NewBuffer([]byte{})
	c.Stderr = io.MultiWriter(stderr, os.Stderr)

	err := c.Run()
	status := 0

	if stderr.Len() != 0 {
		status = 1
	} else {
		if err != nil {
			if e2, ok := err.(*exec.ExitError); ok {
				if s, ok := e2.Sys().(syscall.WaitStatus); ok {
					status = s.ExitStatus()
				} else {
					panic(errors.New("Unimplemented for system where exec.ExitError.Sys() is not syscall.WaitStatus."))
				}
			}
		}
	}
	os.Exit(status)
}
