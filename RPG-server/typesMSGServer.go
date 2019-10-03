package main

import (
	"github.com/tinylib/msgp/msgp"
)

// MSG is the envelope of any MessagePacket
type MSG struct {
	Kind    string   `msg:"type"`
	Payload msgp.Raw `msg:"payload"`
}

// - - - -
// SERVER
// - - - -

type msAuth struct {
	Success bool   `msg:"success"`
	Msg     string `msg:"msg"`
}

type msNewPlayer struct {
	ID      int
	Name    string `msg:"name"`
	Model3D int    `msg:"model3D"`
	Owner   bool   `msg:"owner"`
}

type msPlayerStats struct {
	ID, STR, AGI, INT int
}

type msPlayerState struct {
	ID, HP, MP int
}

type msPos struct {
	ID, PX, PY, PZ, RX, RZ int
}

type msInventory struct {
	ID    int
	Items []dbItem `msg:"items"`
}

type msItem struct {
	ID     int
	ItemID int `msg:"itemID"`
}

type msUseItem struct {
	msItem
}

type msNewItem struct {
	msItem
}

type msDeleteItem struct {
	msItem
}

type msEquipItem struct {
	msItem
}

type msUnequipItem struct {
	msItem
}

type msQuestion struct {
	ID       int
	TargetID int
	Text     string
	Options  []string
}

type msTalk struct {
	ID       int
	TargetID int
	Text     string
}
