package logger

// LogLevel define levels of logging
type LogLevel int

const (
	INFO LogLevel = iota
	WARNING
	ERROR
	DEBUG
)

func (l LogLevel) String() string {
	return [...]string{"INFO", "WARNING", "ERROR", "DEBUG"}[l]
}
