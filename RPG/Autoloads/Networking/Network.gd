extends Node

const SERVER_RESPONSE_FACTORY = preload("res://Autoloads/Networking/ServerResponse.gd")
const CLIENT_REQUEST_FACTORY = preload("res://Autoloads/Networking/ClientRequest.gd")
var serverResponse = null
var clientRequest = null

var ws = null

func _ready():
	ws = WebSocketClient.new()
	ws.set_verify_ssl_enabled(false)
	
	ws.connect("connection_established", self, "_connection_established")
	ws.connect("connection_closed", self, "_connection_closed")
	ws.connect("connection_error", self, "_connection_error")
	
	#var url = "ws://pc.galax.be:1616"
	var url = "ws://localhost:1616"
	# var url = "ws://echo.websocket.org"
	Logger.info("Connecting to " + url, Groups.get.NETWORKING)
	ws.connect_to_url(url)
	
	
func _connection_established(protocol):
	Logger.info("Connection established with protocol: " + protocol, Groups.get.NETWORKING)
	serverResponse = SERVER_RESPONSE_FACTORY.new()
	clientRequest = CLIENT_REQUEST_FACTORY.new()
	clientRequest.initRequester(ws)
	
	clientRequest.login()
	
func _connection_closed(arg):
	Logger.info("Connection closed" + arg, Groups.get.NETWORKING)

func _connection_error():
	Logger.error("Connection error", Groups.get.NETWORKING)

func _process(delta):
	delta = 1
	if ws.get_connection_status() == ws.CONNECTION_CONNECTING || ws.get_connection_status() == ws.CONNECTION_CONNECTED:
		ws.poll()

	if ws.get_peer(1).is_connected_to_host():
		if ws.get_peer(1).get_available_packet_count() > 0 :
			receivedPacket(msgpack.decode(ws.get_peer(1).get_packet()).result)

func receivedPacket(packet):
	serverResponse.processResponse(packet)