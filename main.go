package main

import (
	"log"
	"os"
	"projet/actions"
	"projet/ascii"
	"projet/bdd"
	"projet/graphic"
	"projet/sounds"
	"projet/utils"
	"projet/values"
	"time"

	"github.com/mattn/go-tty"
)

func InitBDD() {
	bdd.Database = bdd.NewQuickDB("data.json")
	bdd.Database.LoadPlayer()
}

func main() {
	pressF11()
	values.Player.Inventory.Money = 100
	InitBDD()
	utils.ClearTerminal()
	if values.Player.Name != "" {
		values.CurrentPage = "real_main_menu"
		go utils.PlaySound(sounds.RealMainMenu, 1)
		graphic.RefreshRealMainMenu()
		StartListening()
	} else {
		utils.Writeln("")
		utils.Writeln("")
		utils.Writeln(ascii.MenuTitle)

		go utils.PlaySound(sounds.MainMenu, 1)

		if len(os.Args) > 1 {
			CreateGameState()
			return
		} else {

			text := "        BMW > AUDI > MERCEDES > RENAULT > PEUGEOT > CITROEN > FIAT > FORD > TOYOTA > NISSAN > VOLKSWAGEN > OPEL > SEAT > SKODA > KIA > HYUNDAI > SUZUKI > MAZDA > HONDA > LEXUS > JAGUAR > LAND ROVER > TESLA > VOLVO > MINI > PORSCHE > ALFA ROMEO > MASERATI > LAMBORGHINI > FERRARI > BUGATTI > ASTON MARTIN > BENTLEY > ROLLS ROYCE > LOTUS > MCLAREN > KOENIGSEGG > PAGANI > SHELBY > DODGE > CHEVROLET > CADILLAC > CHRYSLER > JEEP > GMC > RAM"
			width := 100
			for {
				utils.ScrollTextAtLine(text, width, 15)
				break
			}

			utils.Writeanim("cheh", 3000)
			utils.Writeln("faut lancer avec un argument")
		}
	}
}

func CreateGameState() {
	utils.ClearTerminal()
	utils.Writeln("")
	utils.Writeln("")
	utils.Writeln(ascii.MenuDogsTitle)
	utils.Writeln("Bienvenue cher joueur ! Prêt à dresser des petit chiens mignon ?")
	utils.Writeln("Ici, respect des choix et envies sont notre priorité !")
	utils.Writeln("Nous allons commencer par créer ton personnage.")
	utils.Writeln("Quel est ton nom ?")
	utils.Writeln("")
	utils.Write("Nom : ")
	name := utils.WaitForInput()
	utils.Writeln("")
	utils.Writeln("Bienvenue " + name + " !")
	utils.Writeln("Nous allons maintenant choisir ton premier chien.")
	utils.Writeln("Voici les chiens disponibles :")
	utils.Writeln("")
	utils.Writeln("1. Caniche")
	utils.Writeln("2. Malinois")
	utils.Writeln("3. Pitbull")
	utils.Write("Choix : ")
	valid, choice := utils.WaitForNumberInput()
	if !valid || choice < 1 || choice > 3 {
		utils.Writeln("Choix invalide")
		CreateGameState()
		return
	}
	utils.Writeln("")
	utils.Writeln("Bon on ta choisi le chien numéro 1 par faute de budget !")
	utils.Writeln("Il va falloir choisir un nom pour ton chien, de toute façon en 2024 on peut les changer, mais pas ici !")
	utils.Writeln("")
	utils.Write("Nom (ou determinant): ")
	_ = utils.WaitForInput() //dogName := utils.WaitForInput()
	utils.Writeln("")
	utils.Writeln("Mouais pas ouf Bienvenue \"Bastien\" !")
	utils.Writeln("Nous allons maintenant choisir son sexe (facultatif).")
	utils.Writeln("")
	utils.Writeln("1. Mâle")
	utils.Writeln("2. Femelle")
	utils.Write("Choix : ")
	valid, choice = utils.WaitForNumberInput()
	if !valid || choice < 1 || choice > 2 {
		utils.Writeln("Choix invalide")
		CreateGameState()
		return
	}
	utils.Writeln("")
	utils.Writeln("Tu as choisi le sexe \"Non genré\" (ça ou airbus)")

	time.Sleep(2000 * time.Millisecond)

	utils.Writeln("")
	utils.Writeanim("Que l'aventure commence !", 100)
	values.Player.Name = name
	values.Player.Dogs = append(values.Player.Dogs, values.DogsList["Caniche"])
	values.Player.Inventory.Money = 50
	bdd.Database.SavePlayer(values.Player)
	values.CurrentPage = "real_main_menu"
	go utils.PlaySound(sounds.RealMainMenu, 1)
	graphic.RefreshRealMainMenu()
	StartListening()
}

