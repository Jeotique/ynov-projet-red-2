package graphic

import (
	"fmt"
	"projet/ascii"
	"projet/structs"
	"projet/utils"
	"projet/values"
	"strconv"
)

func DisplayIkea() {
	utils.ClearTerminal()

	utils.Writeln("")
	utils.Writeln("")
	utils.Writeln(ascii.MenuDogsTitle)

	utils.Writeln(utils.CenterText("======== Menu IKEA ========", 75))
	CurrentPrice := values.NichesPrice[values.CurrentOption]
	switch values.CurrentOption {
	case 0, 1, 2, 3, 4, 5:
		if values.Player.GetHouse() == values.CurrentOption {
			DisplayNiches(values.CurrentOption, "blue", "default")
			utils.WriteColor(utils.CenterText("Cette niche est déjà installée", 75), "green")
		} else if structs.GetWood(&values.Player.Inventory) >= CurrentPrice.Wood && structs.GetViscous(&values.Player.Inventory) >= CurrentPrice.Viscous && structs.GetStone(&values.Player.Inventory) >= CurrentPrice.Stone {
			DisplayNiches(values.CurrentOption, "green", "default")
			DisplayResources(CurrentPrice, &values.Player.Inventory)
		} else {
			DisplayNiches(values.CurrentOption, "red", "default")
			DisplayResources(CurrentPrice, &values.Player.Inventory)
		}
	case 6:
		DisplayNiches(-1, "yellow", "default")
		/*case 0:
			if values.Player.GetHouse() == values.CurrentOption {
				utils.WriteColor(utils.CenterText("> Niche tente             |", 75), "blue")
				utils.Writeln(utils.CenterText("| Niche classique         |", 75))
				utils.Writeln(utils.CenterText("| Niche Confortable       |", 75))
				utils.Writeln(utils.CenterText("| Niche Moderne           |", 75))
				utils.Writeln(utils.CenterText("| Niche avec terrasse     |", 75))
				utils.Writeln(utils.CenterText("| Niche de luxe           |", 75))
				utils.Writeln(utils.CenterText("| Retour                  |", 75))
				utils.Writeln(utils.CenterText("===========================", 75))
				utils.WriteColor(utils.CenterText("Cette niche est déjà installée", 75), "green")
			} else if structs.GetWood(&values.Player.Inventory) >= CurrentPrice.Wood && structs.GetViscous(&values.Player.Inventory) >= CurrentPrice.Viscous && structs.GetStone(&values.Player.Inventory) >= CurrentPrice.Stone {
				utils.WriteColor(utils.CenterText("> Niche tente             |", 75), "green")
				utils.Writeln(utils.CenterText("| Niche classique         |", 75))
				utils.Writeln(utils.CenterText("| Niche Confortable       |", 75))
				utils.Writeln(utils.CenterText("| Niche Moderne           |", 75))
				utils.Writeln(utils.CenterText("| Niche avec terrasse     |", 75))
				utils.Writeln(utils.CenterText("| Niche de luxe           |", 75))
				utils.Writeln(utils.CenterText("| Retour                  |", 75))
				utils.Writeln(utils.CenterText("===========================", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(CurrentPrice.Wood)+" | Viscous: "+strconv.Itoa(CurrentPrice.Viscous)+" | Pierre: "+strconv.Itoa(CurrentPrice.Stone), 75))
				utils.Writeln(utils.CenterText("Vous avez :             ", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(structs.GetWood(&values.Player.Inventory))+" | Viscous: "+strconv.Itoa(structs.GetViscous(&values.Player.Inventory))+" | Pierre: "+strconv.Itoa(structs.GetStone(&values.Player.Inventory)), 75))
			} else {
				utils.WriteColor(utils.CenterText("> Niche tente             |", 75), "red")
				utils.Writeln(utils.CenterText("| Niche classique         |", 75))
				utils.Writeln(utils.CenterText("| Niche Confortable       |", 75))
				utils.Writeln(utils.CenterText("| Niche Moderne           |", 75))
				utils.Writeln(utils.CenterText("| Niche avec terrasse     |", 75))
				utils.Writeln(utils.CenterText("| Niche de luxe           |", 75))
				utils.Writeln(utils.CenterText("| Retour                  |", 75))
				utils.Writeln(utils.CenterText("===========================", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(CurrentPrice.Wood)+" | Viscous: "+strconv.Itoa(CurrentPrice.Viscous)+" | Pierre: "+strconv.Itoa(CurrentPrice.Stone), 75))
				utils.Writeln(utils.CenterText("Vous avez :             ", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(structs.GetWood(&values.Player.Inventory))+" | Viscous: "+strconv.Itoa(structs.GetViscous(&values.Player.Inventory))+" | Pierre: "+strconv.Itoa(structs.GetStone(&values.Player.Inventory)), 75))
			}
			break
		case 1:
			if values.Player.GetHouse() == values.CurrentOption {
				utils.Writeln(utils.CenterText("| Niche tente             |", 75))
				utils.WriteColor(utils.CenterText("> Niche classique         |", 75), "blue")
				utils.Writeln(utils.CenterText("| Niche Confortable       |", 75))
				utils.Writeln(utils.CenterText("| Niche Moderne           |", 75))
				utils.Writeln(utils.CenterText("| Niche avec terrasse     |", 75))
				utils.Writeln(utils.CenterText("| Niche de luxe           |", 75))
				utils.Writeln(utils.CenterText("| Retour                  |", 75))
				utils.Writeln(utils.CenterText("===========================", 75))
				utils.WriteColor(utils.CenterText("Cette niche est déjà installée", 75), "green")
			} else if structs.GetWood(&values.Player.Inventory) >= CurrentPrice.Wood && structs.GetViscous(&values.Player.Inventory) >= CurrentPrice.Viscous && structs.GetStone(&values.Player.Inventory) >= CurrentPrice.Stone {
				utils.Writeln(utils.CenterText("| Niche tente             |", 75))
				utils.WriteColor(utils.CenterText("> Niche classique         |", 75), "green")
				utils.Writeln(utils.CenterText("| Niche Confortable       |", 75))
				utils.Writeln(utils.CenterText("| Niche Moderne           |", 75))
				utils.Writeln(utils.CenterText("| Niche avec terrasse     |", 75))
				utils.Writeln(utils.CenterText("| Niche de luxe           |", 75))
				utils.Writeln(utils.CenterText("| Retour                  |", 75))
				utils.Writeln(utils.CenterText("===========================", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(CurrentPrice.Wood)+" | Viscous: "+strconv.Itoa(CurrentPrice.Viscous)+" | Pierre: "+strconv.Itoa(CurrentPrice.Stone), 75))
				utils.Writeln(utils.CenterText("Vous avez :             ", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(structs.GetWood(&values.Player.Inventory))+" | Viscous: "+strconv.Itoa(structs.GetViscous(&values.Player.Inventory))+" | Pierre: "+strconv.Itoa(structs.GetStone(&values.Player.Inventory)), 75))
			} else {
				utils.Writeln(utils.CenterText("| Niche tente             |", 75))
				utils.WriteColor(utils.CenterText("> Niche classique         |", 75), "red")
				utils.Writeln(utils.CenterText("| Niche Confortable       |", 75))
				utils.Writeln(utils.CenterText("| Niche Moderne           |", 75))
				utils.Writeln(utils.CenterText("| Niche avec terrasse     |", 75))
				utils.Writeln(utils.CenterText("| Niche de luxe           |", 75))
				utils.Writeln(utils.CenterText("| Retour                  |", 75))
				utils.Writeln(utils.CenterText("===========================", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(CurrentPrice.Wood)+" | Viscous: "+strconv.Itoa(CurrentPrice.Viscous)+" | Pierre: "+strconv.Itoa(CurrentPrice.Stone), 75))
				utils.Writeln(utils.CenterText("Vous avez :             ", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(structs.GetWood(&values.Player.Inventory))+" | Viscous: "+strconv.Itoa(structs.GetViscous(&values.Player.Inventory))+" | Pierre: "+strconv.Itoa(structs.GetStone(&values.Player.Inventory)), 75))
			}
			break
		case 2:
			if values.Player.GetHouse() == values.CurrentOption {
				utils.Writeln(utils.CenterText("| Niche tente             |", 75))
				utils.Writeln(utils.CenterText("| Niche classique         |", 75))
				utils.WriteColor(utils.CenterText("> Niche Confortable       |", 75), "blue")
				utils.Writeln(utils.CenterText("| Niche Moderne           |", 75))
				utils.Writeln(utils.CenterText("| Niche avec terrasse     |", 75))
				utils.Writeln(utils.CenterText("| Niche de luxe           |", 75))
				utils.Writeln(utils.CenterText("| Retour                  |", 75))
				utils.Writeln(utils.CenterText("===========================", 75))
				utils.WriteColor(utils.CenterText("Cette niche est déjà installée", 75), "green")
			} else if structs.GetWood(&values.Player.Inventory) >= CurrentPrice.Wood && structs.GetViscous(&values.Player.Inventory) >= CurrentPrice.Viscous && structs.GetStone(&values.Player.Inventory) >= CurrentPrice.Stone {
				utils.Writeln(utils.CenterText("| Niche tente             |", 75))
				utils.Writeln(utils.CenterText("| Niche classique         |", 75))
				utils.WriteColor(utils.CenterText("> Niche Confortable       |", 75), "green")
				utils.Writeln(utils.CenterText("| Niche Moderne           |", 75))
				utils.Writeln(utils.CenterText("| Niche avec terrasse     |", 75))
				utils.Writeln(utils.CenterText("| Niche de luxe           |", 75))
				utils.Writeln(utils.CenterText("| Retour                  |", 75))
				utils.Writeln(utils.CenterText("===========================", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(CurrentPrice.Wood)+" | Viscous: "+strconv.Itoa(CurrentPrice.Viscous)+" | Pierre: "+strconv.Itoa(CurrentPrice.Stone), 75))
				utils.Writeln(utils.CenterText("Vous avez :             ", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(structs.GetWood(&values.Player.Inventory))+" | Viscous: "+strconv.Itoa(structs.GetViscous(&values.Player.Inventory))+" | Pierre: "+strconv.Itoa(structs.GetStone(&values.Player.Inventory)), 75))
			} else {
				utils.Writeln(utils.CenterText("| Niche tente			 |", 75))
				utils.Writeln(utils.CenterText("| Niche classique		 |", 75))
				utils.WriteColor(utils.CenterText("> Niche Confortable		 |", 75), "red")
				utils.Writeln(utils.CenterText("| Niche Moderne			 |", 75))
				utils.Writeln(utils.CenterText("| Niche avec terrasse		 |", 75))
				utils.Writeln(utils.CenterText("| Niche de luxe			 |", 75))
				utils.Writeln(utils.CenterText("| Retour				 |", 75))
				utils.Writeln(utils.CenterText("===========================", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(CurrentPrice.Wood)+" | Viscous: "+strconv.Itoa(CurrentPrice.Viscous)+" | Pierre: "+strconv.Itoa(CurrentPrice.Stone), 75))
				utils.Writeln(utils.CenterText("Vous avez :			 ", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(structs.GetWood(&values.Player.Inventory))+" | Viscous: "+strconv.Itoa(structs.GetViscous(&values.Player.Inventory))+" | Pierre: "+strconv.Itoa(structs.GetStone(&values.Player.Inventory)), 75))
			}
			break
		case 3:
			if values.Player.GetHouse() == values.CurrentOption {
				utils.Writeln(utils.CenterText("| Niche tente             |", 75))
				utils.Writeln(utils.CenterText("| Niche classique         |", 75))
				utils.Writeln(utils.CenterText("| Niche Confortable       |", 75))
				utils.WriteColor(utils.CenterText("> Niche Moderne           |", 75), "blue")
				utils.Writeln(utils.CenterText("| Niche avec terrasse     |", 75))
				utils.Writeln(utils.CenterText("| Niche de luxe           |", 75))
				utils.Writeln(utils.CenterText("| Retour                  |", 75))
				utils.Writeln(utils.CenterText("===========================", 75))
				utils.WriteColor(utils.CenterText("Cette niche est déjà installée", 75), "green")
			} else if structs.GetWood(&values.Player.Inventory) >= CurrentPrice.Wood && structs.GetViscous(&values.Player.Inventory) >= CurrentPrice.Viscous && structs.GetStone(&values.Player.Inventory) >= CurrentPrice.Stone {
				utils.Writeln(utils.CenterText("| Niche tente             |", 75))
				utils.Writeln(utils.CenterText("| Niche classique         |", 75))
				utils.Writeln(utils.CenterText("| Niche Confortable       |", 75))
				utils.WriteColor(utils.CenterText("> Niche Moderne           |", 75), "green")
				utils.Writeln(utils.CenterText("| Niche avec terrasse     |", 75))
				utils.Writeln(utils.CenterText("| Niche de luxe           |", 75))
				utils.Writeln(utils.CenterText("| Retour                  |", 75))
				utils.Writeln(utils.CenterText("===========================", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(CurrentPrice.Wood)+" | Viscous: "+strconv.Itoa(CurrentPrice.Viscous)+" | Pierre: "+strconv.Itoa(CurrentPrice.Stone), 75))
				utils.Writeln(utils.CenterText("Vous avez :             ", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(structs.GetWood(&values.Player.Inventory))+" | Viscous: "+strconv.Itoa(structs.GetViscous(&values.Player.Inventory))+" | Pierre: "+strconv.Itoa(structs.GetStone(&values.Player.Inventory)), 75))
			} else {
				utils.Writeln(utils.CenterText("| Niche tente			 |", 75))
				utils.Writeln(utils.CenterText("| Niche classique		 |", 75))
				utils.Writeln(utils.CenterText("| Niche Confortable		 |", 75))
				utils.WriteColor(utils.CenterText("> Niche Moderne			 |", 75), "red")
				utils.Writeln(utils.CenterText("| Niche avec terrasse		 |", 75))
				utils.Writeln(utils.CenterText("| Niche de luxe			 |", 75))
				utils.Writeln(utils.CenterText("| Retour				 |", 75))
				utils.Writeln(utils.CenterText("===========================", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(CurrentPrice.Wood)+" | Viscous: "+strconv.Itoa(CurrentPrice.Viscous)+" | Pierre: "+strconv.Itoa(CurrentPrice.Stone), 75))
				utils.Writeln(utils.CenterText("Vous avez :			 ", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(structs.GetWood(&values.Player.Inventory))+" | Viscous: "+strconv.Itoa(structs.GetViscous(&values.Player.Inventory))+" | Pierre: "+strconv.Itoa(structs.GetStone(&values.Player.Inventory)), 75))
			}
			break
		case 4:
			if values.Player.GetHouse() == values.CurrentOption {
				utils.Writeln(utils.CenterText("| Niche tente             |", 75))
				utils.Writeln(utils.CenterText("| Niche classique         |", 75))
				utils.Writeln(utils.CenterText("| Niche Confortable       |", 75))
				utils.Writeln(utils.CenterText("| Niche Moderne           |", 75))
				utils.WriteColor(utils.CenterText("> Niche avec terrasse     |", 75), "blue")
				utils.Writeln(utils.CenterText("| Niche de luxe           |", 75))
				utils.Writeln(utils.CenterText("| Retour                  |", 75))
				utils.Writeln(utils.CenterText("===========================", 75))
				utils.WriteColor(utils.CenterText("Cette niche est déjà installée", 75), "green")
			} else if structs.GetWood(&values.Player.Inventory) >= CurrentPrice.Wood && structs.GetViscous(&values.Player.Inventory) >= CurrentPrice.Viscous && structs.GetStone(&values.Player.Inventory) >= CurrentPrice.Stone {
				utils.Writeln(utils.CenterText("| Niche tente             |", 75))
				utils.Writeln(utils.CenterText("| Niche classique         |", 75))
				utils.Writeln(utils.CenterText("| Niche Confortable       |", 75))
				utils.Writeln(utils.CenterText("| Niche Moderne           |", 75))
				utils.WriteColor(utils.CenterText("> Niche avec terrasse     |", 75), "green")
				utils.Writeln(utils.CenterText("| Niche de luxe           |", 75))
				utils.Writeln(utils.CenterText("| Retour                  |", 75))
				utils.Writeln(utils.CenterText("===========================", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(CurrentPrice.Wood)+" | Viscous: "+strconv.Itoa(CurrentPrice.Viscous)+" | Pierre: "+strconv.Itoa(CurrentPrice.Stone), 75))
				utils.Writeln(utils.CenterText("Vous avez :             ", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(structs.GetWood(&values.Player.Inventory))+" | Viscous: "+strconv.Itoa(structs.GetViscous(&values.Player.Inventory))+" | Pierre: "+strconv.Itoa(structs.GetStone(&values.Player.Inventory)), 75))
			} else {
				utils.Writeln(utils.CenterText("| Niche tente			 |", 75))
				utils.Writeln(utils.CenterText("| Niche classique		 |", 75))
				utils.Writeln(utils.CenterText("| Niche Confortable		 |", 75))
				utils.Writeln(utils.CenterText("| Niche Moderne			 |", 75))
				utils.WriteColor(utils.CenterText("> Niche avec terrasse		 |", 75), "red")
				utils.Writeln(utils.CenterText("| Niche de luxe			 |", 75))
				utils.Writeln(utils.CenterText("| Retour				 |", 75))
				utils.Writeln(utils.CenterText("===========================", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(CurrentPrice.Wood)+" | Viscous: "+strconv.Itoa(CurrentPrice.Viscous)+" | Pierre: "+strconv.Itoa(CurrentPrice.Stone), 75))
				utils.Writeln(utils.CenterText("Vous avez :			 ", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(structs.GetWood(&values.Player.Inventory))+" | Viscous: "+strconv.Itoa(structs.GetViscous(&values.Player.Inventory))+" | Pierre: "+strconv.Itoa(structs.GetStone(&values.Player.Inventory)), 75))
			}
			break
		case 5:
			if values.Player.GetHouse() == values.CurrentOption {
				utils.Writeln(utils.CenterText("| Niche tente             |", 75))
				utils.Writeln(utils.CenterText("| Niche classique         |", 75))
				utils.Writeln(utils.CenterText("| Niche Confortable       |", 75))
				utils.Writeln(utils.CenterText("| Niche Moderne           |", 75))
				utils.Writeln(utils.CenterText("| Niche avec terrasse     |", 75))
				utils.WriteColor(utils.CenterText("> Niche de luxe           |", 75), "blue")
				utils.Writeln(utils.CenterText("| Retour                  |", 75))
				utils.Writeln(utils.CenterText("===========================", 75))
				utils.WriteColor(utils.CenterText("Cette niche est déjà installée", 75), "green")
			} else if structs.GetWood(&values.Player.Inventory) >= CurrentPrice.Wood && structs.GetViscous(&values.Player.Inventory) >= CurrentPrice.Viscous && structs.GetStone(&values.Player.Inventory) >= CurrentPrice.Stone {
				utils.Writeln(utils.CenterText("| Niche tente             |", 75))
				utils.Writeln(utils.CenterText("| Niche classique         |", 75))
				utils.Writeln(utils.CenterText("| Niche Confortable       |", 75))
				utils.Writeln(utils.CenterText("| Niche Moderne           |", 75))
				utils.Writeln(utils.CenterText("| Niche avec terrasse     |", 75))
				utils.WriteColor(utils.CenterText("> Niche de luxe           |", 75), "green")
				utils.Writeln(utils.CenterText("| Retour                  |", 75))
				utils.Writeln(utils.CenterText("===========================", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(CurrentPrice.Wood)+" | Viscous: "+strconv.Itoa(CurrentPrice.Viscous)+" | Pierre: "+strconv.Itoa(CurrentPrice.Stone), 75))
				utils.Writeln(utils.CenterText("Vous avez :             ", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(structs.GetWood(&values.Player.Inventory))+" | Viscous: "+strconv.Itoa(structs.GetViscous(&values.Player.Inventory))+" | Pierre: "+strconv.Itoa(structs.GetStone(&values.Player.Inventory)), 75))
			} else {
				utils.Writeln(utils.CenterText("| Niche tente			 |", 75))
				utils.Writeln(utils.CenterText("| Niche classique		 |", 75))
				utils.Writeln(utils.CenterText("| Niche Confortable		 |", 75))
				utils.Writeln(utils.CenterText("| Niche Moderne			 |", 75))
				utils.Writeln(utils.CenterText("| Niche avec terrasse		 |", 75))
				utils.WriteColor(utils.CenterText("> Niche de luxe			 |", 75), "red")
				utils.Writeln(utils.CenterText("| Retour				 |", 75))
				utils.Writeln(utils.CenterText("===========================", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(CurrentPrice.Wood)+" | Viscous: "+strconv.Itoa(CurrentPrice.Viscous)+" | Pierre: "+strconv.Itoa(CurrentPrice.Stone), 75))
				utils.Writeln(utils.CenterText("Vous avez :			 ", 75))
				utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(structs.GetWood(&values.Player.Inventory))+" | Viscous: "+strconv.Itoa(structs.GetViscous(&values.Player.Inventory))+" | Pierre: "+strconv.Itoa(structs.GetStone(&values.Player.Inventory)), 75))
			}
			break
		case 6:
			utils.Writeln(utils.CenterText("| Niche tente             |", 75))
			utils.Writeln(utils.CenterText("| Niche classique         |", 75))
			utils.Writeln(utils.CenterText("| Niche Confortable       |", 75))
			utils.Writeln(utils.CenterText("| Niche Moderne           |", 75))
			utils.Writeln(utils.CenterText("| Niche avec terrasse     |", 75))
			utils.Writeln(utils.CenterText("| Niche de luxe           |", 75))
			utils.WriteColor(utils.CenterText("> Retour                  |", 75), "yellow")
			utils.Writeln(utils.CenterText("===========================", 75))
			break
		*/
	}
	utils.Writeln(ascii.Niches[values.CurrentOption])
}

