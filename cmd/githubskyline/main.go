package main

import (
	"fmt"
)

func main() {
	
	st, err := GetSkyline("encse", 120)
	if err == nil {
		fmt.Println(st)
	} else {
		fmt.Println(err)
	}
}
