package enums

//VerbosityLevel is the enum for the different error verbosities filtering on the log part
type VerbosityLevel int

const (
	//Trace is the least log level
	Trace VerbosityLevel = iota
	//Debug is teh second log level
	Debug
	//Info is the second log level
	Info
	//Warning is the the third verbose level
	Warning
	//Error verbose level logs evrything above it
	Error
	//Fatal verbose level is the master level
	Fatal
)

func (v VerbosityLevel) String() string {
	return [...]string{"Trace", "Debug", "Info", "Warning", "Error", "Fatal"}[v]
}
