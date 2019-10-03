package main

import (
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
)

//go:generate msgp

// User is an user
type User struct {
	gorm.Model
	Username string
	ws       *websocket.Conn
}

// ItemCategory represents the type of item (equipable, usable, etc)
type ItemCategory int

const (
	// Consumable an item that can be consumed
	Consumable ItemCategory = iota
	// Equipable an item that can be equiped
	Equipable
	// Quest an item used in quests
	Quest
	// Miscellaneous the rest
	Miscellaneous
)

type dbItem struct {
	gorm.Model
	ID       int
	Category ItemCategory
}
type dbInventory struct {
	Space int
	Items []dbItem
}

// Player is an unique Player controled by an User
type dbPlayer struct {
	gorm.Model
	ID        int         `msg:"ID"`
	Name      string      `msg:"name"`
	Class     string      `msg:"class"`
	HP        int         `msg:"HP"`
	MP        int         `msg:"MP"`
	CurrentHP int         `msg:"currentHP"`
	CurrentMP int         `msg:"currentMP"`
	STR       int         `msg:"STR"`
	AGI       int         `msg:"AGI"`
	INT       int         `msg:"INT"`
	Inventory dbInventory `msg:"inventory"`
}
