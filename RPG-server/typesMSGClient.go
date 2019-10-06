package main

// - - - -
// CLIENT
// - - - -

//go:generate msgp

type MCLogin struct {
	Username string `msg:"username"`
	Password string `msg:"password"`
}

type MCRegister struct {
	Username string `msg:"username"`
	Password string `msg:"password"`
	Email    string `msg:"email"`
}

type MCNewPlayer struct {
	Name string `msg:"name"`
}

type MCPos struct {
	ID, PX, PY, PZ, RX, RY int
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
