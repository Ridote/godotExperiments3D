# warning-ignore-all:unused_class_variable
extends Node
# TODO REMOVE ME
var STATS_FACTORY = preload("res://Scenes/Stats.tscn")

var players = []
func _ready():
	initialisePlayers()
	
func initialisePlayers():
	addPlayer("Rido", 5, 8, 4, 9, 10, 7.5)
	addPlayer("Jairo", 6, 10, 2, 7, 8, 5.4)

func addPlayer(_name : String, _str : int, _agi : int, _int : int, _spd : int, _HP : float, _MP : float):
	var player = STATS_FACTORY.instance()
	player.setName(_name)
	player.initialize(_str, _agi, _int, _spd, _HP, _MP)
	players.push_back(player)