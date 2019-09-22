extends KinematicBody

export var rotationUpSpeed = 2
export var rotationSideSpeed = 5
export var speed = 1000

export var camera3rd : NodePath



func _process(delta):
	var camera3rdnode = get_node(camera3rd)
	camera3rdnode.look_at(get_global_transform().origin, Vector3(0,1,0))
	
	
	#var rotation_val = 0.1
	#camera3rdnode.global_transform.basis.x = global_transform.basis.x.linear_interpolate(camera3rdnode.global_transform.basis.x, rotation_val)
	#camera3rdnode.global_transform.basis.y = global_transform.basis.y.linear_interpolate(camera3rdnode.global_transform.basis.y, rotation_val)
	#camera3rdnode.global_transform.basis.z = global_transform.basis.z.linear_interpolate(camera3rdnode.global_transform.basis.z, rotation_val)
	# camera3rdnode.rotation = camera3rdnode.rotation.linear_interpolate(rotation, delta * 4)]
	# camera3rdnode.global_transform.basis = camera3rdnode.global_transform.basis
	
	if(Input.is_action_pressed("ui_up")):
		rotate_object_local(Vector3(1, 0, 0), delta*rotationUpSpeed)
	if(Input.is_action_pressed("ui_down")):
		rotate_object_local(Vector3(1, 0, 0), -delta*rotationUpSpeed)
	if(Input.is_action_pressed("ui_left")):
		rotate_object_local(Vector3(0, 0, 1), -delta*rotationSideSpeed)
	if(Input.is_action_pressed("ui_right")):
		rotate_object_local(Vector3(0, 0, 1), delta*rotationSideSpeed)
		
	if(Input.is_action_just_pressed("ui_jump")):
		
		if($Camera1st.current):
			camera3rdnode.current = true
		elif(camera3rdnode.current):
			$CameraBack.current = true
		else:
			$Camera1st.current = true
	var direction = -(get_global_transform().origin - $FrontDirection.get_global_transform().origin).normalized()
	
	
	var _moved = move_and_slide(direction * delta * speed, Vector3(0,1,0))
	