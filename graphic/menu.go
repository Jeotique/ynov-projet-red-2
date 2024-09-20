package graphic

import (
	"projet/ascii"
	"projet/utils"
	"projet/values"
)

func RefreshMainMenu() {
	utils.ClearTerminal()

	utils.Writeln("")
	utils.Writeln("")
	utils.Writeln(ascii.MenuTitle)

	utils.Writeln(utils.CenterText("=== Menu du jeu ===", 75))
	switch values.CurrentOption {
	case 0:
		utils.WriteColor(utils.CenterText("> Jouer           |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Quitter         |", 75))
		utils.Writeln(utils.CenterText("===================", 75))
		break
	case 1:
		utils.Writeln(utils.CenterText("| Jouer           |", 75))
		utils.WriteColor(utils.CenterText("> Quitter         |", 75), "yellow")
		utils.Writeln(utils.CenterText("===================", 75))
		break
	}
}

func RefreshRealMainMenu() {
	utils.ClearTerminal()

	utils.Writeln("")
	utils.Writeln("")
	utils.Writeln(ascii.MenuDogsTitle)

	utils.Writeln(utils.CenterText("===== Le vrai de vrai =====", 75))
	switch values.CurrentOption {
	case 0:
		utils.WriteColor(utils.CenterText("> Chercher un adversaire  |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Accéder à la boutique   |", 75))
		utils.Writeln(utils.CenterText("| Consulter l'inventaire  |", 75))
		utils.Writeln(utils.CenterText("| Accéder à IKEA          |", 75))
		utils.Writeln(utils.CenterText("| Voir la niche           |", 75))
		utils.Writeln(utils.CenterText("| Qui sont-ils ?          |", 75))
		utils.Writeln(utils.CenterText("| Quitter                 |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		utils.Writeln(ascii.SwordAD)
		break
	case 1:
		utils.Writeln(utils.CenterText("| Chercher un adversaire  |", 75))
		utils.WriteColor(utils.CenterText("> Accéder à la boutique   |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Consulter l'inventaire  |", 75))
		utils.Writeln(utils.CenterText("| Accéder à IKEA          |", 75))
		utils.Writeln(utils.CenterText("| Voir la niche           |", 75))
		utils.Writeln(utils.CenterText("| Qui sont-ils ?          |", 75))
		utils.Writeln(utils.CenterText("| Quitter                 |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		utils.Writeln(ascii.StandBoutique)
		break
	case 2:
		utils.Writeln(utils.CenterText("| Chercher un adversaire  |", 75))
		utils.Writeln(utils.CenterText("| Accéder à la boutique   |", 75))
		utils.WriteColor(utils.CenterText("> Consulter l'inventaire  |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Accéder à IKEA          |", 75))
		utils.Writeln(utils.CenterText("| Voir la niche           |", 75))
		utils.Writeln(utils.CenterText("| Qui sont-ils ?          |", 75))
		utils.Writeln(utils.CenterText("| Quitter                 |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		utils.Writeln(ascii.InventaireBag)
		break
	case 3:
		utils.Writeln(utils.CenterText("| Chercher un adversaire  |", 75))
		utils.Writeln(utils.CenterText("| Accéder à la boutique   |", 75))
		utils.Writeln(utils.CenterText("| Consulter l'inventaire  |", 75))
		utils.WriteColor(utils.CenterText("> Accéder à IKEA          |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Voir la niche           |", 75))
		utils.Writeln(utils.CenterText("| Qui sont-ils ?          |", 75))
		utils.Writeln(utils.CenterText("| Quitter                 |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		utils.Writeln(ascii.Ikea)
		break
	case 4:
		utils.Writeln(utils.CenterText("| Chercher un adversaire  |", 75))
		utils.Writeln(utils.CenterText("| Accéder à la boutique   |", 75))
		utils.Writeln(utils.CenterText("| Consulter l'inventaire  |", 75))
		utils.Writeln(utils.CenterText("| Accéder à IKEA          |", 75))
		utils.WriteColor(utils.CenterText("> Voir la niche           |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Qui sont-ils ?          |", 75))
		utils.Writeln(utils.CenterText("| Quitter                 |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		utils.Writeln(ascii.DogHouse)
		break
	case 5:
		utils.Writeln(utils.CenterText("| Chercher un adversaire  |", 75))
		utils.Writeln(utils.CenterText("| Accéder à la boutique   |", 75))
		utils.Writeln(utils.CenterText("| Consulter l'inventaire  |", 75))
		utils.Writeln(utils.CenterText("| Accéder à IKEA          |", 75))
		utils.Writeln(utils.CenterText("| Voir la niche           |", 75))
		utils.WriteColor(utils.CenterText("> Qui sont-ils ?          |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Quitter                 |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		utils.Writeln(ascii.Mystery)
		break
	case 6:
		utils.Writeln(utils.CenterText("| Chercher un adversaire  |", 75))
		utils.Writeln(utils.CenterText("| Accéder à la boutique   |", 75))
		utils.Writeln(utils.CenterText("| Consulter l'inventaire  |", 75))
		utils.Writeln(utils.CenterText("| Accéder à IKEA          |", 75))
		utils.Writeln(utils.CenterText("| Voir la niche           |", 75))
		utils.Writeln(utils.CenterText("| Qui sont-ils ?          |", 75))
		utils.WriteColor(utils.CenterText("> Quitter                 |", 75), "yellow")
		utils.Writeln(utils.CenterText("===========================", 75))
		utils.Writeln(ascii.ExitDoor)
	}
}
