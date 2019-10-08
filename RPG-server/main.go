package main

import (
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var upgrader = websocket.Upgrader{} // use default options
var db *gorm.DB

func main() {
	db = setupDB()
	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe("0.0.0.0:1616", nil))
}

func handle(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	var u DBUser

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		fmt.Println("New msg", string(message))
		if mt == websocket.BinaryMessage {
			if u.Username != "" {
				fmt.Println("User exists, handleMsg")
				msgHandler(message, &u, c)
			} else {
				// LOGIN or REGISTER
				l, ok := decodeMsg(message).(MCLogin)
				if !ok {
					// REGISTER
					reg, ok := decodeMsg(message).(MCRegister)
					if !ok {
						fmt.Println("invalid msg sent, when not authenticated you can only send 'login' command")
						c.WriteMessage(websocket.BinaryMessage, encodeMsg("auth", MSAuth{
							Msg:     "invalid msg sent, when not authenticated you can only send 'login' command",
							Success: false,
						}))
						continue
					}
					db.Where("username = ?", reg.Username).First(&u)
					if u.Username != "" {
						fmt.Println("user already exists", u.Username)
						c.WriteMessage(websocket.BinaryMessage, encodeMsg("auth", MSAuth{
							Msg:     "user already exists " + u.Username,
							Success: false,
						}))
						continue
					}
					u.Username = reg.Username
					u.Password = hex.EncodeToString([]byte(reg.Password))

					db.Create(&u)

					fmt.Println("user created", u.Username)
					c.WriteMessage(websocket.BinaryMessage, encodeMsg("auth", MSAuth{
						Msg:     "user created " + u.Username,
						Success: true,
					}))
					continue
				}
				// LOGIN
				db.Where("username = ?", l.Username).First(&u)
				if hex.EncodeToString([]byte(l.Password)) == u.Password {
					c.WriteMessage(websocket.BinaryMessage, encodeMsg("auth", MSAuth{
						Msg:     "authentication succesful!",
						Success: true,
					}))
				} else {
					c.WriteMessage(websocket.BinaryMessage, encodeMsg("auth", MSAuth{
						Msg:     "invalid user!",
						Success: false,
					}))
					u.Username = ""
				}
			}
		} else {
			fmt.Println("received text message, not binary")
			break
		}
	}
}

func msgHandler(msg []byte, u *DBUser, c *websocket.Conn) {
	kind := getMsgKind(msg)
	switch kind {
	case "newCharacter":
		m := decodeMsg(msg).(MCNewCharacter)
		var p DBCharacter
		p.Name = m.Name
		p.Owner = u.ID
		p.HP = 100
		p.CurrentHP = 100
		db.Create(&p)
	case "getCharacters":
		var characters []DBCharacter
		db.Find(&characters).Related(u)
		for _, p := range characters {
			c.WriteMessage(websocket.BinaryMessage, encodeMsg("character", p))
		}
	case "newGame":
		m := decodeMsg(msg).(MCNewGame)
		var g DBGame
		g.Users = append(g.Users, *u)
		g.Lock = m.Lock

		// Debug only, create character
		var c DBCharacter
		c.Look = 1
		c.Name = "DebugElNoob1"
		c.Owner = u.ID
		var c2 DBCharacter
		c2.Look = 0
		c2.Name = "DebugElNoob2"
		c2.Owner = u.ID
		g.Characters = append(g.Characters, c)
		g.Characters = append(g.Characters, c2)
		fmt.Printf("game created %+v\n", g)
		db.Create(&g)
	case "loadGames":
		var games []DBGame
		db.Model(&u).Related(&games, "Games")
		for _, p := range games {
			fmt.Printf("game %+v\n", p)
			c.WriteMessage(websocket.BinaryMessage, encodeMsg("game", p))
		}
	}
}
