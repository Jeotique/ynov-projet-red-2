package actions

import (
	"projet/bdd"
	"projet/graphic"
	"projet/utils"
	"projet/values"
)

func MarketExec() {
	switch values.CurrentOption {
	case 0:
		values.CurrentOption = 0
		values.CurrentPage = "market_buy"
		values.CurrentOptionMax = 7
		graphic.DisplayBuyMarket()
		break
	case 1:
		values.CurrentOption = 0
		values.CurrentPage = "market_sell"
		values.CurrentOptionMax = 7
		graphic.DisplaySellMarket()
		break
	case 2:
		values.CurrentOption = 0
		values.CurrentPage = "real_main_menu"
		values.CurrentOptionMax = 4
		graphic.RefreshRealMainMenu()
		break
	}
}

func MarketBuyExec() {
	switch values.CurrentOption {
	case 0:

	case 1:

	case 2:

	case 3:

	case 4:

	case 5:

	case 6:

	case 7:

	case 8:
		if BuyFromMarket() {
			graphic.DisplayBuyMarket()
		} else {
			utils.WriteColorLn(utils.CenterText("Vous n'avez pas assez d'argent", 75), "red")
		}
	case 9:
		values.CurrentOption = 0
		values.CurrentPage = "market"
		values.CurrentOptionMax = 2
		graphic.DisplayMarket()
		break
	}
}

func calcMarketBuyPrice() int {
	price := 0
	for name, value := range values.MarketCount {
		price += value * values.MarketPrice[name]
	}
	return price
}

func calcMarketSellPrice() int {
	price := 0
	for name, value := range values.MarketSell {
		price += value * (values.MarketPrice[name] / 2)
	}
	return price
}

func BuyFromMarket() bool {
	if values.Player.Inventory.Money >= calcMarketBuyPrice() {
		values.Player.Inventory.Money -= calcMarketBuyPrice()
		values.Player.Inventory.Heal += values.MarketCount["Heal"]
		values.Player.Inventory.Poison += values.MarketCount["Poison"]
		values.Player.Inventory.Kalashnikov = values.MarketCount["Kalashnikov"] > 0
		values.Player.Inventory.Fireball += values.MarketCount["Fireball"]
		values.Player.Inventory.Iceball += values.MarketCount["Iceball"]
		values.Player.Inventory.Bomb += values.MarketCount["Bomb"]
		values.Player.Inventory.Shield += values.MarketCount["Shield"]
		values.Player.Inventory.Heineken += values.MarketCount["Heineken"]
		for name, _ := range values.MarketCount {
			values.MarketCount[name] = 0
		}
		bdd.Database.SavePlayer(values.Player)
		return true
	}
	return false
}

func SellFromMarket() bool {
	if calcMarketSellPrice() > 0 {
		values.Player.Inventory.Money += calcMarketSellPrice()
		values.Player.Inventory.Heal -= values.MarketSell["Heal"]
		values.Player.Inventory.Poison -= values.MarketSell["Poison"]
		values.Player.Inventory.Kalashnikov = values.MarketSell["Kalashnikov"] > 0
		values.Player.Inventory.Fireball -= values.MarketSell["Fireball"]
		values.Player.Inventory.Iceball -= values.MarketSell["Iceball"]
		values.Player.Inventory.Bomb -= values.MarketSell["Bomb"]
		values.Player.Inventory.Shield -= values.MarketSell["Shield"]
		values.Player.Inventory.Heineken -= values.MarketSell["Heineken"]
		for name, _ := range values.MarketSell {
			values.MarketSell[name] = 0
		}
		bdd.Database.SavePlayer(values.Player)
		return true
	}
	return false
}

func MarketSellExec() {
	switch values.CurrentOption {
	case 0:

	case 1:

	case 2:

	case 3:

	case 4:

	case 5:

	case 6:

	case 7:

	case 8:
		if SellFromMarket() {
			graphic.DisplaySellMarket()
		} else {
			utils.WriteColorLn(utils.CenterText("Vous n'avez rien sélectionné", 75), "red")
		}
	case 9:
		values.CurrentOption = 0
		values.CurrentPage = "market"
		values.CurrentOptionMax = 2
		graphic.DisplayMarket()
		break
	}
}
