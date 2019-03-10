# microlog

# This project was canceled

> microlog use logrus library

> Its logrus client.

```
import (
    log "github.com/sillyhatxu/microlog"
    "github.com/sirupsen/logrus"
    "net"
)

func main() {
	logFormatter := &logrus.JSONFormatter{
		FieldMap: *&logrus.FieldMap{
			logrus.FieldKeyMsg: "message",
		},
	}
	conn, err := net.Dial("tcp", "localhost:8001")
	if err != nil {
		log.Fatal("net.Dial error.", err)
	}
	hook := log.New(conn, logFormatter)
	logrusConfig := log.NewLogrusConfig(logFormatter, logrus.DebugLevel, logrus.Fields{"project": "test-project", "module": "mode-test"}, true, hook)
	err = logrusConfig.InstallConfig()
	if err != nil {
		log.Fatal("logrus config initial error.", err)
	}
	
	
	log.Debug("This is info log.")
	log.Info("This is info log.")
}

```
