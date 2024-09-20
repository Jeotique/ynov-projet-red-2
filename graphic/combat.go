package graphic

import (
	"os"
	"projet/ascii"
	"projet/bdd"
	"projet/utils"
	"projet/values"
	"strconv"

	"github.com/olekukonko/tablewriter"
)

func DisplayCombat() {

	utils.ClearTerminal()

	MyDog := values.Player.Dogs[values.Player.CurrentDog]

	data := [][]string{
		[]string{ascii.DogsArt[MyDog.Race], ascii.DogsArt[values.Ennemy.Race]},
		[]string{"\n\n", "\n\n"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	for _, v := range data {
		table.Append(v)
	}
	table.SetFooter([]string{"Race : " + MyDog.Race + "\nVie : " + strconv.Itoa(MyDog.Health) + "/" + strconv.Itoa(MyDog.MaxHealth) + "\nNom : " + MyDog.Name,
		"Type : " + values.Ennemy.Race + "\nVie : " + strconv.Itoa(values.Ennemy.Health) + "/" + strconv.Itoa(values.Ennemy.MaxHealth) + "\nNom : " + values.Ennemy.Name,
	})
	table.Render()

	values.BattleOptions = []string{"Attaquer"}

	// Ajouter les values.BattleOptions selon les objets dans l'inventaire
	if values.Player.Inventory.Heal > 0 {
		values.BattleOptions = append(values.BattleOptions, "Soigner")
	}
	if values.Player.Inventory.Poison > 0 {
		values.BattleOptions = append(values.BattleOptions, "Poison")
	}
	if values.Player.Inventory.Kalashnikov {
		values.BattleOptions = append(values.BattleOptions, "Kalashnikov")
	}
	if values.Player.Inventory.Fireball > 0 {
		values.BattleOptions = append(values.BattleOptions, "Boule de Feu")
	}
	if values.Player.Inventory.Iceball > 0 {
		values.BattleOptions = append(values.BattleOptions, "Boule de Glace")
	}
	if values.Player.Inventory.Bomb > 0 {
		values.BattleOptions = append(values.BattleOptions, "Bombe")
	}
	if values.Player.Inventory.Shield > 0 {
		values.BattleOptions = append(values.BattleOptions, "Bouclier")
	}
	if values.Player.Inventory.Heineken > 0 {
		values.BattleOptions = append(values.BattleOptions, "Heineken")
	}

	// Afficher les values.BattleOptions
	for i, option := range values.BattleOptions {
		if i == values.CurrentOption {
			utils.WriteColor(utils.CenterText("> "+option, 2), "yellow")
		} else {
			utils.Writeln(utils.CenterText(option, 2))
		}
	}
}

func DisplayPlayerAttack() {
	utils.ClearTerminal()

	MyDog := values.Player.Dogs[values.Player.CurrentDog]

	data := [][]string{
		[]string{ascii.DogsArt[MyDog.Race], ascii.DogsArt[values.Ennemy.Race]},
		[]string{"\n\n", "\n\n"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	for _, v := range data {
		table.Append(v)
	}
	table.SetFooter([]string{"Race : " + MyDog.Race + "\nVie : " + strconv.Itoa(MyDog.Health) + "/" + strconv.Itoa(MyDog.MaxHealth) + "\nNom : " + MyDog.Name,
		"Type : " + values.Ennemy.Race + "\nVie : " + strconv.Itoa(values.Ennemy.Health) + "/" + strconv.Itoa(values.Ennemy.MaxHealth) + "\nNom : " + values.Ennemy.Name,
	})
	table.Render()

	if values.EnnemyDodge {
		utils.WriteanimColor(utils.CenterText("L'ennemi a esquivé votre attaque", 2), 10, "red")
	} else {
		if values.PlayerCritical {
			utils.WriteanimColor(utils.CenterText("Vous avez infligé "+strconv.Itoa(values.PlayerDamage)+" dégats (coup critique)", 2), 10, "green")
		} else {
			utils.WriteanimColor(utils.CenterText("Vous avez infligé "+strconv.Itoa(values.PlayerDamage)+" dégats", 2), 10, "green")
		}
	}
}

func DisplayEnnemyAttack() {
	utils.ClearTerminal()

	MyDog := values.Player.Dogs[values.Player.CurrentDog]

	data := [][]string{
		[]string{ascii.DogsArt[MyDog.Race], ascii.DogsArt[values.Ennemy.Race]},
		[]string{"\n\n", "\n\n"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	for _, v := range data {
		table.Append(v)
	}
	table.SetFooter([]string{"Race : " + MyDog.Race + "\nVie : " + strconv.Itoa(MyDog.Health) + "/" + strconv.Itoa(MyDog.MaxHealth) + "\nNom : " + MyDog.Name,
		"Type : " + values.Ennemy.Race + "\nVie : " + strconv.Itoa(values.Ennemy.Health) + "/" + strconv.Itoa(values.Ennemy.MaxHealth) + "\nNom : " + values.Ennemy.Name,
	})
	table.Render()

	if values.PlayerDodge {
		utils.WriteanimColor(utils.CenterText("Vous avez esquivé l'attaque ennemi", 2), 10, "green")
	} else {
		if values.EnnemyCritical {
			utils.WriteanimColor(utils.CenterText("L'ennemi vous a infligé "+strconv.Itoa(values.EnnemyDamage)+" dégats (coup critique)", 2), 10, "red")
		} else {
			utils.WriteanimColor(utils.CenterText("L'ennemi vous a infligé "+strconv.Itoa(values.EnnemyDamage)+" dégats", 2), 10, "red")
		}
	}
}

func DisplayPlayerHeal() {
	utils.ClearTerminal()

	MyDog := values.Player.Dogs[values.Player.CurrentDog]

	data := [][]string{
		[]string{ascii.DogsArt[MyDog.Race], ascii.DogsArt[values.Ennemy.Race]},
		[]string{"\n\n", "\n\n"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	for _, v := range data {
		table.Append(v)
	}
	table.SetFooter([]string{"Race : " + MyDog.Race + "\nVie : " + strconv.Itoa(MyDog.Health) + "/" + strconv.Itoa(MyDog.MaxHealth) + "\nNom : " + MyDog.Name,
		"Type : " + values.Ennemy.Race + "\nVie : " + strconv.Itoa(values.Ennemy.Health) + "/" + strconv.Itoa(values.Ennemy.MaxHealth) + "\nNom : " + values.Ennemy.Name,
	})
	table.Render()

	utils.WriteanimColor(utils.CenterText("Vous avez soigné votre chien de "+strconv.Itoa(values.PlayerHeal)+" points de vie", 2), 10, "green")
}

func DisplayPlayerPoison() {
	utils.ClearTerminal()

	MyDog := values.Player.Dogs[values.Player.CurrentDog]

	data := [][]string{
		[]string{ascii.DogsArt[MyDog.Race], ascii.DogsArt[values.Ennemy.Race]},
		[]string{"\n\n", "\n\n"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	for _, v := range data {
		table.Append(v)
	}
	table.SetFooter([]string{"Race : " + MyDog.Race + "\nVie : " + strconv.Itoa(MyDog.Health) + "/" + strconv.Itoa(MyDog.MaxHealth) + "\nNom : " + MyDog.Name,
		"Type : " + values.Ennemy.Race + "\nVie : " + strconv.Itoa(values.Ennemy.Health) + "/" + strconv.Itoa(values.Ennemy.MaxHealth) + "\nNom : " + values.Ennemy.Name,
	})
	table.Render()

	utils.WriteanimColor(utils.CenterText("Vous avez empoisonné l'ennemi, il perdra 5 points de vie à chaque tour", 2), 10, "green")
}

func DisplayEnnemyPoison() {
	utils.ClearTerminal()

	MyDog := values.Player.Dogs[values.Player.CurrentDog]

	data := [][]string{
		[]string{ascii.DogsArt[MyDog.Race], ascii.DogsArt[values.Ennemy.Race]},
		[]string{"\n\n", "\n\n"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	for _, v := range data {
		table.Append(v)
	}
	table.SetFooter([]string{"Race : " + MyDog.Race + "\nVie : " + strconv.Itoa(MyDog.Health) + "/" + strconv.Itoa(MyDog.MaxHealth) + "\nNom : " + MyDog.Name,
		"Type : " + values.Ennemy.Race + "\nVie : " + strconv.Itoa(values.Ennemy.Health) + "/" + strconv.Itoa(values.Ennemy.MaxHealth) + "\nNom : " + values.Ennemy.Name,
	})
	table.Render()

	utils.WriteanimColor(utils.CenterText("L'ennemi a été empoisonné, il perd 5 points de vie", 2), 10, "green")
}

func DisplayPlayerBoost() {
	utils.ClearTerminal()

	MyDog := values.Player.Dogs[values.Player.CurrentDog]

	data := [][]string{
		[]string{ascii.DogsArt[MyDog.Race], ascii.DogsArt[values.Ennemy.Race]},
		[]string{"\n\n", "\n\n"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	for _, v := range data {
		table.Append(v)
	}
	table.SetFooter([]string{"Race : " + MyDog.Race + "\nVie : " + strconv.Itoa(MyDog.Health) + "/" + strconv.Itoa(MyDog.MaxHealth) + "\nNom : " + MyDog.Name,
		"Type : " + values.Ennemy.Race + "\nVie : " + strconv.Itoa(values.Ennemy.Health) + "/" + strconv.Itoa(values.Ennemy.MaxHealth) + "\nNom : " + values.Ennemy.Name,
	})
	table.Render()

	utils.WriteanimColor(utils.CenterText("Vous avez boosté votre chien à base de heinekein, le prochain coup sera critique", 2), 10, "green")
}

func PlayerWin() {
	utils.ClearTerminal()
	utils.WriteColorLn(ascii.DogsArt[values.Ennemy.Race], "blue")
	utils.WriteanimColorln(utils.CenterText("Vous avez gagné le combat", 2), 10, "green")

	if values.Training {
		values.Player.Dogs[values.Player.CurrentDog].Experience += 3
		utils.WriteanimColorln(utils.CenterText("Votre chien a gagné 3 points d'expérience", 2), 10, "green")
		values.Player.Dogs[values.Player.CurrentDog].Health = values.Player.Dogs[values.Player.CurrentDog].MaxHealth
		bdd.Database.SavePlayer(values.Player)
		values.Training = false
	} else {

		if len(values.Player.Dogs) < 2 {
			values.Player.Dogs[values.Player.CurrentDog].Health = values.Player.Dogs[values.Player.CurrentDog].MaxHealth / 2
		}

		Money := MoneyRandom()
		utils.WriteanimColorln(utils.CenterText("Vous avez gagné "+strconv.Itoa(Money)+" euros", 2), 10, "green")
		values.Player.Inventory.Money += Money
		if WinWood() {
			Wood := WoodRandom()
			utils.WriteanimColorln(utils.CenterText("Vous avez gagné "+strconv.Itoa(Wood)+" bois", 2), 10, "green")
			values.Player.Inventory.Wood += Wood
		}
		if WinViscous() {
			Viscous := ViscousRandom()
			utils.WriteanimColorln(utils.CenterText("Vous avez gagné "+strconv.Itoa(Viscous)+" visqueux", 2), 10, "green")
			values.Player.Inventory.Viscous += Viscous
		}
		if WinStone() {
			Stone := StoneRandom()
			utils.WriteanimColorln(utils.CenterText("Vous avez gagné "+strconv.Itoa(Stone)+" pierres", 2), 10, "green")
			values.Player.Inventory.Stone += Stone
		}
		bdd.Database.SavePlayer(values.Player)
	}
}

func MoneyRandom() int {
	return utils.RandomNumber(10, 30)
}

func WinWood() bool {
	return utils.RandomNumber(1, 100) <= 10
}

func WinViscous() bool {
	return utils.RandomNumber(1, 100) <= 10
}

func WinStone() bool {
	return utils.RandomNumber(1, 100) <= 10
}

func WoodRandom() int {
	return utils.RandomNumber(1, 5)
}

func ViscousRandom() int {
	return utils.RandomNumber(1, 5)
}

func StoneRandom() int {
	return utils.RandomNumber(1, 5)
}

func PlayerLose() {
	utils.ClearTerminal()
	utils.WriteColorLn(ascii.DogsArt[values.Ennemy.Race], "red")
	utils.WriteanimColorln(utils.CenterText("Vous avez perdu le combat", 2), 10, "green")
	if values.Training {
		values.Player.Dogs[values.Player.CurrentDog].Experience += 1
		utils.WriteanimColorln(utils.CenterText("Votre chien a gagné 1 points d'expérience", 2), 10, "green")
		values.Training = false
		values.Player.Dogs[values.Player.CurrentDog].Health = values.Player.Dogs[values.Player.CurrentDog].MaxHealth
	} else {
		if len(values.Player.Dogs) < 2 {
			values.Player.Dogs[values.Player.CurrentDog].Health = values.Player.Dogs[values.Player.CurrentDog].MaxHealth / 2
		}
	}
	bdd.Database.SavePlayer(values.Player)
}

func DisplayPlayerKalashnikov() {
	utils.ClearTerminal()

	MyDog := values.Player.Dogs[values.Player.CurrentDog]

	data := [][]string{
		[]string{ascii.DogsArt[MyDog.Race], ascii.DogsArt[values.Ennemy.Race]},
		[]string{"\n\n", "\n\n"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	for _, v := range data {
		table.Append(v)
	}
	table.SetFooter([]string{"Race : " + MyDog.Race + "\nVie : " + strconv.Itoa(MyDog.Health) + "/" + strconv.Itoa(MyDog.MaxHealth) + "\nNom : " + MyDog.Name,
		"Type : " + values.Ennemy.Race + "\nVie : " + strconv.Itoa(values.Ennemy.Health) + "/" + strconv.Itoa(values.Ennemy.MaxHealth) + "\nNom : " + values.Ennemy.Name,
	})
	table.Render()

	utils.WriteanimColor(utils.CenterText("Vous avez utilisé votre kalashnikov, l'ennemi va sombrer", 2), 10, "green")
}

func DisplayPlayerFireball() {
	utils.ClearTerminal()

	MyDog := values.Player.Dogs[values.Player.CurrentDog]

	data := [][]string{
		[]string{ascii.DogsArt[MyDog.Race], ascii.DogsArt[values.Ennemy.Race]},
		[]string{"\n\n", "\n\n"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	for _, v := range data {
		table.Append(v)
	}
	table.SetFooter([]string{"Race : " + MyDog.Race + "\nVie : " + strconv.Itoa(MyDog.Health) + "/" + strconv.Itoa(MyDog.MaxHealth) + "\nNom : " + MyDog.Name,
		"Type : " + values.Ennemy.Race + "\nVie : " + strconv.Itoa(values.Ennemy.Health) + "/" + strconv.Itoa(values.Ennemy.MaxHealth) + "\nNom : " + values.Ennemy.Name,
	})
	table.Render()

	utils.WriteanimColor(utils.CenterText("Vous avez lancé une boule de feu, l'ennemi va brûler (10 points de vie)", 2), 10, "green")
}

func DisplayPlayerIceball() {
	utils.ClearTerminal()

	MyDog := values.Player.Dogs[values.Player.CurrentDog]

	data := [][]string{
		[]string{ascii.DogsArt[MyDog.Race], ascii.DogsArt[values.Ennemy.Race]},
		[]string{"\n\n", "\n\n"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	for _, v := range data {
		table.Append(v)
	}
	table.SetFooter([]string{"Race : " + MyDog.Race + "\nVie : " + strconv.Itoa(MyDog.Health) + "/" + strconv.Itoa(MyDog.MaxHealth) + "\nNom : " + MyDog.Name,
		"Type : " + values.Ennemy.Race + "\nVie : " + strconv.Itoa(values.Ennemy.Health) + "/" + strconv.Itoa(values.Ennemy.MaxHealth) + "\nNom : " + values.Ennemy.Name,
	})
	table.Render()

	utils.WriteanimColor(utils.CenterText("Vous avez lancé une boule de glace, l'ennemi va geler (5 points de vie)", 2), 10, "green")
}

func DisplayPlayerBomb() {
	utils.ClearTerminal()

	MyDog := values.Player.Dogs[values.Player.CurrentDog]

	data := [][]string{
		[]string{ascii.DogsArt[MyDog.Race], ascii.DogsArt[values.Ennemy.Race]},
		[]string{"\n\n", "\n\n"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	for _, v := range data {
		table.Append(v)
	}
	table.SetFooter([]string{"Race : " + MyDog.Race + "\nVie : " + strconv.Itoa(MyDog.Health) + "/" + strconv.Itoa(MyDog.MaxHealth) + "\nNom : " + MyDog.Name,
		"Type : " + values.Ennemy.Race + "\nVie : " + strconv.Itoa(values.Ennemy.Health) + "/" + strconv.Itoa(values.Ennemy.MaxHealth) + "\nNom : " + values.Ennemy.Name,
	})
	table.Render()

	utils.WriteanimColor(utils.CenterText("Vous avez lancé une bombe, l'ennemi va exploser (20 points de vie)", 2), 10, "green")
}

func DisplayPlayerShield() {
	utils.ClearTerminal()

	MyDog := values.Player.Dogs[values.Player.CurrentDog]

	data := [][]string{
		[]string{ascii.DogsArt[MyDog.Race], ascii.DogsArt[values.Ennemy.Race]},
		[]string{"\n\n", "\n\n"},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoWrapText(false)
	for _, v := range data {
		table.Append(v)
	}
	table.SetFooter([]string{"Race : " + MyDog.Race + "\nVie : " + strconv.Itoa(MyDog.Health) + "/" + strconv.Itoa(MyDog.MaxHealth) + "\nNom : " + MyDog.Name,
		"Type : " + values.Ennemy.Race + "\nVie : " + strconv.Itoa(values.Ennemy.Health) + "/" + strconv.Itoa(values.Ennemy.MaxHealth) + "\nNom : " + values.Ennemy.Name,
	})
	table.Render()

	utils.WriteanimColor(utils.CenterText("Vous avez utilisé un bouclier, vous esquiverez le prochain coup (et +10 points de vie)", 2), 10, "green")
}
