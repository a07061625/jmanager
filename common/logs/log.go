package logs

import (
	"log"
	"os"
	"io"
	"jmanager/common/constants"
)

type jLog struct {
	infoLog *log.Logger
	warningLog *log.Logger
	errorLog *log.Logger
}

var Log = jLog{}

func init() {
	logFile, err := os.OpenFile(constants.JM_LOG, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("打开日志文件失败:", err)
	}

	Log.infoLog = log.New(io.MultiWriter(logFile), "[INFO]", log.Ldate|log.Ltime|log.Lshortfile)
	Log.warningLog = log.New(io.MultiWriter(logFile), "[WARNING]", log.Ldate|log.Ltime|log.Lshortfile)
	Log.errorLog = log.New(io.MultiWriter(logFile), "Error:", log.Ldate|log.Ltime|log.Lshortfile)
}

func (this *jLog) LogError(err interface{}) {
	this.errorLog.Println(err)
}

func (this *jLog) LogInfo(msg string)  {
	this.infoLog.Println(msg)
}

func (this *jLog) LogWarning(msg string)  {
	this.warningLog.Println(msg)
}
