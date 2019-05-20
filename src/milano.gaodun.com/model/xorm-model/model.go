package xorm_model

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"milano.gaodun.com/pkg/utils"
	"reflect"
	"github.com/go-xorm/builder"
)

// model 基础类
type ModelBase struct {
	*xorm.Engine
	s *xorm.Session
}

func (g *ModelBase) NewModelSession() *xorm.Session {
	if g.s == nil {
		g.s = g.NoAutoTime()
	}
	return g.s
}

func (g *ModelBase) GetSession() *xorm.Session {
	return g.NewModelSession()
}

// 拼接查询与条件
func (g *ModelBase) WhereAndParam(fieldName string, value interface{}) {
	g.s.Where(builder.And(builder.Eq{fieldName:value}))
}

// sql 查询条件
func (g *ModelBase) IsExist(bean ...interface{}) (bool, error) {
	return g.s.Exist(bean ...)
}

// sql 查询条件
func (g *ModelBase) SqlFind(query interface{}, args ...interface{}) {
	g.s.SQL(query, args...)
}

// sql 查询条件并返回结果
func (g *ModelBase) SqlQuery(sql ... interface{}) ([]map[string][]byte, error) {
	return g.s.Query(sql...)
}

// 拼接查询或条件
func (g *ModelBase) WhereOrParam(fieldName string, value interface{}) {
	g.s.Or(builder.And(builder.Eq{fieldName:value}))
}

// 分页 默认 50 条
func (g *ModelBase) Page(limit, start int) {
	if start <= 1 {
		start = 0
	}
	if limit == 0 {
		limit = 50
	}
	g.s.Limit(limit, start * limit)
}

// 利害反射 拼接查询条件
func (g *ModelBase) WhereParams(pg interface{}) error {

	getType := reflect.TypeOf(pg)
	getValue := reflect.ValueOf(pg)
	if getValue.Kind() != reflect.Struct {
		return fmt.Errorf("类型错误")
	}

	//
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		kind := getValue.Field(i)
		f := utils.Mapper.Obj2Table(field.Name)
		switch kind.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			if kind.Int() > 0 {
				g.s.Where(builder.And(builder.Eq{f:value}))
			}
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			if kind.Uint() > 0 {
				g.s.Where(builder.And(builder.Eq{f:value}))
			}
		case reflect.Float32, reflect.Float64:
			if kind.Float() > 0 {
				g.s.Where(builder.And(builder.Eq{f:value}))
			}
		case reflect.String:
			if kind.String() != "" {
				g.s.Where(builder.And(builder.Eq{f:value}))
			}
		}
	}

	return nil
}
