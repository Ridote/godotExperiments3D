package main

import (
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

	var u User

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
				msgHandler(message)
			} else {
				// read
				l, ok := decodeMsg(message).(Login)
				if !ok {
					fmt.Println("invalid msg sent, when not authenticated you can only send 'login' command")
					c.WriteMessage(mt, encodeMsg("auth", Auth{
						Msg:     "invalid msg sent, when not authenticated you can only send 'login' command",
						Success: false,
					}))
					break
				}
				if l.Username == "Rido" {
					u.Username = l.Username
					c.WriteMessage(mt, encodeMsg("auth", Auth{
						Msg:     "authentication succesful!",
						Success: true,
					}))
				} else {
					c.WriteMessage(mt, encodeMsg("auth", Auth{
						Msg:     "invalid user!",
						Success: false,
					}))
				}
			}
		} else {
			fmt.Println("received text message, not binary")
			break
		}
	}
}

func msgHandler(msg []byte) {

}
