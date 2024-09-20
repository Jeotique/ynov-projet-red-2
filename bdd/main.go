package bdd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"projet/structs"
	"projet/values"
	"strings"
)

var Database *QuickDB

type QuickDB struct {
	filePath string
	data     map[string]interface{}
}

func NewQuickDB(filePath string) *QuickDB {
	db := &QuickDB{
		filePath: filePath,
		data:     make(map[string]interface{}),
	}
	db.loadData()
	return db
}

func (db *QuickDB) loadData() {
	file, err := ioutil.ReadFile(db.filePath)
	if err != nil {
		return
	}

	if err := json.Unmarshal(file, &db.data); err != nil {
		fmt.Println("Error loading data:", err)
	}
}

func (db *QuickDB) saveData() {
	dataJSON, err := json.MarshalIndent(db.data, "", "  ")
	if err != nil {
		fmt.Println("Error saving data:", err)
		return
	}

	err = ioutil.WriteFile(db.filePath, dataJSON, 0644)
	if err != nil {
		fmt.Println("Error saving data to file:", err)
	}
}

func (db *QuickDB) Set(key string, value interface{}) {
	db.data[key] = value
	db.saveData()
}

func (db *QuickDB) Get(key string) interface{} {
	return db.data[key]
}

func (db *QuickDB) Delete(key string) {
	delete(db.data, key)
	db.saveData()
}

func (db *QuickDB) GetAll() map[string]interface{} {
	return db.data
}

func (db *QuickDB) Find(query string) map[string]interface{} {
	result := make(map[string]interface{})
	for key, value := range db.data {
		if strings.Contains(key, query) {
			result[key] = value
		}
	}
	return result
}

func (db *QuickDB) LoadPlayer() error {
	jsonData, err := ioutil.ReadFile("data.json")
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	var player structs.Character

	err = json.Unmarshal(jsonData, &player)
	if err != nil {
		return fmt.Errorf("error unmarshaling JSON: %v", err)
	}

	// Debugging: print the player struct to check if it was loaded correctly
	fmt.Printf("Loaded player: %+v\n", player)

	// Assuming values.Player is a global variable, make sure it's accessible
	values.Player = player

	return nil
}

func (db *QuickDB) SavePlayer(player structs.Character) error {
	playerJSON, err := json.Marshal(player)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("data.json", playerJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}
