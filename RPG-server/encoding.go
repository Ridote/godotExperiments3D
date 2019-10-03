package main

import "fmt"

func encodeMsg(kind string, payload interface{}) []byte {
	switch kind {
	case "auth":
		r := payload.(mAuth)
		p, _ := r.MarshalMsg(nil)
		m := MSG{
			Kind:    kind,
			Payload: p,
		}
		encM, _ := m.MarshalMsg(nil)
		return encM
	}

	fmt.Println("Unknown kind:", kind)
	return nil
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
		var r mLogin
		_, err := r.UnmarshalMsg([]byte(aux))
		if err != nil {
			fmt.Println("case Login unmarshall Error:", err)
		}
		return r

	default:
		fmt.Println("Unknown kind:", msg.Kind)
	}
	return nil
}
