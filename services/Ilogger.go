package services

//LoggerIface exposes the log package defintions to other modules
type LoggerIface interface {
	Log(message string)
}
