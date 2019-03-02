package microlog

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"net"
	"testing"
)

func TestMicrolog(t *testing.T) {
	logFormatter := &log.JSONFormatter{
		FieldMap: *&log.FieldMap{
			log.FieldKeyMsg: "message",
		},
	}
	conn, err := net.Dial("tcp", "localhost:51401")
	assert.Nil(t, err)
	hook := New(conn, logFormatter)
	logrusConfig := NewLogrusConfig(logFormatter, log.DebugLevel, log.Fields{"project": "test", "module": "stock-backend-internal-api"}, true, hook)
	err = logrusConfig.InstallConfig()
	assert.Nil(t, err)
	//import log "github.com/sillyhatxu/microlog"
	//log.Debug("This is info log.")
	//log.Info("This is info log.")
	Debug("This is debug log.HEIHEI")
	Info("This is info log.")
}
