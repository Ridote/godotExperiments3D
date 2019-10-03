# TODO REMOVE WARNINGS IGNORE
# warning-ignore-all:unused_argument
extends Node

const RESPONSE_TYPE = {
	auth = "auth"
}

func processResponse(packet):
	Logger.info("Packet received(" + packet.type + ")", Groups.get.NETWORKING)
	match packet.type:
		RESPONSE_TYPE.auth:
			auth(packet.payload.success, packet.payload.msg)



# Admin
func auth(status, msg = ""):
	if status:
		Logger.info(msg, "Networking")
	else:
		Logger.error(msg, "Networking")

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



