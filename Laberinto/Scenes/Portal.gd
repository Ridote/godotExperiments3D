extends MeshInstance

onready var portalCam : Camera = $Viewport/Camera

func _ready():
	add_to_group(Constants.G_PORTALS)
	portalCam.set_size(OS.window_size.y)

func updateCamera(main_camera_transform):
	scale.y *= -1
	portalCam.global_transform = main_camera_transform
	scale.y *= -1
	portalCam.global_transform.basis.x *= -1