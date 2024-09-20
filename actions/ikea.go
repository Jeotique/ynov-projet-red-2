package actions

import (
	"projet/bdd"
	"projet/graphic"
	"projet/structs"
	"projet/utils"
	"projet/values"
)

func IkeaExec() {
	switch values.CurrentOption {
	case 0, 1, 2, 3, 4, 5:
		if values.Player.GetHouse() != values.CurrentOption {
			currentPrice := values.NichesPrice[values.CurrentOption]
			if structs.GetWood(&values.Player.Inventory) >= currentPrice.Wood && structs.GetViscous(&values.Player.Inventory) >= currentPrice.Viscous && structs.GetStone(&values.Player.Inventory) >= currentPrice.Stone {
				structs.SetWood(&values.Player.Inventory, structs.GetWood(&values.Player.Inventory)-currentPrice.Wood)
				structs.SetViscous(&values.Player.Inventory, structs.GetViscous(&values.Player.Inventory)-currentPrice.Viscous)
				structs.SetStone(&values.Player.Inventory, structs.GetStone(&values.Player.Inventory)-currentPrice.Stone)
				values.Player.SetHouse(values.CurrentOption)
				graphic.DisplayIkea()
				bdd.Database.SavePlayer(values.Player)
			} else {
				utils.WriteColor(utils.CenterText("Vous n'avez pas assez de ressource", 75), "red")
			}
		}
	case 6:
		values.CurrentOption = 0
		values.CurrentPage = "real_main_menu"
		values.CurrentOptionMax = 4
		graphic.RefreshRealMainMenu()
		break
	}
}
