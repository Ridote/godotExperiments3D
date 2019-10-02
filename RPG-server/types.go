package main

import (
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	"github.com/tinylib/msgp/msgp"
)

//go:generate msgp

// User is an user
type User struct {
	gorm.Model
	Username string          `msg:"username"`
	ws       *websocket.Conn `msg:"-"`
}

// Auth is the response for the client after a login attempt
type Auth struct {
	Success bool   `msg:"success"`
	Msg     string `msg:"msg"`
}

// MSG is the envelope of any MessagePacket
type MSG struct {
	Kind    string   `msg:"type"`
	Payload msgp.Raw `msg:"payload"`
}

// Login is a login attempt from a client
type Login struct {
	Username string `msg:"username"`
	Password string `msg:"password"`
}
