extends KinematicBody

# Declare member variables here. Examples:
# var a = 2
# var b = "text"

# Called when the node enters the scene tree for the first time.
func _ready():
	pass # Replace with function body.

func _process(delta):
	var velocity = Vector3()
	if(Input.is_action_pressed("ui_up")):
		velocity.x += 1
	
	move_and_slide(velocity, Vector3(0,1,0))
