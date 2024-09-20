package structs

import "math/rand"

type Dog struct {
	Name            string  `json:"name"`
	Damage          int     `json:"damage"`
	MaxDamage       int     `json:"max_damage"`
	Health          int     `json:"health"`
	MaxHealth       int     `json:"max_health"`
	DodgeChance     float32 `json:"dodge_chance"`
	DodgeMultiplier int     `json:"dodge_multiplier"`
	Race            string  `json:"race"`
	Experience      int     `json:"experience"`
}

func (d *Dog) SetName(name string) {
	d.Name = name
}

func (d *Dog) SetDamage(damage int) {
	d.Damage = damage
}

func (d *Dog) SetMaxDamage(maxDamage int) {
	d.MaxDamage = maxDamage
}

func (d *Dog) SetHealth(health int) {
	d.Health = health
	if d.Health > d.MaxHealth {
		d.Health = d.MaxHealth
	}
}

func (d *Dog) SetMaxHealth(maxHealth int) {
	d.MaxHealth = maxHealth
}

func (d *Dog) SetDodgeChance(dodgeChance float32) {
	d.DodgeChance = dodgeChance
}

func (d *Dog) IsAlive() bool {
	return d.Health > 0
}

func (d *Dog) TakeDamage(amount int) {
	if amount > d.Health {
		d.Health = 0
	} else {
		d.Health -= amount
	}
}

func (d *Dog) Heal(amount int) {
	if d.Health+amount > d.MaxHealth {
		d.Health = d.MaxHealth
	} else {
		d.Health += amount
	}
}

func (d *Dog) Dodge() bool {
	return rand.Float32() < d.DodgeChance
}

func (d *Dog) Attack() int {
	return rand.Intn(d.MaxDamage-d.Damage+1) + d.Damage
}

func (d *Dog) Critical() bool {
	return rand.Intn(100) < 10
}

func (d *Dog) GetName() string {
	return d.Name
}

func (d *Dog) GetDamage() int {
	return d.Damage
}

func (d *Dog) GetMaxDamage() int {
	return d.MaxDamage
}

func (d *Dog) GetHealth() int {
	return d.Health
}

func (d *Dog) GetMaxHealth() int {
	return d.MaxHealth
}

func (d *Dog) GetDodgeChance() float32 {
	return d.DodgeChance
}

func NewDog(name string, damage, maxDamage, health, maxHealth int, dodgeChance float32, race string) Dog {
	return Dog{
		Name:        name,
		Damage:      damage,
		MaxDamage:   maxDamage,
		Health:      health,
		MaxHealth:   maxHealth,
		DodgeChance: dodgeChance,
		Race:        race,
	}
}
