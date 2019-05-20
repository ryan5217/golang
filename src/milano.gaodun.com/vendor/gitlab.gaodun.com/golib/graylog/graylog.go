package graylog

// graylog
import (
	"os"
	"strconv"
	"strings"
	"sync"

	"fmt"
	"time"

	"github.com/chenxuey/graylog-golang"
)

var (
	GRAY_ITEM string = "golib-graylog"
)

// 获取主机名
func getHostName() string {
	name, err := os.Hostname()

	if err != nil {
		name = "localhost"
	}
	return name
}

var g = gelf.New(gelf.Config{
	GraylogPort:     5504,
	GraylogHostname: "g02.graylog.gaodunwangxiao.com",
	//GraylogHostname: "graylog.gaodunwangxiao.com",
	//MaxChunkSizeWan: 42,
	//MaxChunkSizeLan: 1337,
})

//GdLog 写入日志
// param log string 日志内容
// param item map[string]interface{}
func GdLog(log string, item map[string]interface{}) {
	grayLog.GdLog(log, LOG_INFO, item)
}

//AddGrayLog 添加 info 日志
func AddGrayLog(logInfo ...interface{}) {
	AddGrayLevelLog(LOG_INFO, logInfo...)
}

//AddGrayLevelLog 添加 lever 日志
func AddGrayLevelLog(level LOG_LEVEL, logInfo ...interface{}) {
	var logStr string
	m := make(map[string]interface{})
	if len(logInfo) == 1 {
		logStr = logInfo[0].(string)
	} else {
		logStr = logInfo[0].(string)
		m = logInfo[1].(map[string]interface{})
	}
	m["item"] = SysEnv + GRAY_ITEM
	grayLog.GdLog(logStr, level, m)
}

//GrayXormSql 实现 xrom 打印日志接口
type GrayXormSql struct {
}

func (GrayXormSql) Write(p []byte) (n int, err error) {
	// todo sql 打印到 gray log
	if strings.Contains(strings.ToLower(string(p)), "[sql]") {
		s := string(p)
		t := strings.Split(s, "took:")
		if len(t) > 1 {
			tst := strings.Trim(t[1], " ")
			tst = strings.Replace(tst, "ms\n", "", -1)
			sqlTime, _ := strconv.ParseFloat(tst, 10)
			m := map[string]interface{}{
				"user_sql":           s,
				"user_sql_exec_time": sqlTime,
			}
			AddGrayLevelLog(LOG_DEBUG, s, m)
		} else {
			m := map[string]interface{}{
				"user_sql": s,
			}
			AddGrayLevelLog(LOG_DEBUG, s, m)
		}

	}
	return 0, nil
}

type LOG_LEVEL int

const (
	// !nashtsai! following level also match syslog.Priority value
	LOG_DEBUG LOG_LEVEL = iota
	LOG_INFO
	LOG_WARNING
	LOG_ERR
	LOG_OFF
	LOG_UNKNOWN
)

type GrayLogger struct {
	g         *gelf.Gelf
	GrayCount int
	p         sync.RWMutex
	flag      LOG_LEVEL
}

func NewGray(g *gelf.Gelf, grayCount int, flag LOG_LEVEL) *GrayLogger {
	return &GrayLogger{g: g, GrayCount: grayCount, flag: flag}
}

var grayLog = NewGray(g, 10, LOG_INFO)

// 定时15分钟解锁 gray log
func init() {
	go grayLog.TurnOn()
}

func (gl *GrayLogger) TurnOn() {
	grayCountInt := gl.GrayCount
	for {
		t := time.NewTicker(15 * time.Minute)
		select {
		case <-t.C:
			if gl.GrayCount <= 0 {
				gl.GrayCount = grayCountInt
			}
		}
	}
}

func (gl *GrayLogger) GdLog(log string, level LOG_LEVEL, item map[string]interface{}) {
	if level < gl.flag {
		return
	}
	if gl.GrayCount > 0 {
		gl.p.Lock()
		item["level"] = level
		item["env"] = SysEnv
		item["item"] = GRAY_ITEM
		item["domain"] = SysEnv + GRAY_ITEM
		gl.p.Unlock()
		Client = NewGdAliLogger(ProjectName, LogStoreName, SlsClient)
		err := Client.WithMap(item).PutLog(log)
		if err != nil {
			gl.GrayCount--
			fmt.Println(err)
		}
	}
}

func (gl *GrayLogger) SetFlags(flag LOG_LEVEL) {
	gl.flag = flag
}

// SetFlags sets the output flags for the gray log.
func SetFlags(flag LOG_LEVEL) {
	grayLog.SetFlags(flag)
}
