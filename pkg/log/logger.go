package log

var Log *Logger

type Logger struct {
}

func NewLogger() *Logger {
	return &Logger{}
}

func (l *Logger) Load() {
	Log = l
}
