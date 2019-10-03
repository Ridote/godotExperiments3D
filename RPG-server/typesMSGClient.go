package main

// - - - -
// CLIENT
// - - - -

type mcLogin struct {
	Username string `msg:"username"`
	Password string `msg:"password"`
}

type mcRegister struct {
	Username string `msg:"username"`
	Password string `msg:"password"`
	Email    string `msg:"email"`
}

type mcNewPlayer struct {
	Name string `msg:"name"`
}

type mcPos struct {
	ID, PX, PY, PZ, RX, RY int
}

type mcItem struct {
	ID     int
	ItemID int `msg:"itemID"`
}

type mcEquipItem struct {
	mcItem
}

type mcUnequipItem struct {
	mcItem
}

type mcUseItem struct {
	mcItem
}

type mcDeleteItem struct {
	mcItem
}

type mcTalk struct {
	ID       int
	TargetID int `msg:"targetID"`
}

type mcAnswer struct {
	ID       int
	TargetID int `msg:"targetID"`
	Option   int `msg:"option"`
}
