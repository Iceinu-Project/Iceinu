package util

import "fmt"

func ANSI(color int) string {
	return fmt.Sprintf("\x1b[%dm", color)
}
