package menu

import (
	"Tolnkee-Backup-Test/backup"
	"Tolnkee-Backup-Test/messages"
	"Tolnkee-Backup-Test/utils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type OptionMainMenu struct {
	number     int
	nameOption string
	function   func()
}

var optionMainMenuList []*OptionMainMenu

func init() {
	optionMainMenuList = []*OptionMainMenu{
		{
			number:     1,
			nameOption: "Full OS Backup",
			function: func() {
				inputResponse := backup.FullBackup()
				fmt.Println(inputResponse)
				if strings.EqualFold(inputResponse, "y") {
					utils.ClearTerminal()
					MainMenu()
					return
				}

				if strings.EqualFold(inputResponse, "n") {
					utils.ClearTerminal()
					fmt.Print(messages.EXIT_APP)
					time.Sleep(3 * time.Second)
					return
				}
			},
		},
		{
			number:     2,
			nameOption: "Backup of a specific directory",
			function: func() {
				utils.ClearTerminalAndOpenFunc(3*time.Second, messages.ERROR_FEATURE_SOON, MainMenu)
				return
			},
		},
		{
			number:     3,
			nameOption: "Exit",
			function: func() {
				utils.ClearTerminal()
				fmt.Print(messages.EXIT_APP)
				time.Sleep(3 * time.Second)
				return
			},
		},
	}
}

func MainMenu() {
	fmt.Println("╔══════════════════════════════════╗" +
		"\n" +
		"|             ᗷᗩᑕKᑌᑭᑭEᗪ            |" +
		"\n" +
		"|              кѕтαяѕ℠             |" +
		"\n" +
		"╚══════════════════════════════════╝")

	for i, menuOption := range optionMainMenuList {
		if i == 0 {
			fmt.Println(fmt.Sprintf("╔-[%d] %s", menuOption.number, menuOption.nameOption))
		} else if i == len(optionMainMenuList)-1 {
			fmt.Print(fmt.Sprintf("╚-[%d] %s\n\n➤ ", menuOption.number, menuOption.nameOption))
		} else {
			fmt.Println(fmt.Sprintf("╠-[%d] %s", menuOption.number, menuOption.nameOption))
		}
	}

	var input string
	fmt.Scanln(&input)

	inputInt, err := strconv.Atoi(input)
	if err != nil {
		utils.ClearTerminalAndOpenFunc(3*time.Second, messages.ERROR_STRING_CONVERTION_SYNTAX, MainMenu)
		return
	}

	for _, menuOption := range optionMainMenuList {
		if inputInt == menuOption.number {
			utils.ClearTerminal()
			menuOption.function()
			return
		}
	}

	utils.ClearTerminalAndOpenFunc(3*time.Second, messages.ERROR_STRING_CONVERTION_SYNTAX, MainMenu)
	return
}
