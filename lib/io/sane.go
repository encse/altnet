package io

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

func RunWithSavedTerminalState(cb func() error) error {
	fd := int(os.Stdin.Fd())
	termios, err := unix.IoctlGetTermios(fd, ioctlReadTermios)
	if err != nil {
		return err
	}
	defer func() {
		unix.IoctlSetTermios(fd, ioctlWriteTermios, termios)
		fmt.Print("\033[?47l")                  // alternate screen buffer off
		fmt.Print("\033]1337;SetColumns=0\007") // reset columns to screen width
	}()

	return cb()
}

func Sane() error {

	fd := int(os.Stdin.Fd())

	termios, err := unix.IoctlGetTermios(fd, ioctlReadTermios)
	if err != nil {
		return err
	}

	termios.Lflag |= unix.ICANON | unix.ISIG | unix.ECHO |
		unix.IEXTEN | unix.ECHO | unix.ECHOE | unix.ECHOK

	termios.Lflag &^= unix.ECHONL | unix.NOFLSH | unix.TOSTOP

	return unix.IoctlSetTermios(fd, ioctlWriteTermios, termios)
}
