package dbconf

import (
	"fmt"
	"milano.gaodun.com/pkg/consul"
	"os"
)

var dbConf = consul.GdConsul

type GdDb struct {
	DriverName string
	DriverDns  string
}

var GDDBConf = GetMySqlConfig()
var GaodunDBConf = GetGaodunMySqlConfig()
var NewtikuDBConf = GetNewtikuMySqlConfig()

func GetMySqlConfig() []GdDb {
	// 数据库配置
	dbAllConfig := []GdDb{
		{DriverName: "mysql", DriverDns: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbConf["GD_PRIMARY_MYSQL_USER"], dbConf["GD_PRIMARY_MYSQL_PASSWORD"], dbConf["GD_PRIMARY_MYSQL_HOST"], dbConf["PUBLIC_MYSQL_DB_PORT"], "gd_primary")},
	}
	return dbAllConfig
}
func GetGaodunMySqlConfig() []GdDb {
	// 数据库配置
	dbAllConfig := []GdDb{
		{DriverName: "mysql", DriverDns: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbConf["GD_PRIMARY_MYSQL_USER"], dbConf["GD_PRIMARY_MYSQL_PASSWORD"], dbConf["GD_PRIMARY_MYSQL_HOST"], dbConf["PUBLIC_MYSQL_DB_PORT"], "gaodun")},
		{DriverName: "mysql", DriverDns: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbConf["TIKU_MYSQL_USER"], dbConf["TIKU_MYSQL_PASSWORD"], dbConf["PUBLIC_MYSQL_DB_HOST"], dbConf["PUBLIC_MYSQL_DB_PORT"], "gaodun")},
		{DriverName: "mysql", DriverDns: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbConf["APIDEA_MYSQL_USER"], dbConf["APIDEA_MYSQL_PASSWORD"], dbConf["PUBLIC_MYSQL_DB_HOST"], dbConf["PUBLIC_MYSQL_DB_PORT"], "gaodun")},
		{DriverName: "mysql", DriverDns: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbConf["NEWTIKU_MYSQL_USER"], dbConf["NEWTIKU_MYSQL_PASSWORD"], dbConf["NEWTIKU_MYSQL_HOST"], dbConf["NEWTIKU_MYSQL_PORT"], dbConf["NEWTIKU_MYSQL_DATABASE"])},
	}
	fmt.Println(dbAllConfig)
	return dbAllConfig
}
func GetNewtikuMySqlConfig() []GdDb {
	// 数据库配置
	dbAllConfig := []GdDb{
		{DriverName: "mysql", DriverDns: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbConf["NEWTIKU_MYSQL_USER"], dbConf["NEWTIKU_MYSQL_PASSWORD"], dbConf["NEWTIKU_MYSQL_HOST"], dbConf["NEWTIKU_MYSQL_PORT"], dbConf["NEWTIKU_MYSQL_DATABASE"])},
		{DriverName: "mysql", DriverDns: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbConf["GAODUN_TIKU_MYSQL_USER"], dbConf["GAODUN_TIKU_MYSQL_PASSWORD"], dbConf["GAODUN_TIKU_MYSQL_HOST"], dbConf["GAODUN_TIKU_MYSQL_DB_PORT"], dbConf["GAODUN_TIKU_MYSQL_DATABASE"])},
		{DriverName: "mysql", DriverDns: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbConf["TIKU_DBO_MYSQL_USER"], dbConf["TIKU_DBO_MYSQL_PASSWORD"], dbConf["TIKU_DBO_MYSQL_HOST"], dbConf["TIKU_DBO_MYSQL_PORT"], dbConf["TIKU_DBO_MYSQL_DATABASE"])},
	}
	fmt.Println(dbAllConfig)
	return dbAllConfig
}

// IsDev ...
func IsDev() bool {
	env := os.Getenv("SYSTEM_ENV")
	if env == "dev" || env == "" {
		return true
	}
	return false
}
