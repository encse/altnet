package io

import (
	"fmt"

	"github.com/encse/altnet/lib/log"
)

func Freshline() {
	pos, err := GetCursorPosition()
	if err != nil {
		log.Error(err)
		fmt.Println()
	}
	if pos.Column != 1 {
		fmt.Println(Bold + Inverse + "%" + Reset)
	}
}