func StartListening() {
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {
		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		if r == 65 || r == 66 || r == 67 || r == 68 { // 65 = fleche haut & 66 = fleche bas
			var maxIndex int
			var funcType func()
			//go utils.PlaySound(sounds.Navigate, 1)
			switch values.CurrentPage {
			case "main_menu":
				maxIndex = 1
				funcType = graphic.RefreshMainMenu
				break
			case "real_main_menu":
				maxIndex = 6
				funcType = graphic.RefreshRealMainMenu
				break
			case "inventory":
				maxIndex = 11
				funcType = graphic.DisplayInventory
				break
			case "who":
				maxIndex = 2
				funcType = graphic.DisplayWho
				break
			case "market":
				maxIndex = 2
				funcType = graphic.DisplayMarket
				break
			case "market_buy":
				maxIndex = 9
				funcType = graphic.DisplayBuyMarket
				break
			case "market_sell":
				maxIndex = 9
				funcType = graphic.DisplaySellMarket
				break
			case "ikea":
				maxIndex = 6
				funcType = graphic.DisplayIkea
				break
			case "niche":
				maxIndex = 2
				funcType = graphic.DisplayNiche
				break
			case "niche_dog":
				maxIndex = 3
				funcType = graphic.DisplayDogInfo
				break
			case "niche_change":
				maxIndex = len(values.Player.Dogs) + 1
				funcType = graphic.DiplayChangeDog
				break
			case "combat":
				maxIndex = len(values.BattleOptions) - 1
				funcType = graphic.DisplayCombat
				break
			case "niche_buy":
				maxIndex = len(values.DogsRace)
				funcType = graphic.DisplayBuyDog
				break
			case "niche_comp":
				maxIndex = 4
				funcType = graphic.DisplayCompetences
				break
			default:
				continue
			}
			if r == 65 {
				if values.CurrentOption == 0 {
					values.CurrentOption = maxIndex
				} else {
					values.CurrentOption--
				}
			} else if r == 66 {
				if values.CurrentOption == maxIndex {
					values.CurrentOption = 0
				} else {
					values.CurrentOption++
				}
			} else if r == 67 || r == 68 { // 67 = fleche droite & 68 = fleche gauche
				if values.CurrentPage == "market_buy" {
					itemsnames := []string{"Heal", "Poison", "Kalashnikov", "Fireball", "Iceball", "Bomb", "Shield", "Heineken"}
					if r == 68 {
						if values.MarketCount[itemsnames[values.CurrentOption]] > 0 {
							values.MarketCount[itemsnames[values.CurrentOption]]--
						} else {
							values.MarketCount[itemsnames[values.CurrentOption]] = 9
						}
					} else if r == 67 {
						if values.MarketCount[itemsnames[values.CurrentOption]] < 9 {
							values.MarketCount[itemsnames[values.CurrentOption]]++
						} else {
							values.MarketCount[itemsnames[values.CurrentOption]] = 0
						}
					}
				} else if values.CurrentPage == "market_sell" {
					itemsnames := []string{"Heal", "Poison", "Kalashnikov", "Fireball", "Iceball", "Bomb", "Shield", "Heineken"}
					if r == 68 {
						if values.MarketSell[itemsnames[values.CurrentOption]] > 0 {
							values.MarketSell[itemsnames[values.CurrentOption]]--
						} else {
							switch itemsnames[values.CurrentOption] {
							case "Heal":
								values.MarketSell[itemsnames[values.CurrentOption]] = values.Player.Inventory.Heal
								break
							case "Poison":
								values.MarketSell[itemsnames[values.CurrentOption]] = values.Player.Inventory.Poison
								break
							case "Kalashnikov":
								if values.Player.Inventory.Kalashnikov {
									values.MarketSell[itemsnames[values.CurrentOption]] = 1
								}
								break
							case "Fireball":
								values.MarketSell[itemsnames[values.CurrentOption]] = values.Player.Inventory.Fireball
								break
							case "Iceball":
								values.MarketSell[itemsnames[values.CurrentOption]] = values.Player.Inventory.Iceball
								break
							case "Bomb":
								values.MarketSell[itemsnames[values.CurrentOption]] = values.Player.Inventory.Bomb
								break
							case "Shield":
								values.MarketSell[itemsnames[values.CurrentOption]] = values.Player.Inventory.Shield
								break
							case "Heineken":
								values.MarketSell[itemsnames[values.CurrentOption]] = values.Player.Inventory.Heineken
								break
							}
						}
					} else if r == 67 {
						switch itemsnames[values.CurrentOption] {
						case "Heal":
							if values.MarketSell[itemsnames[values.CurrentOption]] < values.Player.Inventory.Heal {
								values.MarketSell[itemsnames[values.CurrentOption]]++
							} else {
								values.MarketSell[itemsnames[values.CurrentOption]] = 0
							}
							break
						case "Poison":
							if values.MarketSell[itemsnames[values.CurrentOption]] < values.Player.Inventory.Poison {
								values.MarketSell[itemsnames[values.CurrentOption]]++
							} else {
								values.MarketSell[itemsnames[values.CurrentOption]] = 0
							}
							break
						case "Kalashnikov":
							if values.MarketSell[itemsnames[values.CurrentOption]] < 1 {
								values.MarketCount[itemsnames[values.CurrentOption]]++
							} else {
								values.MarketCount[itemsnames[values.CurrentOption]] = 0
							}
							break
						case "Fireball":
							if values.MarketSell[itemsnames[values.CurrentOption]] < values.Player.Inventory.Fireball {
								values.MarketSell[itemsnames[values.CurrentOption]]++
							} else {
								values.MarketSell[itemsnames[values.CurrentOption]] = 0
							}
							break
						case "Iceball":
							if values.MarketSell[itemsnames[values.CurrentOption]] < values.Player.Inventory.Iceball {
								values.MarketSell[itemsnames[values.CurrentOption]]++
							} else {
								values.MarketSell[itemsnames[values.CurrentOption]] = 0
							}
							break
						case "Bomb":
							if values.MarketSell[itemsnames[values.CurrentOption]] < values.Player.Inventory.Bomb {
								values.MarketSell[itemsnames[values.CurrentOption]]++
							} else {
								values.MarketSell[itemsnames[values.CurrentOption]] = 0
							}
							break
						case "Shield":
							if values.MarketSell[itemsnames[values.CurrentOption]] < values.Player.Inventory.Shield {
								values.MarketSell[itemsnames[values.CurrentOption]]++
							} else {
								values.MarketSell[itemsnames[values.CurrentOption]] = 0
							}
							break
						case "Heineken":
							if values.MarketSell[itemsnames[values.CurrentOption]] < values.Player.Inventory.Heineken {
								values.MarketSell[itemsnames[values.CurrentOption]]++
							} else {
								values.MarketSell[itemsnames[values.CurrentOption]] = 0
							}
							break
						}
					}
				} else if values.CurrentPage == "niche_comp" {
					if r == 68 { // gauche
						switch values.CurrentOption {
						case 0:
							if values.Player.Dogs[values.Player.CurrentDog].MaxHealth > 5 {
								values.Player.Dogs[values.Player.CurrentDog].MaxHealth--
								values.Player.Dogs[values.Player.CurrentDog].Experience++
							}
							break
						case 1:
							if values.Player.Dogs[values.Player.CurrentDog].Damage > 1 {
								values.Player.Dogs[values.Player.CurrentDog].Damage--
								values.Player.Dogs[values.Player.CurrentDog].Experience++
							}
							break
						case 2:
							if values.Player.Dogs[values.Player.CurrentDog].MaxDamage > 1 {
								values.Player.Dogs[values.Player.CurrentDog].MaxDamage--
								values.Player.Dogs[values.Player.CurrentDog].Experience++
							}
							break
						case 3:
							if values.Player.Dogs[values.Player.CurrentDog].DodgeMultiplier > 0 {
								values.Player.Dogs[values.Player.CurrentDog].DodgeMultiplier--
								values.Player.Dogs[values.Player.CurrentDog].Experience++
							}
						}
					} else if r == 67 { // droite
						switch values.CurrentOption {
						case 0:
							if values.Player.Dogs[values.Player.CurrentDog].Experience > 0 {
								values.Player.Dogs[values.Player.CurrentDog].MaxHealth++
								values.Player.Dogs[values.Player.CurrentDog].Experience--
							}
							break
						case 1:
							if values.Player.Dogs[values.Player.CurrentDog].Experience > 0 {
								values.Player.Dogs[values.Player.CurrentDog].Damage++
								values.Player.Dogs[values.Player.CurrentDog].Experience--
							}
							break
						case 2:
							if values.Player.Dogs[values.Player.CurrentDog].Experience > 0 {
								values.Player.Dogs[values.Player.CurrentDog].MaxDamage++
								values.Player.Dogs[values.Player.CurrentDog].Experience--
							}
							break
						case 3:
							if values.Player.Dogs[values.Player.CurrentDog].Experience > 0 {
								values.Player.Dogs[values.Player.CurrentDog].DodgeMultiplier++
								values.Player.Dogs[values.Player.CurrentDog].Experience--
							}
							break
						}
					}
				}
			}
			funcType()
		} else {
			if r == 13 { // 13 = touche entrée...
				//go utils.PlaySound(sounds.Click, 1)
				switch values.CurrentPage {
				case "main_menu":
					actions.MenuExec()
					break
				case "real_main_menu":
					actions.RealMenuExec()
					break
				case "inventory":
					actions.InventoryExec()
					break
				case "who":
					actions.WhoExec()
					break
				case "market":
					actions.MarketExec()
					break
				case "market_buy":
					actions.MarketBuyExec()
					break
				case "market_sell":
					actions.MarketSellExec()
					break
				case "ikea":
					actions.IkeaExec()
					break
				case "niche":
					actions.NicheExec()
					break
				case "niche_dog":
					actions.NicheDogExec()
					break
				case "niche_change":
					actions.NicheChangeExec()
					break
				case "combat":
					actions.CombatExec()
					break
				case "niche_buy":
					actions.NicheBuyExec()
					break
				case "niche_comp":
					actions.NicheCompExec()
					break
				default:
					continue
				}
			}
		}
	}
}
