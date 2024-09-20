package structs

type Inventory struct {
	Heal        int  `json:"heal"`
	Poison      int  `json:"poison"`
	Kalashnikov bool `json:"kalashnikov"`
	Fireball    int  `json:"fireball"`
	Iceball     int  `json:"iceball"`
	Bomb        int  `json:"bomb"`
	Shield      int  `json:"shield"`
	Heineken    int  `json:"heineken"`
	Wood        int  `json:"wood"`
	Viscous     int  `json:"viscous"`
	Stone       int  `json:"stone"`
	Money       int  `json:"money"`
}

func AddHeal(i *Inventory) {
	i.Heal++
}

func AddPoison(i *Inventory) {
	i.Poison++
}

func RemoveHeal(i *Inventory) {
	i.Heal--
}

func RemovePoison(i *Inventory) {
	i.Poison--
}

func AddFireball(i *Inventory) {
	i.Fireball++
}

func RemoveFireball(i *Inventory) {
	i.Fireball--
}

func AddIceball(i *Inventory) {
	i.Iceball++
}

func RemoveIceball(i *Inventory) {
	i.Fireball--
}

func AddBomb(i *Inventory) {
	i.Bomb++
}

func RemoveBomb(i *Inventory) {
	i.Bomb--
}

func AddShield(i *Inventory) {
	i.Shield++
}

func RemoveShield(i *Inventory) {
	i.Shield--
}

func AddHeineken(i *Inventory) {
	i.Heineken++
}

func RemoveHeineken(i *Inventory) {
	i.Heineken--
}

func AddMoney(i *Inventory, m int) {
	i.Money += m
}

func RemoveMoney(i *Inventory, m int) {
	i.Money -= m
}

func EnableKalashnikov(i *Inventory) {
	i.Kalashnikov = true
}

func DisableKalashnikov(i *Inventory) {
	i.Kalashnikov = false
}

func GetHeal(i *Inventory) int {
	return i.Heal
}

func GetPoison(i *Inventory) int {
	return i.Poison
}

func GetKalashnikov(i *Inventory) bool {
	return i.Kalashnikov
}

func GetFireball(i *Inventory) int {
	return i.Fireball
}

func GetIceball(i *Inventory) int {
	return i.Iceball
}

func GetBomb(i *Inventory) int {
	return i.Bomb
}

func GetShield(i *Inventory) int {
	return i.Shield
}

func GetHeineken(i *Inventory) int {
	return i.Heineken
}

func GetMoney(i *Inventory) int {
	return i.Money
}

func GetDefaultInventory() Inventory {
	return Inventory{
		Heal:        0,
		Poison:      0,
		Kalashnikov: false,
		Fireball:    0,
		Iceball:     0,
		Bomb:        0,
		Shield:      0,
		Heineken:    0,
		Wood:        0,
		Viscous:     0,
		Stone:       0,
		Money:       100,
	}
}

func GetInventory(i *Inventory) Inventory {
	return *i
}

func SetInventory(i *Inventory, inv Inventory) {
	i.Heal = inv.Heal
	i.Poison = inv.Poison
	i.Kalashnikov = inv.Kalashnikov
	i.Fireball = inv.Fireball
	i.Iceball = inv.Iceball
	i.Bomb = inv.Bomb
	i.Shield = inv.Shield
	i.Heineken = inv.Heineken
	i.Wood = inv.Wood
	i.Viscous = inv.Viscous
	i.Stone = inv.Stone
	i.Money = inv.Money
}

func GetWood(i *Inventory) int {
	return i.Wood
}

func GetViscous(i *Inventory) int {
	return i.Viscous
}

func GetStone(i *Inventory) int {
	return i.Stone
}

func AddWood(i *Inventory) {
	i.Wood++
}

func AddViscous(i *Inventory) {
	i.Viscous++
}

func AddStone(i *Inventory) {
	i.Stone++
}

func SetWood(i *Inventory, w int) {
	i.Wood = w
}

func SetViscous(i *Inventory, v int) {
	i.Viscous = v
}

func SetStone(i *Inventory, s int) {
	i.Stone = s
}
