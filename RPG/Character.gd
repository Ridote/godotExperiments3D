extends Spatial


class_name Character

func initialize(_str, _agi, _int) -> void:
	$Stats.initialize(_str, _agi, _int)

func attack(enemy : Character, duration : float = 2.0):
	var tween : Tween = Tween.new()
	var currentPosition : Vector3 = translation
	# We need to calculate the offset between global positions and apply that to a local one since tween doesn't seem to manage well global positions through transform.origin
	var targetPosition : Vector3 = translation + (enemy.getFrontPosition() - transform.origin)
	get_tree().get_root().call_deferred("add_child", tween)
# warning-ignore:return_value_discarded
	tween.interpolate_property(self, 'translation', currentPosition, targetPosition, duration,Tween.TRANS_LINEAR,Tween.EASE_IN)
# warning-ignore:return_value_discarded
	tween.start()
	yield(tween, "tween_completed")
# warning-ignore:return_value_discarded
	tween.interpolate_property(self, 'translation', targetPosition, currentPosition, duration,Tween.TRANS_LINEAR,Tween.EASE_IN)
# warning-ignore:return_value_discarded
	tween.start()
	yield(tween, "tween_completed")
	print("finished")

func getFrontPosition() -> Vector3:
	return transform.origin + $FrontPosition.translation