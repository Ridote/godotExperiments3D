package main

import (
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
)

//go:generate msgp

// User is an user
type DBUser struct {
	gorm.Model `msg:"-"`
	Username   string
	Password   string      // hashed
	Inventory  DBInventory `msg:"inventory"`
	ws         *websocket.Conn
}

// ItemCategory represents the type of item (equipable, usable, etc)
type DBItemCategory uint8

const (
	itemConsumable DBItemCategory = iota
	itemEquipable
	itemQuest
	itemMiscellaneous
)

type DBItem struct {
	gorm.Model `msg:"-"`
	Category   DBItemCategory
}
type DBInventory struct {
	gorm.Model `msg:"-"`
	Space      int
	Items      []DBItem
}

// Player is an unique Player controled by an User
type DBPlayer struct {
	gorm.Model         `msg:"-"`
	Owner              DBUser
	PX, PY, PZ, RX, RZ int
	Name               string `msg:"name"`
	Class              string `msg:"class"`
	HP                 int    `msg:"HP"`
	MP                 int    `msg:"MP"`
	CurrentHP          int    `msg:"currentHP"`
	CurrentMP          int    `msg:"currentMP"`
	STR                int    `msg:"STR"`
	AGI                int    `msg:"AGI"`
	INT                int    `msg:"INT"`
}
