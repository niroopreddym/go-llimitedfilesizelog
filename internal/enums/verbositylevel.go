package enums

//VerbosityLevel is the enum for the different error verbosities filtering on the log part
type VerbosityLevel int

const (
	//Info is the least verbose level
	Info VerbosityLevel = iota
	//Warning is the the second verbose level
	Warning
	//Error verbose level logs evrything above it
	Error
	//Fatal verbose level is the master level
	Fatal
)
