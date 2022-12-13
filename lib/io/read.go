package io

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func Readline() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	st, err := reader.ReadString('\n')
	if err != nil {
		return st, err
	}
	return strings.TrimRight(st, "\n"), nil
}

func ReadPassword() (string, error) {
	res, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func ReadKey() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	r, _, err := reader.ReadRune()
	return string(r), err
}

func ReadOption(prompt, options string) (string, error) {
	for {
		fmt.Printf("Select an item [%v]:", options)
		key, err := Readline()
		if err != nil {
			return "", err
		}
		fmt.Println(len(key))
		if len(key) == 1 && strings.Contains(strings.ToLower(options), strings.ToLower(key)) {
			return strings.ToLower(key), nil
		}
	}
}
