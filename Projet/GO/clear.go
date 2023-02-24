package main

import "os"

// clear le terminal //
func Clear() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
}
