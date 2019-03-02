# microlog

> microlog use logrus library

> Its logrus client.

```
import log "github.com/sillyhatxu/microlog"

func main(){
    logFormatter := &log.JSONFormatter{
        FieldMap: *&log.FieldMap{
            log.FieldKeyMsg: "message",
        },
    }
    conn, err := net.Dial("tcp", "localhost:8080")
    assert.Nil(t, err)
    hook := New(conn, logFormatter)
    logrusConfig := NewLogrusConfig(logFormatter, log.DebugLevel, log.Fields{"project": "test-project", "module": "project-module"}, true, hook)
    err = logrusConfig.InstallConfig()

	log.Debug("This is info log.")
	log.Info("This is info log.")
}

```