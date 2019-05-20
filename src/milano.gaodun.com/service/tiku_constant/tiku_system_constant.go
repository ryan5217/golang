package tiku_constant

// 获取 gd_tiku_system_constant 常量

import (
	"fmt"
	tsc "milano.gaodun.com/model/tiku_system_constant"
	"sync"
	"time"
)

var TkConstList = NewTikuConstantList()

const (
	MaxSize = 50 // 前 50 个
	MaxTime = 10  // 10 分钟过期时间
)

type tikuConstant struct {
	key  string
	val  string
	next *tikuConstant
	t    time.Time
}

type TikuConstantList struct {
	first  *tikuConstant
	length int
	last   *tikuConstant
	rw     *sync.RWMutex
}

func NewTikuConstant(key, val string) *tikuConstant {
	return &tikuConstant{key, val, nil, time.Now()}
}

func NewTikuConstantList() *TikuConstantList {
	tl := new(TikuConstantList)
	tl.first = NewTikuConstant("", "")
	tl.last = NewTikuConstant("", "")
	tl.rw = new(sync.RWMutex)
	return tl
}

// 获取 key 对应的值, 注：10分钟缓存
func (tl *TikuConstantList) GetKey(key string) (string, error) {
	defer tl.rw.Unlock()
	tl.rw.Lock()
	next := tl.first
	var last *tikuConstant
	for next.next != nil && next.next.key != key {
		last = next
		next = next.next
	}
	var rear *tikuConstant
	// 如果找到
	if next.next != nil {
		rear = next.next
		if time.Since(rear.t).Minutes() > MaxTime {
			s, err := tl.getKeyByDB(key)
			if err != nil {
				return "", err
			}
			if s.Id == 0 {
				return "", fmt.Errorf("empty")
			}
			rear.t = time.Now()
			rear.val = s.Thevalue
		}
		next.next = rear.next
		rear.next = tl.first.next
		tl.first.next = rear
		return rear.val, nil
	}
	s, err := tl.getKeyByDB(key)
	if err != nil {
		return "", err
	}
	if s.Id == 0 {
		return "", fmt.Errorf("empty")
	}

	rear = NewTikuConstant(key, s.Thevalue)
	rear.next = tl.first.next
	tl.first.next = rear
	tl.last = last
	if tl.length >= MaxSize {
		tl.last.next = nil
	} else {
		tl.length++
	}

	return rear.val, nil
}

// 数据库中读
func (tl *TikuConstantList) getKeyByDB(key string) (tsc.GdTikuSystemConstant, error) {
	parm := tsc.SearchParam{Thekey: key}
	tscm := tsc.NewTikuSystemConstantModel(&parm)
	return tscm.GetKey()
}
