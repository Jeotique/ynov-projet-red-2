package graphic

import (
	"projet/ascii"
	"projet/utils"
	"projet/values"
	"strconv"
)

func DisplayMarket() {
	utils.ClearTerminal()

	utils.Writeln("")
	utils.Writeln("")
	utils.Writeln(ascii.MenuDogsTitle)

	utils.Writeln(utils.CenterText("====== Menu Marchand ======", 75))
	switch values.CurrentOption {
	case 0:
		utils.WriteColor(utils.CenterText("> Acheter                 |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Vendre                  |", 75))
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 1:
		utils.Writeln(utils.CenterText("| Acheter                 |", 75))
		utils.WriteColor(utils.CenterText("> Vendre                  |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 2:
		utils.Writeln(utils.CenterText("| Acheter                 |", 75))
		utils.Writeln(utils.CenterText("| Vendre                  |", 75))
		utils.WriteColor(utils.CenterText("> Retour                  |", 75), "yellow")
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	}
}

func DisplayBuyMarket() {
	utils.ClearTerminal()

	utils.Writeln("")
	utils.Writeln("")
	utils.Writeln(ascii.MenuDogsTitle)

	utils.Writeln(utils.CenterText("========= Acheter =========", 75))
	switch values.CurrentOption {
	case 0:
		utils.WriteColor(utils.CenterText("> Heal              <["+strconv.Itoa(values.MarketCount["Heal"])+"]> |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Poison            <%d> |", 75), values.MarketCount["Poison"])
		utils.Writeln(utils.CenterText("| Kalashnikov       <%d> |", 75), values.MarketCount["Kalashnikov"])
		utils.Writeln(utils.CenterText("| Fireball          <%d> |", 75), values.MarketCount["Fireball"])
		utils.Writeln(utils.CenterText("| Iceball           <%d> |", 75), values.MarketCount["Iceball"])
		utils.Writeln(utils.CenterText("| Bomb              <%d> |", 75), values.MarketCount["Bomb"])
		utils.Writeln(utils.CenterText("| Shield            <%d> |", 75), values.MarketCount["Shield"])
		utils.Writeln(utils.CenterText("| Heineken          <%d> |", 75), values.MarketCount["Heineken"])
		utils.Writeln(utils.CenterText("| Confirmer les achats    |", 75))
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 1:
		utils.Writeln(utils.CenterText("| Heal              <%d> |", 75), values.MarketCount["Heal"])
		utils.WriteColor(utils.CenterText("> Poison            <["+strconv.Itoa(values.MarketCount["Poison"])+"]> |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Kalashnikov       <%d> |", 75), values.MarketCount["Kalashnikov"])
		utils.Writeln(utils.CenterText("| Fireball          <%d> |", 75), values.MarketCount["Fireball"])
		utils.Writeln(utils.CenterText("| Iceball           <%d> |", 75), values.MarketCount["Iceball"])
		utils.Writeln(utils.CenterText("| Bomb              <%d> |", 75), values.MarketCount["Bomb"])
		utils.Writeln(utils.CenterText("| Shield            <%d> |", 75), values.MarketCount["Shield"])
		utils.Writeln(utils.CenterText("| Heineken          <%d> |", 75), values.MarketCount["Heineken"])
		utils.Writeln(utils.CenterText("| Confirmer les achats    |", 75))
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 2:
		utils.Writeln(utils.CenterText("| Heal              <%d> |", 75), values.MarketCount["Heal"])
		utils.Writeln(utils.CenterText("| Poison            <%d> |", 75), values.MarketCount["Poison"])
		utils.WriteColor(utils.CenterText("> Kalashnikov       <["+strconv.Itoa(values.MarketCount["Kalashnikov"])+"]> |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Fireball          <%d> |", 75), values.MarketCount["Fireball"])
		utils.Writeln(utils.CenterText("| Iceball           <%d> |", 75), values.MarketCount["Iceball"])
		utils.Writeln(utils.CenterText("| Bomb              <%d> |", 75), values.MarketCount["Bomb"])
		utils.Writeln(utils.CenterText("| Shield            <%d> |", 75), values.MarketCount["Shield"])
		utils.Writeln(utils.CenterText("| Heineken          <%d> |", 75), values.MarketCount["Heineken"])
		utils.Writeln(utils.CenterText("| Confirmer les achats    |", 75))
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 3:
		utils.Writeln(utils.CenterText("| Heal              <%d> |", 75), values.MarketCount["Heal"])
		utils.Writeln(utils.CenterText("| Poison            <%d> |", 75), values.MarketCount["Poison"])
		utils.Writeln(utils.CenterText("| Kalashnikov       <%d> |", 75), values.MarketCount["Kalashnikov"])
		utils.WriteColor(utils.CenterText("> Fireball          <["+strconv.Itoa(values.MarketCount["Fireball"])+"]> |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Iceball           <%d> |", 75), values.MarketCount["Iceball"])
		utils.Writeln(utils.CenterText("| Bomb              <%d> |", 75), values.MarketCount["Bomb"])
		utils.Writeln(utils.CenterText("| Shield            <%d> |", 75), values.MarketCount["Shield"])
		utils.Writeln(utils.CenterText("| Heineken          <%d> |", 75), values.MarketCount["Heineken"])
		utils.Writeln(utils.CenterText("| Confirmer les achats    |", 75))
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 4:
		utils.Writeln(utils.CenterText("| Heal              <%d> |", 75), values.MarketCount["Heal"])
		utils.Writeln(utils.CenterText("| Poison            <%d> |", 75), values.MarketCount["Poison"])
		utils.Writeln(utils.CenterText("| Kalashnikov       <%d> |", 75), values.MarketCount["Kalashnikov"])
		utils.Writeln(utils.CenterText("| Fireball          <%d> |", 75), values.MarketCount["Fireball"])
		utils.WriteColor(utils.CenterText("> Iceball           <["+strconv.Itoa(values.MarketCount["Iceball"])+"]> |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Bomb              <%d> |", 75), values.MarketCount["Bomb"])
		utils.Writeln(utils.CenterText("| Shield            <%d> |", 75), values.MarketCount["Shield"])
		utils.Writeln(utils.CenterText("| Heineken          <%d> |", 75), values.MarketCount["Heineken"])
		utils.Writeln(utils.CenterText("| Confirmer les achats    |", 75))
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 5:
		utils.Writeln(utils.CenterText("| Heal              <%d> |", 75), values.MarketCount["Heal"])
		utils.Writeln(utils.CenterText("| Poison            <%d> |", 75), values.MarketCount["Poison"])
		utils.Writeln(utils.CenterText("| Kalashnikov       <%d> |", 75), values.MarketCount["Kalashnikov"])
		utils.Writeln(utils.CenterText("| Fireball          <%d> |", 75), values.MarketCount["Fireball"])
		utils.Writeln(utils.CenterText("| Iceball           <%d> |", 75), values.MarketCount["Iceball"])
		utils.WriteColor(utils.CenterText("> Bomb              <["+strconv.Itoa(values.MarketCount["Bomb"])+"]> |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Shield            <%d> |", 75), values.MarketCount["Shield"])
		utils.Writeln(utils.CenterText("| Heineken          <%d> |", 75), values.MarketCount["Heineken"])
		utils.Writeln(utils.CenterText("| Confirmer les achats    |", 75))
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 6:
		utils.Writeln(utils.CenterText("| Heal              <%d> |", 75), values.MarketCount["Heal"])
		utils.Writeln(utils.CenterText("| Poison            <%d> |", 75), values.MarketCount["Poison"])
		utils.Writeln(utils.CenterText("| Kalashnikov       <%d> |", 75), values.MarketCount["Kalashnikov"])
		utils.Writeln(utils.CenterText("| Fireball          <%d> |", 75), values.MarketCount["Fireball"])
		utils.Writeln(utils.CenterText("| Iceball           <%d> |", 75), values.MarketCount["Iceball"])
		utils.Writeln(utils.CenterText("| Bomb              <%d> |", 75), values.MarketCount["Bomb"])
		utils.WriteColor(utils.CenterText("> Shield            <["+strconv.Itoa(values.MarketCount["Shield"])+"]> |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Heineken          <%d> |", 75), values.MarketCount["Heineken"])
		utils.Writeln(utils.CenterText("| Confirmer les achats    |", 75))
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 7:
		utils.Writeln(utils.CenterText("| Heal              <%d> |", 75), values.MarketCount["Heal"])
		utils.Writeln(utils.CenterText("| Poison            <%d> |", 75), values.MarketCount["Poison"])
		utils.Writeln(utils.CenterText("| Kalashnikov       <%d> |", 75), values.MarketCount["Kalashnikov"])
		utils.Writeln(utils.CenterText("| Fireball          <%d> |", 75), values.MarketCount["Fireball"])
		utils.Writeln(utils.CenterText("| Iceball           <%d> |", 75), values.MarketCount["Iceball"])
		utils.Writeln(utils.CenterText("| Bomb              <%d> |", 75), values.MarketCount["Bomb"])
		utils.Writeln(utils.CenterText("| Shield            <%d> |", 75), values.MarketCount["Shield"])
		utils.WriteColor(utils.CenterText("> Heineken          <["+strconv.Itoa(values.MarketCount["Heineken"])+"]> |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Confirmer les achats    |", 75))
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 8:
		utils.Writeln(utils.CenterText("| Heal              <%d> |", 75), values.MarketCount["Heal"])
		utils.Writeln(utils.CenterText("| Poison            <%d> |", 75), values.MarketCount["Poison"])
		utils.Writeln(utils.CenterText("| Kalashnikov       <%d> |", 75), values.MarketCount["Kalashnikov"])
		utils.Writeln(utils.CenterText("| Fireball          <%d> |", 75), values.MarketCount["Fireball"])
		utils.Writeln(utils.CenterText("| Iceball           <%d> |", 75), values.MarketCount["Iceball"])
		utils.Writeln(utils.CenterText("| Bomb              <%d> |", 75), values.MarketCount["Bomb"])
		utils.Writeln(utils.CenterText("| Shield            <%d> |", 75), values.MarketCount["Shield"])
		utils.Writeln(utils.CenterText("| Heineken          <%d> |", 75), values.MarketCount["Heineken"])
		utils.WriteColor(utils.CenterText("> Confirmer les achats    |", 75), "green")
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 9:
		utils.Writeln(utils.CenterText("| Heal              <%d> |", 75), values.MarketCount["Heal"])
		utils.Writeln(utils.CenterText("| Poison            <%d> |", 75), values.MarketCount["Poison"])
		utils.Writeln(utils.CenterText("| Kalashnikov       <%d> |", 75), values.MarketCount["Kalashnikov"])
		utils.Writeln(utils.CenterText("| Fireball          <%d> |", 75), values.MarketCount["Fireball"])
		utils.Writeln(utils.CenterText("| Iceball           <%d> |", 75), values.MarketCount["Iceball"])
		utils.Writeln(utils.CenterText("| Bomb              <%d> |", 75), values.MarketCount["Bomb"])
		utils.Writeln(utils.CenterText("| Shield            <%d> |", 75), values.MarketCount["Shield"])
		utils.Writeln(utils.CenterText("| Heineken          <%d> |", 75), values.MarketCount["Heineken"])
		utils.Writeln(utils.CenterText("| Confirmer les achats    |", 75))
		utils.WriteColor(utils.CenterText("> Retour                  |", 75), "yellow")
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	}
	utils.Writeln("")

	utils.Writeln(utils.CenterText("Montant de votre panier : %d €", 75), CalcMarketBuyPrice())
	utils.Writeln(utils.CenterText("Vous avez "+strconv.Itoa(values.Player.Inventory.Money)+"€", 75))
}

func CalcMarketBuyPrice() int {
	price := 0
	for name, value := range values.MarketCount {
		price += value * values.MarketPrice[name]
	}
	return price
}

func DisplaySellMarket() {
	utils.ClearTerminal()

	utils.Writeln("")
	utils.Writeln("")
	utils.Writeln(ascii.MenuDogsTitle)

	utils.Writeln(utils.CenterText("========= Vendre =========", 75))
	switch values.CurrentOption {
	case 0:
		utils.WriteColor(utils.CenterText("> Heal              <["+strconv.Itoa(values.MarketSell["Heal"])+"]> |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Poison            <%d> |", 75), values.MarketSell["Poison"])
		utils.Writeln(utils.CenterText("| Kalashnikov       <%d> |", 75), values.MarketSell["Kalashnikov"])
		utils.Writeln(utils.CenterText("| Fireball          <%d> |", 75), values.MarketSell["Fireball"])
		utils.Writeln(utils.CenterText("| Iceball           <%d> |", 75), values.MarketSell["Iceball"])
		utils.Writeln(utils.CenterText("| Bomb              <%d> |", 75), values.MarketSell["Bomb"])
		utils.Writeln(utils.CenterText("| Shield            <%d> |", 75), values.MarketSell["Shield"])
		utils.Writeln(utils.CenterText("| Heineken          <%d> |", 75), values.MarketSell["Heineken"])
		utils.Writeln(utils.CenterText("| Confirmer les ventes    |", 75))
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 1:
		utils.Writeln(utils.CenterText("| Heal              <%d> |", 75), values.MarketSell["Heal"])
		utils.WriteColor(utils.CenterText("> Poison            <["+strconv.Itoa(values.MarketSell["Poison"])+"]> |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Kalashnikov       <%d> |", 75), values.MarketSell["Kalashnikov"])
		utils.Writeln(utils.CenterText("| Fireball          <%d> |", 75), values.MarketSell["Fireball"])
		utils.Writeln(utils.CenterText("| Iceball           <%d> |", 75), values.MarketSell["Iceball"])
		utils.Writeln(utils.CenterText("| Bomb              <%d> |", 75), values.MarketSell["Bomb"])
		utils.Writeln(utils.CenterText("| Shield            <%d> |", 75), values.MarketSell["Shield"])
		utils.Writeln(utils.CenterText("| Heineken          <%d> |", 75), values.MarketSell["Heineken"])
		utils.Writeln(utils.CenterText("| Confirmer les ventes    |", 75))
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 2:
		utils.Writeln(utils.CenterText("| Heal              <%d> |", 75), values.MarketSell["Heal"])
		utils.Writeln(utils.CenterText("| Poison            <%d> |", 75), values.MarketSell["Poison"])
		utils.WriteColor(utils.CenterText("> Kalashnikov       <["+strconv.Itoa(values.MarketSell["Kalashnikov"])+"]> |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Fireball          <%d> |", 75), values.MarketSell["Fireball"])
		utils.Writeln(utils.CenterText("| Iceball           <%d> |", 75), values.MarketSell["Iceball"])
		utils.Writeln(utils.CenterText("| Bomb              <%d> |", 75), values.MarketSell["Bomb"])
		utils.Writeln(utils.CenterText("| Shield            <%d> |", 75), values.MarketSell["Shield"])
		utils.Writeln(utils.CenterText("| Heineken          <%d> |", 75), values.MarketSell["Heineken"])
		utils.Writeln(utils.CenterText("| Confirmer les ventes    |", 75))
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 3:
		utils.Writeln(utils.CenterText("| Heal              <%d> |", 75), values.MarketSell["Heal"])
		utils.Writeln(utils.CenterText("| Poison            <%d> |", 75), values.MarketSell["Poison"])
		utils.Writeln(utils.CenterText("| Kalashnikov       <%d> |", 75), values.MarketSell["Kalashnikov"])
		utils.WriteColor(utils.CenterText("> Fireball          <["+strconv.Itoa(values.MarketSell["Fireball"])+"]> |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Iceball           <%d> |", 75), values.MarketSell["Iceball"])
		utils.Writeln(utils.CenterText("| Bomb              <%d> |", 75), values.MarketSell["Bomb"])
		utils.Writeln(utils.CenterText("| Shield            <%d> |", 75), values.MarketSell["Shield"])
		utils.Writeln(utils.CenterText("| Heineken          <%d> |", 75), values.MarketSell["Heineken"])
		utils.Writeln(utils.CenterText("| Confirmer les ventes    |", 75))
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 4:
		utils.Writeln(utils.CenterText("| Heal              <%d> |", 75), values.MarketSell["Heal"])
		utils.Writeln(utils.CenterText("| Poison            <%d> |", 75), values.MarketSell["Poison"])
		utils.Writeln(utils.CenterText("| Kalashnikov       <%d> |", 75), values.MarketSell["Kalashnikov"])
		utils.Writeln(utils.CenterText("| Fireball          <%d> |", 75), values.MarketSell["Fireball"])
		utils.WriteColor(utils.CenterText("> Iceball           <["+strconv.Itoa(values.MarketSell["Iceball"])+"]> |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Bomb              <%d> |", 75), values.MarketSell["Bomb"])
		utils.Writeln(utils.CenterText("| Shield            <%d> |", 75), values.MarketSell["Shield"])
		utils.Writeln(utils.CenterText("| Heineken          <%d> |", 75), values.MarketSell["Heineken"])
		utils.Writeln(utils.CenterText("| Confirmer les ventes    |", 75))
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 5:
		utils.Writeln(utils.CenterText("| Heal              <%d> |", 75), values.MarketSell["Heal"])
		utils.Writeln(utils.CenterText("| Poison            <%d> |", 75), values.MarketSell["Poison"])
		utils.Writeln(utils.CenterText("| Kalashnikov       <%d> |", 75), values.MarketSell["Kalashnikov"])
		utils.Writeln(utils.CenterText("| Fireball          <%d> |", 75), values.MarketSell["Fireball"])
		utils.Writeln(utils.CenterText("| Iceball           <%d> |", 75), values.MarketSell["Iceball"])
		utils.WriteColor(utils.CenterText("> Bomb              <["+strconv.Itoa(values.MarketSell["Bomb"])+"]> |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Shield            <%d> |", 75), values.MarketSell["Shield"])
		utils.Writeln(utils.CenterText("| Heineken          <%d> |", 75), values.MarketSell["Heineken"])
		utils.Writeln(utils.CenterText("| Confirmer les ventes    |", 75))
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 6:
		utils.Writeln(utils.CenterText("| Heal              <%d> |", 75), values.MarketSell["Heal"])
		utils.Writeln(utils.CenterText("| Poison            <%d> |", 75), values.MarketSell["Poison"])
		utils.Writeln(utils.CenterText("| Kalashnikov       <%d> |", 75), values.MarketSell["Kalashnikov"])
		utils.Writeln(utils.CenterText("| Fireball          <%d> |", 75), values.MarketSell["Fireball"])
		utils.Writeln(utils.CenterText("| Iceball           <%d> |", 75), values.MarketSell["Iceball"])
		utils.Writeln(utils.CenterText("| Bomb              <%d> |", 75), values.MarketSell["Bomb"])
		utils.WriteColor(utils.CenterText("> Shield            <["+strconv.Itoa(values.MarketSell["Shield"])+"]> |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Heineken          <%d> |", 75), values.MarketSell["Heineken"])
		utils.Writeln(utils.CenterText("| Confirmer les ventes    |", 75))
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 7:
		utils.Writeln(utils.CenterText("| Heal              <%d> |", 75), values.MarketSell["Heal"])
		utils.Writeln(utils.CenterText("| Poison            <%d> |", 75), values.MarketSell["Poison"])
		utils.Writeln(utils.CenterText("| Kalashnikov       <%d> |", 75), values.MarketSell["Kalashnikov"])
		utils.Writeln(utils.CenterText("| Fireball          <%d> |", 75), values.MarketSell["Fireball"])
		utils.Writeln(utils.CenterText("| Iceball           <%d> |", 75), values.MarketSell["Iceball"])
		utils.Writeln(utils.CenterText("| Bomb              <%d> |", 75), values.MarketSell["Bomb"])
		utils.Writeln(utils.CenterText("| Shield            <%d> |", 75), values.MarketSell["Shield"])
		utils.WriteColor(utils.CenterText("> Heineken          <["+strconv.Itoa(values.MarketSell["Heineken"])+"]> |", 75), "yellow")
		utils.Writeln(utils.CenterText("| Confirmer les ventes    |", 75))
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 8:
		utils.Writeln(utils.CenterText("| Heal              <%d> |", 75), values.MarketSell["Heal"])
		utils.Writeln(utils.CenterText("| Poison            <%d> |", 75), values.MarketSell["Poison"])
		utils.Writeln(utils.CenterText("| Kalashnikov       <%d> |", 75), values.MarketSell["Kalashnikov"])
		utils.Writeln(utils.CenterText("| Fireball          <%d> |", 75), values.MarketSell["Fireball"])
		utils.Writeln(utils.CenterText("| Iceball           <%d> |", 75), values.MarketSell["Iceball"])
		utils.Writeln(utils.CenterText("| Bomb              <%d> |", 75), values.MarketSell["Bomb"])
		utils.Writeln(utils.CenterText("| Shield            <%d> |", 75), values.MarketSell["Shield"])
		utils.Writeln(utils.CenterText("| Heineken          <%d> |", 75), values.MarketSell["Heineken"])
		utils.WriteColor(utils.CenterText("> Confirmer les ventes    |", 75), "green")
		utils.Writeln(utils.CenterText("| Retour                  |", 75))
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	case 9:
		utils.Writeln(utils.CenterText("| Heal              <%d> |", 75), values.MarketSell["Heal"])
		utils.Writeln(utils.CenterText("| Poison            <%d> |", 75), values.MarketSell["Poison"])
		utils.Writeln(utils.CenterText("| Kalashnikov       <%d> |", 75), values.MarketSell["Kalashnikov"])
		utils.Writeln(utils.CenterText("| Fireball          <%d> |", 75), values.MarketSell["Fireball"])
		utils.Writeln(utils.CenterText("| Iceball           <%d> |", 75), values.MarketSell["Iceball"])
		utils.Writeln(utils.CenterText("| Bomb              <%d> |", 75), values.MarketSell["Bomb"])
		utils.Writeln(utils.CenterText("| Shield            <%d> |", 75), values.MarketSell["Shield"])
		utils.Writeln(utils.CenterText("| Heineken          <%d> |", 75), values.MarketSell["Heineken"])
		utils.Writeln(utils.CenterText("| Confirmer les ventes    |", 75))
		utils.WriteColor(utils.CenterText("> Retour                  |", 75), "yellow")
		utils.Writeln(utils.CenterText("===========================", 75))
		break
	}
	utils.Writeln("")
	utils.Writeln(utils.CenterText("Montant de vos vente : %d €", 75), CalcMarketSellPrice())
	utils.Writeln(utils.CenterText("Vous avez "+strconv.Itoa(values.Player.Inventory.Money)+"€", 75))
}

func CalcMarketSellPrice() int {
	price := 0
	for name, value := range values.MarketSell {
		price += value * (values.MarketPrice[name] / 2)
	}
	return price
}
