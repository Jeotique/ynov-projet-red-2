package structs

type Character struct {
	Name       string    `json:"name"`
	Level      int       `json:"level"`
	Health     int       `json:"health"`
	Inventory  Inventory `json:"inventory"`
	MaxHealth  int       `json:"max_health"`
	Dogs       []Dog     `json:"dogs"`
	CurrentDog int       `json:"current_dog"`
	House      int       `json:"house"`
}

func (c *Character) SetName(name string) {
	c.Name = name
}

func (c *Character) SetLevel(level int) {
	c.Level = level
}

func (c *Character) SetHealth(health int) {
	c.Health = health
}

func (c *Character) SetMaxHealth(maxHealth int) {
	c.MaxHealth = maxHealth
}

func (c *Character) SetInventory(inventory Inventory) {
	c.Inventory = inventory
}

func (c *Character) SetDogs(dogs []Dog) {
	c.Dogs = dogs
}

func (c *Character) SetCurrentDog(currentDog int) {
	c.CurrentDog = currentDog
}

func (c *Character) GetName() string {
	return c.Name
}

func (c *Character) GetLevel() int {
	return c.Level
}

func (c *Character) GetHealth() int {
	return c.Health
}

func (c *Character) GetMaxHealth() int {
	return c.MaxHealth
}

func (c *Character) GetInventory() Inventory {
	return c.Inventory
}

func (c *Character) GetDogs() []Dog {
	return c.Dogs
}

func (c *Character) GetCurrentDog() Dog {
	return c.Dogs[c.CurrentDog]
}

func (c *Character) GetHouse() int {
	return c.House
}

func (c *Character) SetHouse(house int) {
	c.House = house
}
