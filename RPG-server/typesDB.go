package main

import (
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
)

//go:generate msgp
//msgp:ignore User

// User is an user
type dbUser struct {
	gorm.Model
	Username string
	Password string // hashed
	ws       *websocket.Conn
}

// ItemCategory represents the type of item (equipable, usable, etc)
type ItemCategory uint8

const (
	itemConsumable ItemCategory = iota
	itemEquipable
	itemQuest
	itemMiscellaneous
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
	ID                 int
	Owner              dbUser
	PX, PY, PZ, RX, RZ int
	Name               string      `msg:"name"`
	Class              string      `msg:"class"`
	HP                 int         `msg:"HP"`
	MP                 int         `msg:"MP"`
	CurrentHP          int         `msg:"currentHP"`
	CurrentMP          int         `msg:"currentMP"`
	STR                int         `msg:"STR"`
	AGI                int         `msg:"AGI"`
	INT                int         `msg:"INT"`
	Inventory          dbInventory `msg:"inventory"`
}
