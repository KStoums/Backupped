package utils

import (
	"fmt"
	"time"
)

func ClearTerminal() {
	for i := 0; i < 50; i++ {
		fmt.Println("\n")
	}
	return
}

func ClearTerminalAndOpenFunc(t time.Duration, e string, m func()) {
	ClearTerminal()
	fmt.Println(e)
	time.Sleep(t)
	ClearTerminal()
	m()
}
