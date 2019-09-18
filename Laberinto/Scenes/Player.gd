extends KinematicBody

var speed : float = 80
var rotationSpeed : float = 5

var gravityUp : float = 9.8
var gravityDown : float = 14.5

var jumpSpeed : float = 500.0
var velocity : Vector3 = Vector3()

var lerpFactor : float = 0.1

func _ready():
	pass # Replace with function body.

func _process(delta):
	var direction = Vector3()
	
	if(Input.is_action_pressed("ui_left")):
		rotate(Vector3(0,1,0), delta*rotationSpeed)
	if(Input.is_action_pressed("ui_right")):
		rotate(Vector3(0,1,0), -delta*rotationSpeed)
	if(Input.is_action_pressed("ui_up")):
		direction = Vector3(-sin(rotation.y), 0.0, -cos(rotation.y)).normalized()
	if(Input.is_action_pressed("ui_down")):
		direction = Vector3(sin(rotation.y), 0.0, cos(rotation.y)).normalized()	
	direction *= speed * delta
	if(!is_on_floor()):
		if(velocity.y > 0):
			velocity.y -= gravityUp * delta
		else:
			velocity.y -= gravityDown * delta
	else:
		velocity.y = 0.0
		if (Input.is_action_pressed("ui_jump")):
			velocity.y = delta * jumpSpeed
		
	velocity += direction
	velocity = move_and_slide(velocity, Vector3(0.0, 1.0, 0.0))
	velocity.x = lerp(velocity.x,  0, lerpFactor)
	velocity.z = lerp(velocity.z, 0, lerpFactor)
