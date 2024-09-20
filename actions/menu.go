package actions

import (
	"os"
	"projet/graphic"
	"projet/sounds"
	"projet/utils"
	"projet/values"
	"syscall"
	"time"
)

var (
	user32               = syscall.NewLazyDLL("user32.dll")
	procKeybdEvent       = user32.NewProc("keybd_event")
	VK_F11          byte = 0x7A
	KEYEVENTF_KEYUP      = 0x0002
)

func pressF11() {
	procKeybdEvent.Call(uintptr(VK_F11), 0, 0, 0)
	time.Sleep(100 * time.Millisecond)
	procKeybdEvent.Call(uintptr(VK_F11), 0, uintptr(KEYEVENTF_KEYUP), 0)
}

func MenuExec() {
	switch values.CurrentOption {
	case 0:
		values.CurrentOption = 0
		values.CurrentPage = "real_main_menu"
		values.CurrentOptionMax = 5
		graphic.RefreshRealMainMenu()
		go utils.PlaySound(sounds.RealMainMenu, 1)
		break
	case 1:
		utils.ClearTerminal()
		os.Exit(0)
		break
	}
}

func RealMenuExec() {
	switch values.CurrentOption {
	case 0:
		if values.Player.Dogs[values.Player.CurrentDog].IsAlive() {
			values.Training = false
			values.Ennemy = utils.GetRandomDogClone(values.DogsList)
			values.CurrentOption = 0
			values.CurrentPage = "combat"
			values.CurrentOptionMax = -1
			graphic.DisplayCombat()
			go utils.PlaySound(sounds.Battle, 1)
		} else {
			utils.WriteColorLn(utils.CenterText("Votre chien est mort", 75), "red")
		}
		break
	case 1:
		values.CurrentOption = 0
		values.CurrentPage = "market"
		values.CurrentOptionMax = 5
		graphic.DisplayMarket()
		break
	case 2:
		values.CurrentOption = 0
		values.CurrentPage = "inventory"
		values.CurrentOptionMax = 9
		graphic.DisplayInventory()
		break
	case 3:
		values.CurrentOption = 0
		values.CurrentPage = "ikea"
		values.CurrentOptionMax = 5
		graphic.DisplayIkea()
		break
	case 4:
		values.CurrentOption = 0
		values.CurrentPage = "niche"
		values.CurrentOptionMax = 2
		graphic.DisplayNiche()
		break
	case 5:
		values.CurrentOption = 0
		values.CurrentPage = "who"
		values.CurrentOptionMax = 2
		graphic.DisplayWho()
		break
	case 6:
		pressF11()
		utils.ClearTerminal()
		os.Exit(0)
		break
	}
}

func InventoryExec() {
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

	case 9:

	case 10:

	case 11:
		values.CurrentOption = 0
		values.CurrentPage = "real_main_menu"
		values.CurrentOptionMax = 4
		graphic.RefreshRealMainMenu()
		break
	}
}

func WhoExec() {
	switch values.CurrentOption {
	case 0:

	case 1:

	case 2:
		values.CurrentOption = 0
		values.CurrentPage = "real_main_menu"
		values.CurrentOptionMax = 4
		graphic.RefreshRealMainMenu()
		break
	}
}
