package actions

import (
	"projet/bdd"
	"projet/graphic"
	"projet/utils"
	"projet/values"
)

func NicheExec() {
	switch values.CurrentOption {
	case 0:
		values.CurrentOption = 0
		values.CurrentPage = "niche_dog"
		values.CurrentOptionMax = 2
		graphic.DisplayDogInfo()
		break
	case 1:
		values.CurrentOption = 0
		values.CurrentPage = "niche_change"
		values.CurrentOptionMax = len(values.Player.Dogs)
		graphic.DiplayChangeDog()
		break
	case 2:
		values.CurrentOption = 0
		values.CurrentPage = "real_main_menu"
		values.CurrentOptionMax = 4
		graphic.RefreshRealMainMenu()
		break
	}
}

func NicheDogExec() {
	switch values.CurrentOption {
	case 0:
		if values.Player.Inventory.Heal > 0 {
			if values.Player.Dogs[values.Player.CurrentDog].Health < values.Player.Dogs[values.Player.CurrentDog].MaxHealth {
				values.Player.Inventory.Heal--
				values.Player.Dogs[values.Player.CurrentDog].Heal(5)
				graphic.DisplayDogInfo()
				bdd.Database.SavePlayer(values.Player)
			}
		}
		break
	case 1:
		if values.Player.Dogs[values.Player.CurrentDog].IsAlive() {
			values.Training = true
			values.Ennemy = utils.GetRandomDogClone(values.DogsList)
			values.CurrentOption = 0
			values.CurrentPage = "combat"
			values.CurrentOptionMax = -1
			graphic.DisplayCombat()
		} else {
			utils.WriteColorLn(utils.CenterText("Votre chien est mort", 75), "red")
		}
		break
	case 2:
		values.CurrentOption = 0
		values.CurrentPage = "niche_comp"
		values.CurrentOptionMax = 4
		graphic.DisplayCompetences()
		break
	case 3:
		values.CurrentOption = 0
		values.CurrentPage = "niche"
		values.CurrentOptionMax = 2
		graphic.DisplayNiche()
		break
	}
}

func NicheChangeExec() {
	if values.CurrentOption < len(values.Player.Dogs) {
		values.Player.CurrentDog = values.CurrentOption
		bdd.Database.SavePlayer(values.Player)
		graphic.DiplayChangeDog()
	} else {
		if values.CurrentOption == len(values.Player.Dogs)+1 {
			values.CurrentOption = 0
			values.CurrentPage = "niche"
			values.CurrentOptionMax = 2
			graphic.DisplayNiche()
		} else {
			values.CurrentOption = 0
			values.CurrentPage = "niche_buy"
			values.CurrentOptionMax = 7
			graphic.DisplayBuyDog()
		}
	}
}

func NicheBuyExec() {
	if values.CurrentOption < len(values.DogsRace) {
		if values.Player.Inventory.Money >= values.DogsPrice[values.DogsRace[values.CurrentOption]] {
			values.Player.Inventory.Money -= values.DogsPrice[values.DogsRace[values.CurrentOption]]
			values.Player.Dogs = append(values.Player.Dogs, utils.GetClonedDog(values.DogsList, values.DogsRace[values.CurrentOption]))
			bdd.Database.SavePlayer(values.Player)
			values.CurrentOption = 0
			values.CurrentPage = "niche_change"
			values.CurrentOptionMax = len(values.Player.Dogs)
			graphic.DiplayChangeDog()
		} else {
			utils.WriteColor(utils.CenterText("Vous n'avez pas assez d'argent", 75), "red")
		}
	} else {
		values.CurrentOption = 0
		values.CurrentPage = "niche_change"
		values.CurrentOptionMax = len(values.Player.Dogs)
		graphic.DiplayChangeDog()
	}
}

func NicheCompExec() {
	switch values.CurrentOption {
	case 0:
		break
	case 1:
		break
	case 2:
		break
	case 3:
		break
	case 4:
		values.CurrentOption = 0
		values.CurrentPage = "niche_dog"
		values.CurrentOptionMax = 2
		graphic.DisplayDogInfo()
		bdd.Database.SavePlayer(values.Player)
		break
	}
}
