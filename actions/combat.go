package actions

import (
	"projet/graphic"
	"projet/sounds"
	"projet/utils"
	"projet/values"
	"time"
)

func CombatExec() {
	Action := values.BattleOptions[values.CurrentOption]

	switch Action {
	case "Attaquer":
		PlayerAttack()
	case "Soigner":
		if values.Player.Dogs[values.Player.CurrentDog].Health == values.Player.Dogs[values.Player.CurrentDog].MaxHealth {
			return
		}
		values.Player.Inventory.Heal--
		values.PlayerHeal = 5
		values.Player.Dogs[values.Player.CurrentDog].Heal(5)
		values.PlayerTurn = false
		graphic.DisplayPlayerHeal()
		time.Sleep(1 * time.Second)
		EnnemyAttack()
	case "Poison":
		values.Player.Inventory.Poison--
		values.EnnemyPoison = 3
		values.PlayerTurn = false
		graphic.DisplayPlayerPoison()
		time.Sleep(1 * time.Second)
		EnnemyAttack()
	case "Kalashnikov":
		values.Player.Inventory.Kalashnikov = false
		values.Ennemy.SetHealth(0)
		values.PlayerTurn = false
		graphic.DisplayPlayerKalashnikov()
		time.Sleep(1 * time.Second)
		CombatReset()
		values.CurrentPage = "combat_win"
		go utils.PlaySound(sounds.Win, 1)
		graphic.PlayerWin()
		time.Sleep(5 * time.Second)
		ReturnToMenu()
		return
	case "Boule de Feu":
		values.Player.Inventory.Fireball--
		values.Ennemy.TakeDamage(10)
		values.PlayerTurn = false
		graphic.DisplayPlayerFireball()
		time.Sleep(1 * time.Second)
		if !values.Ennemy.IsAlive() {
			CombatReset()
			values.CurrentPage = "combat_win"
			go utils.PlaySound(sounds.Win, 1)
			graphic.PlayerWin()
			time.Sleep(5 * time.Second)
			ReturnToMenu()
			return
		}
		EnnemyAttack()
	case "Boule de Glace":
		values.Player.Inventory.Iceball--
		values.Ennemy.TakeDamage(5)
		values.PlayerTurn = false
		graphic.DisplayPlayerIceball()
		time.Sleep(1 * time.Second)
		if !values.Ennemy.IsAlive() {
			CombatReset()
			values.CurrentPage = "combat_win"
			go utils.PlaySound(sounds.Win, 1)
			graphic.PlayerWin()
			time.Sleep(5 * time.Second)
			ReturnToMenu()
			return
		}
		EnnemyAttack()
	case "Bombe":
		values.Player.Inventory.Bomb--
		values.Ennemy.TakeDamage(20)
		values.PlayerTurn = false
		graphic.DisplayPlayerBomb()
		time.Sleep(1 * time.Second)
		if !values.Ennemy.IsAlive() {
			CombatReset()
			values.CurrentPage = "combat_win"
			go utils.PlaySound(sounds.Win, 1)
			graphic.PlayerWin()
			time.Sleep(5 * time.Second)
			ReturnToMenu()
			return
		}
		EnnemyAttack()
	case "Bouclier":
		values.Player.Inventory.Shield--
		values.ForcePlayerDodge = true
		values.PlayerTurn = false
		values.Player.Dogs[values.Player.CurrentDog].Heal(10)
		graphic.DisplayPlayerShield()
		time.Sleep(1 * time.Second)
		EnnemyAttack()
	case "Heineken":
		values.Player.Inventory.Heineken--
		values.PlayerNextAttackCritical = true
		values.PlayerTurn = false
		graphic.DisplayPlayerBoost()
		time.Sleep(1 * time.Second)
		EnnemyAttack()
		break
	}
}

