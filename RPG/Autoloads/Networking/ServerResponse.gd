# TODO REMOVE WARNINGS IGNORE
# warning-ignore-all:unused_argument
extends Node

const RESPONSE_TYPE = {
	auth = "auth",
	signup = "signup"
}

signal auth_success(success, msg)
signal signup_success(success, msg)

func processResponse(packet):
	Logger.info("Packet received(" + packet.type + ")", Groups.NETWORKING)
	match packet.type:
		RESPONSE_TYPE.auth:
			auth(packet.payload.success, packet.payload.msg)
		RESPONSE_TYPE.signup:
			signup(packet.payload.success, packet.payload.msg)

# Admin
func auth(status, msg = ""):
	if status:
		Logger.info(msg, "Networking")
		emit_signal("auth_success", true, msg)
		# TODO load entry
		get_tree().quit()
	else:
		Logger.error(msg, "Networking")
		emit_signal("auth_success", false, msg)

func signup(status, msg=""):
	if status:
		Logger.info(msg, "Networking")
		emit_signal("signup_success", true)
	else:
		Logger.error(msg, "Networking")
		emit_signal("signup_success", false)

func newPlayer(id, name, model3D):
	pass

# Info
func playerStats(ID, stats):
	pass

func playerState(ID, state):
	pass

func pos(ID, PX, PY, PZ, RX, RZ):
	pass

# Actions
func useItem(ID, itemID):
	pass

func newItem(ID, itemID):
	pass

func deleteItem(ID, itemID):
	pass
	
func equipItem(ID, itemID):
	pass
	
func unequipItem(ID, itemID):
	pass
	
func question(ID, targetID, options):
	pass
	
func talk(id, targetID, text):
	pass



