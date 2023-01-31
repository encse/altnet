package io

const (
	Escape      = "\033"
	Bold        = Escape + "[1m"
	Inverse     = Escape + "[7m"
	Reset       = Escape + "[0m"
	ClearScreen = Escape + "[2J"
	Home        = Escape + "[H"
)
