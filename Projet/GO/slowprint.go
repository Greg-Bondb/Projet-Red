package main

import (
	"time"

	"github.com/01-edu/z01"
)

func slow(s string) {
	for _, letter := range s {
		time.Sleep(1 * time.Nanosecond)
		z01.PrintRune(letter)
	}
}