func DisplayNiches(selectedIndex int, selectedColor, defaultColor string) {
	niches := []string{
		"Niche tente",
		"Niche classique",
		"Niche Confortable",
		"Niche Moderne",
		"Niche avec terrasse",
		"Niche de luxe",
	}

	for i, niche := range niches {
		color := defaultColor
		left := "|"
		if i == selectedIndex {
			color = selectedColor
			left = ">"
		}
		utils.WriteColor(utils.CenterText(fmt.Sprintf(left+" %-23s |", niche), 75), color)
	}
	if selectedIndex == -1 {
		utils.WriteColor(utils.CenterText("> Retour                  |", 75), selectedColor)
	} else {
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
	}
	utils.Writeln(utils.CenterText("===========================", 75))
}

// Fonction pour afficher les ressources
func DisplayResources(CurrentPrice values.NichePrice, inventory *structs.Inventory) {
	utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(CurrentPrice.Wood)+" | Viscous: "+strconv.Itoa(CurrentPrice.Viscous)+" | Pierre: "+strconv.Itoa(CurrentPrice.Stone), 75))
	utils.Writeln(utils.CenterText("Vous avez :             ", 75))
	utils.Writeln(utils.CenterText("Bois: "+strconv.Itoa(structs.GetWood(inventory))+" | Viscous: "+strconv.Itoa(structs.GetViscous(inventory))+" | Pierre: "+strconv.Itoa(structs.GetStone(inventory)), 75))
}
