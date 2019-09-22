extends Spatial

export (NodePath) var viewport

# Called when the node enters the scene tree for the first time.
func _ready():
	get_node(viewport).asignViewport($Viewport)

func asignViewport(friendViewport):
	$PortalSprite.texture.set_viewport_path_in_scene(friendViewport.get_path())