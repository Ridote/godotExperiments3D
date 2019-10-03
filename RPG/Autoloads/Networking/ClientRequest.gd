extends Node

const REQUEST_TYPE = {
	login = "login"
}

var ws = null

func initRequester(_ws):
	ws = _ws

func sendRequest(_type, _payload):
	var msg = null
	if _payload:
		var payload = msgpack.encode(_payload).result
		msg = msgpack.encode({type = _type, payload = payload}).result
	else:
		msg = msgpack.encode({type = _type}).result
	
	ws.get_peer(1).put_packet(msg)

func login():
	sendRequest(REQUEST_TYPE.login, {username="Rido", password="123456"})
	