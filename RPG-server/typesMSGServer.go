package main

import (
	"github.com/tinylib/msgp/msgp"
)

// MSG is the envelope of any MessagePacket
type MSG struct {
	Kind    string   `msg:"type"`
	Payload msgp.Raw `msg:"payload"`
}

//go:generate msgp

// - - - -
// SERVER
// - - - -

type MSAuth struct {
	Success bool   `msg:"success"`
	Msg     string `msg:"msg"`
}

type MSNewPlayer struct {
	ID      int
	Name    string `msg:"name"`
	Model3D int    `msg:"model3D"`
	Owner   bool   `msg:"owner"`
}

type MSPlayerStats struct {
	ID, STR, AGI, INT int
}

type MSPlayerState struct {
	ID, HP, MP int
}

type MSPos struct {
	ID, PX, PY, PZ, RX, RZ int
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
	TargetID int
	Text     string
	Options  []string
}

type MSTalk struct {
	ID       int
	TargetID int
	Text     string
}
