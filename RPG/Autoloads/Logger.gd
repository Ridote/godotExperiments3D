extends Node

const SEVERITY_LEVELS = {
	ERROR = 0,
	WARNING = 1,
	INFO = 2
}

const SEVERITY = SEVERITY_LEVELS.INFO

func logger(msg, entity = "", type=SEVERITY_LEVELS.WARNING):
	var severityString = ""
	match type:
		SEVERITY_LEVELS.ERROR: severityString = "ERROR"
		SEVERITY_LEVELS.WARNING: severityString = "WARNING"
		SEVERITY_LEVELS.INFO: severityString = "INFO"
	msg = str(msg)
	if(SEVERITY >= type):
		if entity:
			msg = severityString + "(" + entity + "): " + msg
		else:
			msg = severityString + ": " + msg
		print(msg)

func info(msg, entity = ""):
	logger(msg, entity, SEVERITY_LEVELS.INFO)
func warning(msg, entity = ""):
	logger(msg, entity, SEVERITY_LEVELS.WARNING)
func error(msg, entity = ""):
	logger(msg, entity, SEVERITY_LEVELS.ERROR)