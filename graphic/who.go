package graphic

import (
	"projet/ascii"
	"projet/utils"
	"projet/values"
)

func DisplayWho() {
	utils.ClearTerminal()

	utils.Writeln("")
	utils.Writeln("")
	utils.Writeln(ascii.MenuDogsTitle)

	utils.Writeln(utils.CenterText("========= Qui sont-ils ? =========", 75))
	switch values.CurrentOption {
	case 0:
		utils.WriteColor(utils.CenterText("> READY PLAYER ONE = SPIELBERG   |", 75), "yellow")
		utils.Writeln(utils.CenterText("| ON AND ON AND ON = ABBA        |", 75))
		utils.Writeln(utils.CenterText("| Retour                         |", 75))
		utils.Writeln(utils.CenterText("==================================", 75))
		utils.Writeln(ascii.Mystery)
		break
	case 1:
		utils.Writeln(utils.CenterText("| READY PLAYER ONE = SPIELBERG   |", 75))
		utils.WriteColor(utils.CenterText("> ON AND ON AND ON = ABBA        |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Retour                         |", 75))
		utils.Writeln(utils.CenterText("==================================", 75))
		utils.Writeln(ascii.Mystery)
		break
	case 2:
		utils.Writeln(utils.CenterText("| READY PLAYER ONE = SPIELBERG   |", 75))
		utils.Writeln(utils.CenterText("| ON AND ON AND ON = ABBA        |", 75))
		utils.WriteColor(utils.CenterText("> Retour                         |", 75), "yellow")
		utils.Writeln(utils.CenterText("==================================", 75))
		utils.Writeln(ascii.ExitDoor)
		break
	}
}