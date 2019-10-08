package main

import (
	"fmt"

	"github.com/tinylib/msgp/msgp"
)

func encodeMsg(kind string, payload interface{}) []byte {
	var p msgp.Raw
	switch kind {
	case "auth":
		r := payload.(MSAuth)
		p, _ = r.MarshalMsg(nil)
	case "newCharacter":
		r := payload.(MSNewCharacter)
		p, _ = r.MarshalMsg(nil)
	case "characterStats":
		r := payload.(MSCharacterStats)
		p, _ = r.MarshalMsg(nil)
	case "characterState":
		r := payload.(MSCharacterState)
		p, _ = r.MarshalMsg(nil)
	case "pos":
		r := payload.(MSPos)
		p, _ = r.MarshalMsg(nil)
	case "inventory":
		r := payload.(MSInventory)
		p, _ = r.MarshalMsg(nil)
	case "useItem":
		r := payload.(MSUseItem)
		p, _ = r.MarshalMsg(nil)
	case "newItem":
		r := payload.(MSNewItem)
		p, _ = r.MarshalMsg(nil)
	case "deleteItem":
		r := payload.(MSDeleteItem)
		p, _ = r.MarshalMsg(nil)
	case "equipItem":
		r := payload.(MSEquipItem)
		p, _ = r.MarshalMsg(nil)
	case "unequipItem":
		r := payload.(MSUnequipItem)
		p, _ = r.MarshalMsg(nil)
	case "question":
		r := payload.(MSQuestion)
		p, _ = r.MarshalMsg(nil)
	case "talk":
		r := payload.(MSTalk)
		p, _ = r.MarshalMsg(nil)
	case "game":
		r := payload.(DBGame)
		p, _ = r.MarshalMsg(nil)
	}
	m := MSG{
		Kind:    kind,
		Payload: p,
	}
	encM, _ := m.MarshalMsg(nil)

	return encM
}

func getMsgKind(m []byte) string {
	var msg MSG
	msg.UnmarshalMsg(m)
	return msg.Kind
}

func decodeMsg(m []byte) interface{} {
	var msg MSG
	_, err := msg.UnmarshalMsg(m)
	if err != nil {
		fmt.Println("msgKind:", err)
	}
	aux := msg.Payload[2:len(msg.Payload)]

	switch msg.Kind {
	case "login":
		var r MCLogin
		_, err := r.UnmarshalMsg([]byte(aux))
		if err != nil {
			fmt.Println("case Login unmarshall Error:", err)
		}
		return r
	case "register":
		var r MCRegister
		_, err := r.UnmarshalMsg([]byte(aux))
		if err != nil {
			fmt.Println("case register unmarshall Error:", err)
		}
		return r
	case "newCharacter":
		var r MCNewCharacter
		_, err := r.UnmarshalMsg([]byte(aux))
		if err != nil {
			fmt.Println("case newcharacter unmarshall Error:", err)
		}
		return r
	case "pos":
		var r MCPos
		_, err := r.UnmarshalMsg([]byte(aux))
		if err != nil {
			fmt.Println("case pos unmarshall Error:", err)
		}
		return r
	case "equipItem":
		var r MCEquipItem
		_, err := r.UnmarshalMsg([]byte(aux))
		if err != nil {
			fmt.Println("case equipitem unmarshall Error:", err)
		}
		return r
	case "unequipItem":
		var r MCUnequipItem
		_, err := r.UnmarshalMsg([]byte(aux))
		if err != nil {
			fmt.Println("case unequipitem unmarshall Error:", err)
		}
		return r
	case "useItem":
		var r MCUseItem
		_, err := r.UnmarshalMsg([]byte(aux))
		if err != nil {
			fmt.Println("case useitem unmarshall Error:", err)
		}
		return r
	case "deleteItem":
		var r MCDeleteItem
		_, err := r.UnmarshalMsg([]byte(aux))
		if err != nil {
			fmt.Println("case deleteitem unmarshall Error:", err)
		}
		return r
	case "talk":
		var r MCTalk
		_, err := r.UnmarshalMsg([]byte(aux))
		if err != nil {
			fmt.Println("case talk unmarshall Error:", err)
		}
		return r
	case "answer":
		var r MCAnswer
		_, err := r.UnmarshalMsg([]byte(aux))
		if err != nil {
			fmt.Println("case answer unmarshall Error:", err)
		}
		return r
	case "newGame":
		var r MCNewGame
		_, err := r.UnmarshalMsg([]byte(aux))
		if err != nil {
			fmt.Println("case answer unmarshall Error:", err)
		}
		return r
	default:
		fmt.Println("Unknown kind:", msg.Kind)
	}
	return nil
}