func EnnemyAttack() {
	if values.Player.Dogs[values.Player.CurrentDog].Dodge() || values.ForcePlayerDodge {
		values.ForcePlayerDodge = false
		values.EnnemyCritical = false
		values.PlayerTurn = true
		values.PlayerDodge = true
		values.CurrentPage = "combat_enemy"
		graphic.DisplayEnnemyAttack()
		time.Sleep(1 * time.Second)
		if values.EnnemyPoison > 0 {
			values.EnnemyPoison--
			values.Ennemy.TakeDamage(5)
			graphic.DisplayEnnemyPoison()
			time.Sleep(1 * time.Second)
			if !values.Ennemy.IsAlive() {
				CombatReset()
				values.CurrentPage = "combat_win"
				go utils.PlaySound(sounds.Win, 1)
				graphic.PlayerWin()
				time.Sleep(5 * time.Second)
				ReturnToMenu()
				return
			}
		}
		values.CurrentPage = "combat"
		graphic.DisplayCombat()
		return
	}
	Critical := values.Ennemy.Critical()
	Damage := 0
	if !Critical {
		values.EnnemyCritical = false
		Damage = values.Ennemy.Attack()
		values.Player.Dogs[values.Player.CurrentDog].TakeDamage(Damage)
	} else {
		values.EnnemyCritical = true
		Damage = values.Ennemy.Attack() * 2
		values.Player.Dogs[values.Player.CurrentDog].TakeDamage(Damage)
	}
	values.EnnemyDamage = Damage
	values.PlayerTurn = true
	values.PlayerDodge = false
	values.CurrentPage = "combat_enemy"
	graphic.DisplayEnnemyAttack()
	time.Sleep(1 * time.Second)
	if !values.Player.Dogs[values.Player.CurrentDog].IsAlive() {
		CombatReset()
		values.CurrentPage = "combat_lose"
		go utils.PlaySound(sounds.Lose, 1)
		graphic.PlayerLose()
		time.Sleep(10 * time.Second)
		ReturnToMenu()
		return
	}
	if values.EnnemyPoison > 0 {
		values.EnnemyPoison--
		values.Ennemy.TakeDamage(5)
		graphic.DisplayEnnemyPoison()
		time.Sleep(1 * time.Second)
		if !values.Ennemy.IsAlive() {
			CombatReset()
			values.CurrentPage = "combat_win"
			go utils.PlaySound(sounds.Win, 1)
			graphic.PlayerWin()
			time.Sleep(5 * time.Second)
			ReturnToMenu()
			return
		}
	}
	values.CurrentPage = "combat"
	graphic.DisplayCombat()
}

func PlayerAttack() {
	Damage := 0
	values.EnnemyDodge = false
	if values.PlayerNextAttackCritical {
		values.PlayerCritical = true
		Damage = values.Player.Dogs[values.Player.CurrentDog].Attack() * 2
		values.Ennemy.TakeDamage(Damage)
		values.PlayerNextAttackCritical = false
	} else {
		if values.Ennemy.Dodge() {
			values.PlayerCritical = false
			values.PlayerTurn = false
			values.EnnemyDodge = true
			values.CurrentPage = "combat_player"
			graphic.DisplayPlayerAttack()
			time.Sleep(1 * time.Second)
			EnnemyAttack()
			return
		}
		if values.Player.Dogs[values.Player.CurrentDog].Critical() {
			values.PlayerCritical = true
			Damage = values.Player.Dogs[values.Player.CurrentDog].Attack() * 2
			values.Ennemy.TakeDamage(Damage)
		} else {
			values.PlayerCritical = false
			Damage = values.Player.Dogs[values.Player.CurrentDog].Attack()
			values.Ennemy.TakeDamage(Damage)
		}
	}
	values.PlayerDamage = Damage
	values.PlayerTurn = false
	values.CurrentPage = "combat_player"
	graphic.DisplayPlayerAttack()
	time.Sleep(1 * time.Second)

	if !values.Ennemy.IsAlive() {
		CombatReset()
		values.CurrentPage = "combat_win"
		go utils.PlaySound(sounds.Win, 1)
		graphic.PlayerWin()
		time.Sleep(5 * time.Second)
		ReturnToMenu()
		return
	}
	EnnemyAttack()
}

func CombatReset() {
	values.PlayerTurn = true
	values.PlayerDodge = false
	values.EnnemyDodge = false
	values.EnnemyCritical = false
	values.PlayerCritical = false
	values.PlayerDamage = 0
	values.EnnemyDamage = 0
	values.PlayerNextAttackCritical = false
	values.EnnemyNextAttackCritical = false
	values.PlayerHeal = 0
	//values.Training = false
	values.EnnemyPoison = 0
}

func ReturnToMenu() {
	values.CurrentOption = 0
	values.CurrentPage = "real_main_menu"
	values.CurrentOptionMax = 5
	graphic.RefreshRealMainMenu()
	go utils.PlaySound(sounds.RealMainMenu, 1)
}
