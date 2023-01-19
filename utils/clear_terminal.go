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

func ClearTerminalAndOpenFunc(t time.Duration, msg string, function func()) {
	ClearTerminal()
	fmt.Println(msg)
	time.Sleep(t)
	ClearTerminal()
	function()
}
