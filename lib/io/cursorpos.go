package io

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"golang.org/x/sys/unix"
)

type Position struct {
	Row    int
	Column int
}

func GetCursorPosition() (Position, error) {
	cmdCursorPosition := escape + "[6n"
	st, err := read(cmdCursorPosition)
	if err != nil {
		return Position{}, fmt.Errorf("Could not get cursor postion, %w", err)
	}

	e := escape + "\\[(\\d+);(\\d+)R"
	r := regexp.MustCompile(e)

	m := r.FindStringSubmatchIndex(st)
	if m == nil {
		return Position{}, errors.New("Could not get cursor postion, invalid format")
	}

	row, err := strconv.Atoi(st[m[2]:m[3]])
	if err != nil {
		return Position{}, errors.New("Could not get cursor postion, invalid format")
	}
	col, err := strconv.Atoi(st[m[4]:m[5]])
	if err != nil {
		return Position{}, errors.New("Could not get cursor postion, invalid format")
	}

	return Position{Row: row, Column: col}, nil

}
func read(cmd string) (string, error) {

	fd := int(os.Stdin.Fd())

	termios, err := unix.IoctlGetTermios(fd, ioctlReadTermios)
	if err != nil {
		return "", err
	}

	newState := *termios
	newState.Lflag = newState.Lflag &^ unix.ECHO
	newState.Lflag = newState.Lflag &^ unix.ICANON
	if err := unix.IoctlSetTermios(fd, ioctlWriteTermios, &newState); err != nil {
		return "", err
	}

	defer unix.IoctlSetTermios(fd, ioctlWriteTermios, termios)

	fmt.Print(cmd)

	buf := make([]byte, 100)
	n, err := os.Stdin.Read(buf)
	if err != nil {
		panic(err)
	}

	buf = buf[:n]
	return string(buf), nil
}
