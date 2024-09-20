package graphic

import (
	"fmt"
	"os"
	"projet/ascii"
	"projet/utils"
	"projet/values"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func DisplayNiche() {
	utils.ClearTerminal()

	utils.Writeln("")
	utils.Writeln("")
	utils.Writeln(ascii.MenuDogsTitle)

	utils.Writeln(utils.CenterText("====== Voir la niche ======", 75))
	switch values.CurrentOption {
	case 0:
		utils.WriteColor(utils.CenterText("> Voir le chien           |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Changer de chien        |", 75))
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 1:
		utils.Writeln(utils.CenterText("| Voir le chien           |", 75))
		utils.WriteColor(utils.CenterText("> Changer de chien        |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 2:
		utils.Writeln(utils.CenterText("| Voir le chien           |", 75))
		utils.Writeln(utils.CenterText("| Changer de chien        |", 75))
		utils.WriteColor(utils.CenterText("> Retour                  |", 75), "yellow")
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	}

	utils.DisplayTwoAsciiSideBySide(ascii.Niches[values.Player.GetHouse()], ascii.DogsArt[values.Player.GetCurrentDog().Race], 5)
}

func DisplayDogInfo() {
	utils.ClearTerminal()

	utils.Writeln("")
	utils.Writeln("")
	utils.Writeln(ascii.MenuDogsTitle)

	Dog := values.Player.GetCurrentDog()

	data := [][]string{
		[]string{"Nom : " + Dog.Name + "\nVie : " + strconv.Itoa(Dog.Health) + "/" + strconv.Itoa(Dog.MaxHealth) + "\nDégat min/max : " + strconv.Itoa(Dog.Damage) + "/" + strconv.Itoa(Dog.MaxDamage) + "\nRace : " + Dog.Race, ascii.DogsArt[Dog.Race]},
		//[]string{"\n\n", "\n\n"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	for _, v := range data {
		table.Append(v)
	}
	table.Render()

	switch values.CurrentOption {
	case 0:
		if values.Player.Inventory.Heal > 0 {
			utils.WriteColor(utils.CenterText("> Soigner le chien", 2), "yellow")
		} else {
			utils.WriteColor(utils.CenterText("> Soigner le chien", 2), "yellow")
			utils.WriteColor(utils.CenterText("   [x] vous n'avez pas de soin", 2), "red")
		}
		utils.Writeln(utils.CenterText("Entrainer le chien", 2))
		utils.Writeln(utils.CenterText("Gérer les compétences", 2))
		utils.Writeln(utils.CenterText("Retour", 2))
		break
	case 1:
		utils.Writeln(utils.CenterText("Soigner le chien", 2))
		utils.WriteColor(utils.CenterText("> Entrainer le chien", 2), "yellow")
		utils.Writeln(utils.CenterText("Gérer les compétences", 2))
		utils.Writeln(utils.CenterText("Retour", 2))
		break
	case 2:
		utils.Writeln(utils.CenterText("Soigner le chien", 2))
		utils.Writeln(utils.CenterText("Entrainer le chien", 2))
		utils.WriteColor(utils.CenterText("> Gérer les compétences", 2), "yellow")
		utils.Writeln(utils.CenterText("Retour", 2))
		break
	case 3:
		utils.Writeln(utils.CenterText("Soigner le chien", 2))
		utils.Writeln(utils.CenterText("Entrainer le chien", 2))
		utils.Writeln(utils.CenterText("Gérer les compétences", 2))
		utils.WriteColor(utils.CenterText("> Retour", 2), "yellow")
	}
}

func DiplayChangeDog() {
	utils.ClearTerminal()

	utils.Writeln("")
	utils.Writeln("")
	utils.Writeln(ascii.MenuDogsTitle)

	utils.Writeln(utils.CenterText("====== Changer de chien ======", 75))
	for index, dog := range values.Player.Dogs {
		if index == values.CurrentOption {
			utils.WriteColor(utils.CenterText("> "+dog.Name, 75), "yellow")
		} else {
			utils.Writeln(utils.CenterText(dog.Name, 75))
		}
	}
	if values.CurrentOption == len(values.Player.Dogs) {
		utils.WriteColor(utils.CenterText("> Acheter un chien", 75), "yellow")
	} else {
		utils.Writeln(utils.CenterText("Acheter un chien", 75))
	}
	if values.CurrentOption == len(values.Player.Dogs)+1 {
		utils.WriteColor(utils.CenterText("> Retour", 75), "yellow")
		utils.Writeln(utils.CenterText(ascii.ExitDoor, 75))
	} else {
		utils.Writeln(utils.CenterText("Retour", 75))
		if values.CurrentOption < len(values.Player.Dogs) {
			utils.Writeln(utils.CenterText(ascii.DogsArt[values.Player.GetCurrentDog().Race], 75))
		} else {
			utils.Writeln(utils.CenterText(ascii.SPA, 75))
		}
	}
}

func DisplayBuyDog() {
	utils.ClearTerminal()

	utils.Writeln("")
	utils.Writeln("")
	//utils.Writeln(ascii.MenuDogsTitle)

	utils.Writeln(utils.CenterText("====== Acheter un chien ======", 75))
	for index, dog := range values.DogsRace {
		dogPrice := values.DogsPrice[dog]
		if index == values.CurrentOption {
			utils.WriteColor(utils.CenterText("> "+dog+" - "+strconv.Itoa(dogPrice)+"$", 75), "yellow")
		} else {
			utils.Writeln(utils.CenterText(dog+" - "+strconv.Itoa(dogPrice)+"$", 75))
		}
	}
	if values.CurrentOption == len(values.DogsRace) {
		utils.WriteColor(utils.CenterText("> Retour", 75), "yellow")
		utils.Writeln(utils.CenterText(ascii.ExitDoor, 75))
	} else {
		utils.Writeln(utils.CenterText("Retour", 75))
		utils.Writeln(utils.CenterText(ascii.DogsArt[values.DogsRace[values.CurrentOption]], 75))
	}
}

func DisplayCompetences() {
	utils.ClearTerminal()

	utils.Writeln("")
	utils.Writeln("")
	utils.Writeln(ascii.MenuDogsTitle)

	utils.Writeln(utils.CenterText("======== Compétences =======", 75))
	Dog := values.Player.Dogs[values.Player.CurrentDog]
	switch values.CurrentOption {
	case 0:
		utils.WriteColor(utils.CenterText(fmt.Sprintf("> %-24s |", strconv.Itoa(Dog.MaxHealth)+" Vie max"), 75), "yellow")
		utils.Writeln(utils.CenterText(fmt.Sprintf("| %-24s |", strconv.Itoa(Dog.Damage)+" Dégat min"), 75))
		utils.Writeln(utils.CenterText(fmt.Sprintf("| %-24s |", strconv.Itoa(Dog.MaxDamage)+" Dégat max"), 75))
		utils.Writeln(utils.CenterText(fmt.Sprintf("| %-24s |", strconv.Itoa(Dog.DodgeMultiplier)+" Bonus esquive"), 75))
		utils.Writeln(utils.CenterText("| Retour                   |", 75))
		utils.Writeln(utils.CenterText("============================", 75))
		break
	case 1:
		utils.Writeln(utils.CenterText(fmt.Sprintf("| %-24s |", strconv.Itoa(Dog.MaxHealth)+" Vie max"), 75))
		utils.WriteColor(utils.CenterText(fmt.Sprintf("> %-24s |", strconv.Itoa(Dog.Damage)+" Dégat min"), 75), "yellow")
		utils.Writeln(utils.CenterText(fmt.Sprintf("| %-24s |", strconv.Itoa(Dog.MaxDamage)+" Dégat max"), 75))
		utils.Writeln(utils.CenterText(fmt.Sprintf("| %-24s |", strconv.Itoa(Dog.DodgeMultiplier)+" Bonus esquive"), 75))
		utils.Writeln(utils.CenterText("| Retour                   |", 75))
		utils.Writeln(utils.CenterText("============================", 75))
		break
	case 2:
		utils.Writeln(utils.CenterText(fmt.Sprintf("| %-24s |", strconv.Itoa(Dog.MaxHealth)+" Vie max"), 75))
		utils.Writeln(utils.CenterText(fmt.Sprintf("| %-24s |", strconv.Itoa(Dog.Damage)+" Dégat min"), 75))
		utils.WriteColor(utils.CenterText(fmt.Sprintf("> %-24s |", strconv.Itoa(Dog.MaxDamage)+" Dégat max"), 75), "yellow")
		utils.Writeln(utils.CenterText(fmt.Sprintf("| %-24s |", strconv.Itoa(Dog.DodgeMultiplier)+" Bonus esquive"), 75))
		utils.Writeln(utils.CenterText("| Retour                   |", 75))
		utils.Writeln(utils.CenterText("============================", 75))
		break
	case 3:
		utils.Writeln(utils.CenterText(fmt.Sprintf("| %-24s |", strconv.Itoa(Dog.MaxHealth)+" Vie max"), 75))
		utils.Writeln(utils.CenterText(fmt.Sprintf("| %-24s |", strconv.Itoa(Dog.Damage)+" Dégat min"), 75))
		utils.Writeln(utils.CenterText(fmt.Sprintf("| %-24s |", strconv.Itoa(Dog.MaxDamage)+" Dégat max"), 75))
		utils.WriteColor(utils.CenterText(fmt.Sprintf("> %-24s |", strconv.Itoa(Dog.DodgeMultiplier)+" Bonus esquive"), 75), "yellow")
		utils.Writeln(utils.CenterText("| Retour                   |", 75))
		utils.Writeln(utils.CenterText("============================", 75))
		break
	case 4:
		utils.Writeln(utils.CenterText(fmt.Sprintf("| %-24s |", strconv.Itoa(Dog.MaxHealth)+" Vie max"), 75))
		utils.Writeln(utils.CenterText(fmt.Sprintf("| %-24s |", strconv.Itoa(Dog.Damage)+" Dégat min"), 75))
		utils.Writeln(utils.CenterText(fmt.Sprintf("| %-24s |", strconv.Itoa(Dog.MaxDamage)+" Dégat max"), 75))
		utils.Writeln(utils.CenterText(fmt.Sprintf("| %-24s |", strconv.Itoa(Dog.DodgeMultiplier)+" Bonus esquive"), 75))
		utils.WriteColor(utils.CenterText("> Retour                   |", 75), "yellow")
		utils.Writeln(utils.CenterText("============================", 75))
		break
	}

	utils.Writeln(utils.CenterText("Points exprérience : "+strconv.Itoa(Dog.Experience), 75))
}
