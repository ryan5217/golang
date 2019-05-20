package graylog

import (
	"fmt"
	"time"

	sls "github.com/aliyun/aliyun-log-go-sdk"
	"github.com/gogo/protobuf/proto"
	"sync"
)

// ali yun 日志类 新
type GdAliLogger struct {
	Client       *sls.Client
	ProjectName  string
	LogStoreName string
	cols         map[string]interface{}
	p            sync.RWMutex
}

// 返回新的 client
func NewGdAliLogger(ProjectName, LogStoreName string, client *sls.Client) *GdAliLogger {
	g := GdAliLogger{}
	g.Client = client
	g.ProjectName = ProjectName
	g.LogStoreName = LogStoreName
	g.cols = map[string]interface{}{}
	return &g
}

// 重置
func (g *GdAliLogger) ResetField(ProjectName, LogStoreName string) {
	g.ProjectName = ProjectName
	g.LogStoreName = LogStoreName
}

func (g *GdAliLogger) WithField(key string, val interface{}) *GdAliLogger {
	g.p.Lock()
	g.cols[key] = val
	g.p.Unlock()
	return g
}

func (g *GdAliLogger) WithMap(m map[string]interface{}) *GdAliLogger {
	g.p.Lock()
	for k, v := range m {
		g.cols[k] = v
	}
	g.p.Unlock()

	return g
}

func (g *GdAliLogger) PutLog(message interface{}) error {
	logs := []*sls.Log{}
	content := []*sls.LogContent{}
	g.cols["short_message"] = message
	for k, v := range g.cols {
		content = append(content, &sls.LogContent{
			Key:   proto.String(fmt.Sprintf("col_%s", k)),
			Value: proto.String(fmt.Sprint(v)),
		})
	}
	log := &sls.Log{
		Time:     proto.Uint32(uint32(time.Now().Unix())),
		Contents: content,
	}
	logs = append(logs, log)

	loggroup := &sls.LogGroup{
		Topic:  proto.String(getHostName()),
		Source: proto.String(LocalIPV4),
		Logs:   logs,
	}
	return g.Client.PutLogs(ProjectName, LogStoreName, loggroup)
}

// info
func (g *GdAliLogger) InfoLog(message interface{}) error {
	g.cols["level"] = LOG_INFO
	return g.PutLog(message)
}

// LOG_WARNING
func (g *GdAliLogger) WarningLog(message interface{}) error {
	g.cols["level"] = LOG_WARNING
	return g.PutLog(message)
}

// LOG_ERR
func (g *GdAliLogger) ErrLog(message interface{}) error {
	g.cols["level"] = LOG_ERR
	return g.PutLog(message)
}

// LOG_OFF
func (g *GdAliLogger) OffLog(message interface{}) error {
	g.cols["level"] = LOG_OFF
	return g.PutLog(message)
}
