package values

import "projet/structs"

var CurrentPage string = "main_menu"
var CurrentOption int = 0
var CurrentOptionMax int = 1

var Player structs.Character

var MarketCount map[string]int = map[string]int{
	"Heal":        0,
	"Poison":      0,
	"Kalashnikov": 0,
	"Fireball":    0,
	"Iceball":     0,
	"Bomb":        0,
	"Shield":      0,
	"Heineken":    0,
}

var MarketSell map[string]int = map[string]int{
	"Heal":        0,
	"Poison":      0,
	"Kalashnikov": 0,
	"Fireball":    0,
	"Iceball":     0,
	"Bomb":        0,
	"Shield":      0,
	"Heineken":    0,
}

var MarketPrice map[string]int = map[string]int{
	"Heal":        10,
	"Poison":      15,
	"Kalashnikov": 6000,
	"Fireball":    20,
	"Iceball":     20,
	"Bomb":        30,
	"Shield":      25,
	"Heineken":    5,
}

var DogsList map[string]structs.Dog = map[string]structs.Dog{
	"Caniche":        structs.NewDog("Bastien", 1, 3, 10, 10, 0.1, "Caniche"),
	"Bulldog":        structs.NewDog("Brutus", 3, 6, 20, 20, 0.15, "Bulldog"),
	"Husky":          structs.NewDog("Fantôme", 5, 8, 30, 30, 0.2, "Husky"),
	"Labrador":       structs.NewDog("Buddy", 4, 7, 25, 25, 0.18, "Labrador"),
	"Rottweiler":     structs.NewDog("Maximus", 6, 10, 35, 35, 0.25, "Rottweiler"),
	"Beagle":         structs.NewDog("Snoopy", 2, 4, 15, 15, 0.12, "Beagle"),
	"BergerAllemand": structs.NewDog("Rex", 5, 9, 28, 28, 0.22, "BergerAllemand"),
	"Chihuahua":      structs.NewDog("Tiny", 1, 2, 8, 8, 0.3, "Chihuahua"),
	"CanicheRoyal":   structs.NewDog("Fluffy", 2, 5, 18, 18, 0.14, "CanicheRoyal"),
	"Dobermann":      structs.NewDog("Ace", 7, 12, 40, 40, 0.28, "Dobermann"),
	"Poney":          structs.NewDog("Gathan", 100, 100, 300, 300, 0.5, "Poney"),
	"BergerSuisse":   structs.NewDog("Rayden", 100, 100, 300, 300, 0.5, "BergerSuisse"),
}

var RobotDog structs.Dog = structs.NewDog("Robot", 4, 6, 50, 50, 0.4, "Robot")

type NichePrice struct {
	Wood    int
	Stone   int
	Viscous int
}

var NichesPrice map[int]NichePrice = map[int]NichePrice{
	0: {Wood: 0, Stone: 0, Viscous: 0},
	1: {Wood: 10, Stone: 5, Viscous: 2},
	2: {Wood: 20, Stone: 10, Viscous: 5},
	3: {Wood: 30, Stone: 15, Viscous: 8},
	4: {Wood: 40, Stone: 20, Viscous: 10},
	5: {Wood: 50, Stone: 25, Viscous: 12},
}

var BattleOptions []string
var Ennemy structs.Dog

var PlayerTurn bool = true
var PlayerDodge bool = false
var EnnemyDodge bool = false
var EnnemyCritical bool = false
var PlayerCritical bool = false
var PlayerDamage int = 0
var EnnemyDamage int = 0
var PlayerNextAttackCritical bool = false
var EnnemyNextAttackCritical bool = false
var PlayerHeal int = 0
var Training bool = false
var EnnemyPoison int = 0
var ForcePlayerDodge bool = false

var DogsPrice map[string]int = map[string]int{
	"Caniche":        1,
	"Bulldog":        90,
	"Husky":          60,
	"Labrador":       60,
	"Rottweiler":     50,
	"Beagle":         30,
	"BergerAllemand": 70,
	"Chihuahua":      8,
	"CanicheRoyal":   180,
	"Dobermann":      200,
	"Poney":          10000,
	"BergerSuisse":   300,
}
var DogsRace = []string{
	"Caniche",
	"Bulldog",
	"Husky",
	"Labrador",
	"Rottweiler",
	"Beagle",
	"BergerAllemand",
	"Chihuahua",
	"CanicheRoyal",
	"Dobermann",
	"Poney",
	"BergerSuisse",
}
