extends Node

const REQUEST_TYPE = {
	login = "login",
	signup = "register",
	new_game = "newGame",
	load_games = "loadGames"
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

func login(username : String, password : String):
	sendRequest(REQUEST_TYPE.login, {username=username, password=password})

func signup(username : String, password : String):
	sendRequest(REQUEST_TYPE.signup, {username=username, password=password})

func newGame(lock : bool):
	sendRequest(REQUEST_TYPE.new_game, {lock=lock})

func loadGames():
	sendRequest(REQUEST_TYPE.load_games, null)