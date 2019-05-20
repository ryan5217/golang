package utils

import (
	"milano.gaodun.com/pkg/setting"
	"strconv"
	"strings"
)

var GraySql GrayXormSql

// 实现 xrom 打印日志接口
type GrayXormSql struct {
}

func (GrayXormSql) Write(p []byte) (n int, err error) {
	if strings.Contains(strings.ToUpper(string(p)), "[SQL]") {
		s := string(p)
		t := strings.Split(s, "took:")
		tst := strings.Trim(t[1], " ")
		tst = strings.Replace(tst, "ms\n", "", -1)
		sqlTime, _ := strconv.ParseFloat(tst, 10)
		setting.MySQLLogger.WithField("user_sql_exec_time", sqlTime).Info(s)
	}
	return 0, nil
}
