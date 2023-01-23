package io

import (
	"os"

	"golang.org/x/sys/unix"
)

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
