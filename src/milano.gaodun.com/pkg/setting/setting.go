package setting

import (
	"os"
	"strings"

	"github.com/apex/log"
	"github.com/apex/log/handlers/multi"
	"github.com/apex/log/handlers/text"
	"milano.gaodun.com/conf"
)

//var Logger *log.Entry
// TODO log 不应该在setting当中
// 设置log
var Logger *log.Entry = GrayLog()
var MySQLLogger = grayMySQLLog()

var grayLog = getGrayLog()

func getGrayLog() log.Handler {
	return &alg
}

type GdGrayLog struct {
	LogFields string
	LogLevel int
}

func NewGdGrayLog() *GdGrayLog {
	return &GdGrayLog{LogFields:conf.LOG_FIELDS, LogLevel:conf.LOG_LEVEL}
}

func (gl *GdGrayLog) GrayLog(newFields ...map[string]interface{}) *log.Entry {
	if conf.IS_SHOW_CONSOLE {
		t := text.New(os.Stderr)
		log.SetHandler(multi.New(grayLog, t))
	} else {
		log.SetHandler(multi.New(grayLog))
	}

	fields := make(log.Fields)
	grayFields := gl.LogFields
	grayFieldsArray := strings.Split(grayFields, ",")
	if len(grayFieldsArray) > 0 {
		for i := 0; i < len(grayFieldsArray); i++ {
			temp := strings.Split(grayFieldsArray[i], ":")
			if len(temp) > 1 {
				fields[string(temp[0])] = temp[1]
			}
		}
	}

	if newFields != nil {
		for k, v := range newFields[0] {
			fields[k] = v
		}
	}
	level := gl.LogLevel
	log.SetLevel(log.Level(level))
	return log.WithFields(fields)
}

// 通用 log
func GrayLog(newFields ...map[string]interface{}) *log.Entry {
	return NewGdGrayLog().GrayLog(newFields...)
}

// MySQL log
func grayMySQLLog(newFields ...map[string]interface{}) *log.Entry {
	g := NewGdGrayLog()
	g.LogFields = conf.LOG_FIELDS_MYSQL
	return g.GrayLog(newFields...)
}