package microlog

import (
	log "github.com/sirupsen/logrus"
	"io"
	"time"
)

var (
	// std is the name of the standard logger in stdlib `log`
	mic = NewMicrolog()
)

type Micro struct {
	Std    *log.Logger
	Fields log.Fields
}

func NewMicrolog() *Micro {
	return &Micro{
		Std: log.New(),
	}
}

func StandardLogger() *log.Logger {
	return mic.Std
}

func SetFields(Fields log.Fields) {
	mic.Fields = Fields
}

// SetOutput sets the standard logger output.
func SetOutput(out io.Writer) {
	mic.Std.SetOutput(out)
}

// SetFormatter sets the standard logger formatter.
func SetFormatter(formatter log.Formatter) {
	mic.Std.SetFormatter(formatter)
}

// SetReportCaller sets whether the standard logger will include the calling
// method as a field.
func SetReportCaller(include bool) {
	mic.Std.SetReportCaller(include)
}

// SetLevel sets the standard logger level.
func SetLevel(level log.Level) {
	mic.Std.SetLevel(level)
}

// GetLevel returns the standard logger level.
func GetLevel() log.Level {
	return mic.Std.GetLevel()
}

// IsLevelEnabled checks if the log level of the standard logger is greater than the level paramvelEnabled checks if the log level of the standard logger is greater than the level param
func IsLevelEnabled(level log.Level) bool {
	return mic.Std.IsLevelEnabled(level)
}

// AddHook adds a hook to the standard logger hooks.
func AddHook(hook Hook) {
	mic.Std.AddHook(hook)
}

func AddHookOrigin(hook log.Hook) {
	mic.Std.AddHook(hook)
}

// WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.
func WithError(err error) *log.Entry {
	return mic.Std.WithField(log.ErrorKey, err)
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *log.Entry {
	return mic.Std.WithField(key, value)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields log.Fields) *log.Entry {
	return mic.Std.WithFields(fields)
}

// WithTime creats an entry from the standard logger and overrides the time of
// logs generated with it.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithTime(t time.Time) *log.Entry {
	return mic.Std.WithTime(t)
}

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Trace(args...)
	} else {
		mic.Std.Trace(args...)
	}
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Debug(args...)
	} else {
		mic.Std.Debug(args...)
	}
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Print(args...)
	} else {
		mic.Std.Print(args...)
	}
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Info(args...)
	} else {
		mic.Std.Info(args...)
	}
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Warn(args...)
	} else {
		mic.Std.Warn(args...)
	}
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Warning(args...)
	} else {
		mic.Std.Warning(args...)
	}
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Error(args...)
	} else {
		mic.Std.Error(args...)
	}
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Panic(args...)
	} else {
		mic.Std.Panic(args...)
	}
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Fatal(args...)
	} else {
		mic.Std.Fatal(args...)
	}
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(format string, args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Tracef(format, args...)
	} else {
		mic.Std.Tracef(format, args...)
	}
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Debugf(format, args...)
	} else {
		mic.Std.Debugf(format, args...)
	}
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Printf(format, args...)
	} else {
		mic.Std.Printf(format, args...)
	}
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Infof(format, args...)
	} else {
		mic.Std.Infof(format, args...)
	}
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Warnf(format, args...)
	} else {
		mic.Std.Warnf(format, args...)
	}
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Warningf(format, args...)
	} else {
		mic.Std.Warningf(format, args...)
	}
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Errorf(format, args...)
	} else {
		mic.Std.Errorf(format, args...)
	}
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Panicf(format, args...)
	} else {
		mic.Std.Panicf(format, args...)
	}
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Fatalf(format, args...)
	} else {
		mic.Std.Fatalf(format, args...)
	}
}

// Traceln logs a message at level Trace on the standard logger.
func Traceln(args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Traceln(args...)
	} else {
		mic.Std.Traceln(args...)
	}
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Debugln(args...)
	} else {
		mic.Std.Debugln(args...)
	}
}

// Println logs a message at level Info on the standard logger.
func Println(args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Println(args...)
	} else {
		mic.Std.Println(args...)
	}
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Infoln(args...)
	} else {
		mic.Std.Infoln(args...)
	}
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Warnln(args...)
	} else {
		mic.Std.Warnln(args...)
	}
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Warningln(args...)
	} else {
		mic.Std.Warningln(args...)
	}
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Errorln(args...)
	} else {
		mic.Std.Errorln(args...)
	}
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Panicln(args...)
	} else {
		mic.Std.Panicln(args...)
	}
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalln(args ...interface{}) {
	if mic.Fields != nil {
		mic.Std.WithFields(mic.Fields).Fatalln(args...)
	} else {
		mic.Std.Fatalln(args...)
	}
}

type Hook struct {
	writer    io.Writer
	formatter log.Formatter
}

func (h Hook) Fire(e *log.Entry) error {
	dataBytes, err := h.formatter.Format(e)
	if err != nil {
		return err
	}
	_, err = h.writer.Write(dataBytes)
	return err
}

func (h Hook) Levels() []log.Level {
	return log.AllLevels
}
