package microlog

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net"
	"os"
	"testing"
	"time"
)

func TestMicrologNoConfig(t *testing.T) {
	Debug("Test Debug")
	Debugf("Test Debugf %v insert", "0.0")
	Info("Test Info")
	Infof("Test Infof %v insert", "0.0")
	Warning("Test Warning")
	Warningf("Test Warningf %v insert", "0.0")
}

func TestMicrologWithFields(t *testing.T) {
	SetFields(log.Fields{"project": "game-api", "module": "mode-test"})

	Debug("Test Debug")
	Debugf("Test Debugf %v insert", "0.0")
	Info("Test Info")
	Infof("Test Infof %v insert", "0.0")
	Warning("Test Warning")
	Warningf("Test Warningf %v insert", "0.0")
}

func TestMicrologBasic(t *testing.T) {
	logFormatter := &log.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		//TimestampFormat:string("2006-01-02 15:04:05"),
		FieldMap: *&log.FieldMap{
			log.FieldKeyMsg:  "message",
			log.FieldKeyTime: "@timestamp",
		},
	}
	SetOutput(os.Stdout)
	SetLevel(log.DebugLevel)
	SetReportCaller(true)
	SetFormatter(logFormatter)

	Debug("Test Debug")
	Debugf("Test Debugf %v insert", "0.0")
	Info("Test Info")
	Infof("Test Infof %v insert", "0.0")
	Warning("Test Warning")
	Warningf("Test Warningf %v insert", "0.0")
}

func TestMicrologProjectConfig(t *testing.T) {
	logFormatter := &log.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		//TimestampFormat:string("2006-01-02 15:04:05"),
		FieldMap: *&log.FieldMap{
			log.FieldKeyMsg:  "message",
			log.FieldKeyTime: "@timestamp",
		},
	}
	SetOutput(os.Stdout)
	SetLevel(log.DebugLevel)
	SetReportCaller(true)
	SetFormatter(logFormatter)
	SetFields(log.Fields{"project": "game-api", "module": "mode-test"})

	Debug("Test Debug")
	Debugf("Test Debugf %v insert", "0.0")
	Info("Test Info")
	Infof("Test Infof %v insert", "0.0")
	Warning("Test Warning")
	Warningf("Test Warningf %v insert", "0.0")
}

func TestMicrologProjectConfigAndHook(t *testing.T) {
	logFormatter := &log.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		//TimestampFormat:string("2006-01-02 15:04:05"),
		FieldMap: *&log.FieldMap{
			log.FieldKeyMsg:  "message",
			log.FieldKeyTime: "@timestamp",
		},
	}
	SetOutput(os.Stdout)
	SetLevel(log.DebugLevel)
	SetReportCaller(true)
	SetFormatter(logFormatter)
	SetFields(log.Fields{"project": "game-api", "module": "mode-test"})

	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		assert.NotNil(t, err)
	}

	AddHook(Hook{
		writer:    conn,
		formatter: logFormatter,
	})

	Debug("Test Debug")
	Debugf("Test Debugf %v insert", "0.0")
	Info("Test Info")
	Infof("Test Infof %v insert", "0.0")
	Warning("Test Warning")
	Warningf("Test Warningf %v insert", "0.0")
}

type DefaultFieldHook struct {
	GetValue func() string
}

func (h *DefaultFieldHook) Levels() []log.Level {
	return log.AllLevels
}

func (h *DefaultFieldHook) Fire(e *log.Entry) error {
	e.Data["aDefaultField"] = h.GetValue()
	return nil
}

func TestMicrologProjectConfigAndHookMulti(t *testing.T) {
	logFormatter := &log.JSONFormatter{
		TimestampFormat: time.RFC3339Nano,
		//TimestampFormat:string("2006-01-02 15:04:05"),
		FieldMap: *&log.FieldMap{
			log.FieldKeyMsg:  "message",
			log.FieldKeyTime: "@timestamp",
		},
	}
	SetOutput(os.Stdout)
	SetLevel(log.DebugLevel)
	SetReportCaller(true)
	SetFormatter(logFormatter)
	SetFields(log.Fields{"project": "game-api", "module": "mode-test"})

	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		assert.NotNil(t, err)
	}

	AddHook(Hook{
		writer:    conn,
		formatter: logFormatter,
	})
	AddHookOrigin(&DefaultFieldHook{GetValue: GetValueImpl})

	Debug("Test Debug")
	Debugf("Test Debugf %v insert", "0.0")
	Info("Test Info")
	Infof("Test Infof %v insert", "0.0")
	Warning("Test Warning")
	Warningf("Test Warningf %v insert", "0.0")
}

func GetValueImpl() string {
	return "with its default value"
}
