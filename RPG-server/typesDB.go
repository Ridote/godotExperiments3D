package main

import (
	"github.com/gorilla/websocket"
	"github.com/tinylib/msgp/msgp"
)

//go:generate msgp

// ItemCategory represents the type of item (equipable, usable, etc)
type DBItemCategory uint8

const (
	itemConsumable DBItemCategory = iota
	itemEquipable
	itemQuest
	itemMiscellaneous
)

type DBItem struct {
	ID       uint `gorm:"primary_key"`
	Category DBItemCategory
}
type DBInventory struct {
	ID    uint `gorm:"primary_key"`
	Space int
	Items []DBItem
}

type DBUser struct {
	ID        uint        `gorm:"primary_key"`
	Username  string      `msg:"username"`
	Password  string      `msg:"password"`
	Inventory DBInventory `msg:"inventory"`
	ws        *websocket.Conn
	Games     []DBGame `msg:"games" gorm:"many2many:user_games;"`
}

type DBCharacter struct {
	ID             uint `gorm:"primary_key"`
	Owner          uint `msg:"owner"`
	Game           uint
	PX, PY, PZ, RY int
	Look           int    `msg:"look"`
	Name           string `msg:"name"`
	Class          string `msg:"class"`
	HP             int    `msg:"HP"`
	MP             int    `msg:"MP"`
	CurrentHP      int    `msg:"currentHP"`
	CurrentMP      int    `msg:"currentMP"`
	STR            int    `msg:"STR"`
	AGI            int    `msg:"AGI"`
	INT            int    `msg:"INT"`
	Map            int    `msg:"map"`
}

type DBGame struct {
	ID         uint          `gorm:"primary_key"`
	Time       int           `msg:"time"`
	Users      []DBUser      `msg:"users" gorm:"many2many:user_games;"`
	Characters []DBCharacter `msg:"characters" gorm:"foreignkey:Game"`
	Lock       bool          `msg:"lock"` // defines if all users must be connected to be able to play this game
}

// MSG is the envelope of any MessagePacket
type MSG struct {
	Kind    string   `msg:"type"`
	Payload msgp.Raw `msg:"payload"`
}

// - - - -
// SERVER
// - - - -

type MSAuth struct {
	Success bool   `msg:"success"`
	Msg     string `msg:"msg"`
}

type MSNewCharacter struct {
	ID      int
	Name    string `msg:"name"`
	Model3D int    `msg:"model3D"`
	Owner   bool   `msg:"owner"`
}

type MSCharacterStats struct {
	ID, STR, AGI, INT int
}

type MSCharacterState struct {
	ID, HP, MP int
}

type MSPos struct {
	ID, PX, PY, PZ, RY int
}

type MSInventory struct {
	ID    int
	Items []DBItem `msg:"items"`
}

type MSItem struct {
	ID     int
	ItemID int `msg:"itemID"`
}

type MSUseItem struct {
	MSItem
}

type MSNewItem struct {
	MSItem
}

type MSDeleteItem struct {
	MSItem
}

type MSEquipItem struct {
	MSItem
}

type MSUnequipItem struct {
	MSItem
}

type MSQuestion struct {
	ID       int
	TargetID int      `msg:"targetID"`
	Text     string   `msg:"text"`
	Options  []string `msg:"options"`
}

type MSTalk struct {
	ID       int
	TargetID int    `msg:"targetID"`
	Text     string `msg:"text"`
}

// - - - -
// CLIENT
// - - - -

type MCLogin struct {
	Username string `msg:"username"`
	Password string `msg:"password"`
}

type MCRegister struct {
	Username string `msg:"username"`
	Password string `msg:"password"`
	Email    string `msg:"email"`
}

type MCNewCharacter struct {
	Name string `msg:"name"`
}

type MCPos struct {
	ID, PX, PY, PZ, RY int
}

type MCItem struct {
	ID     int
	ItemID int `msg:"itemID"`
}

type MCEquipItem struct {
	MCItem
}

type MCUnequipItem struct {
	MCItem
}

type MCUseItem struct {
	MCItem
}

type MCDeleteItem struct {
	MCItem
}

type MCTalk struct {
	ID       int
	TargetID int `msg:"targetID"`
}

type MCAnswer struct {
	ID       int
	TargetID int `msg:"targetID"`
	Option   int `msg:"option"`
}

type MCNewGame struct {
	Lock bool `msg:"lock"`
}
