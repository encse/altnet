package io

import (
	"bufio"
	"errors"
	"fmt"
	stdio "io"
	"os"
	"runtime"
	"strings"

	"github.com/encse/altnet/lib/log"
	"github.com/encse/altnet/lib/slices"
	"golang.org/x/term"
)

type UserFriendlyError struct {
	Err error
}

func (e UserFriendlyError) Error() string { return e.Err.Error() }
func (e UserFriendlyError) Unwrap() error { return e.Err }

func FatalIfError(err error, message ...any) {
	if errors.Is(err, stdio.EOF) {
		os.Exit(0)
	}

	if err != nil {
		if len(message) > 0 {
			fmt.Println(message...)
		} else if uerr, ok := err.(UserFriendlyError); ok {
			fmt.Println(uerr.Error())
		} else {
			fmt.Println("An error occurred.")
		}
		_, filename, line, _ := runtime.Caller(1)
		log.Errorf("[error] %s:%d %v", filename, line, err)
		os.Exit(1)
	}
}

func ReadArg(prompt string, args []string, iarg int) (string, error) {
	arg := ""
	if len(args) > iarg {
		arg = args[iarg]
	}

	if arg == "" {
		var err error
		arg, err = ReadNotEmpty(prompt + ":")
		if err != nil {
			return "", err
		}
	}
	return arg, nil
}

func ReadArgFromList(prompt string, args []string, iarg int, options []string) (string, error) {
	arg := ""
	if len(args) > iarg {
		arg = args[iarg]
	}

	if slices.Contains(options, arg) {
		return arg, nil
	}

	for {
		var err error
		arg, err = ReadNotEmpty(prompt + " (? for list): ")
		if err != nil {
			return "", err
		}

		if slices.Contains(options, arg) {
			return arg, err
		}

		if arg == "?" {
			for _, option := range options {
				fmt.Println(option)
			}
		}
	}
}

func ReadNotEmpty(prompt string) (string, error) {
	for {
		res, err := Readline(prompt)
		if err != nil {
			return "", err
		}
		res = strings.TrimSpace(res)
		if res != "" {
			return res, nil
		}
	}
}

func Readline(prompt string) (string, error) {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	st, err := reader.ReadString('\n')
	if err != nil {
		return st, err
	}
	return strings.TrimRight(st, "\n"), nil
}

func ReadPassword(prompt string) (string, error) {
	fmt.Print(prompt)
	res, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}
	fmt.Println()
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
		key, err := Readline(fmt.Sprintf("Select an item [%v]:", options))
		if err != nil {
			return "", err
		}
		if len(key) == 1 && strings.Contains(strings.ToLower(options), strings.ToLower(key)) {
			return strings.ToLower(key), nil
		}
	}
}
