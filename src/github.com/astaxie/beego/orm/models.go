// Beego (http://beego.me/)

// @description beego is an open-source, high-performance web framework for the Go programming language.

// @link        http://github.com/astaxie/beego for the canonical source repository

// @license     http://github.com/astaxie/beego/blob/master/LICENSE

// @authors     astaxie, slene

package orm

import (
	"sync"
)

const (
	od_CASCADE            = "cascade"
	od_SET_NULL           = "set_null"
	od_SET_DEFAULT        = "set_default"
	od_DO_NOTHING         = "do_nothing"
	defaultStructTagName  = "orm"
	defaultStructTagDelim = ";"
)

var (
	modelCache = &_modelCache{
		cache:     make(map[string]*modelInfo),
		cacheByFN: make(map[string]*modelInfo),
	}
	supportTag = map[string]int{
		"-":            1,
		"null":         1,
		"index":        1,
		"unique":       1,
		"pk":           1,
		"auto":         1,
		"auto_now":     1,
		"auto_now_add": 1,
		"size":         2,
		"column":       2,
		"default":      2,
		"rel":          2,
		"reverse":      2,
		"rel_table":    2,
		"rel_through":  2,
		"digits":       2,
		"decimals":     2,
		"on_delete":    2,
		"type":         2,
	}
)

// model info collection
type _modelCache struct {
	sync.RWMutex
	orders    []string
	cache     map[string]*modelInfo
	cacheByFN map[string]*modelInfo
	done      bool
}

// get all model info
func (mc *_modelCache) all() map[string]*modelInfo {
	m := make(map[string]*modelInfo, len(mc.cache))
	for k, v := range mc.cache {
		m[k] = v
	}
	return m
}

// get orderd model info
func (mc *_modelCache) allOrdered() []*modelInfo {
	m := make([]*modelInfo, 0, len(mc.orders))
	for _, table := range mc.orders {
		m = append(m, mc.cache[table])
	}
	return m
}

// get model info by table name
func (mc *_modelCache) get(table string) (mi *modelInfo, ok bool) {
	mi, ok = mc.cache[table]
	return
}

// get model info by field name
func (mc *_modelCache) getByFN(name string) (mi *modelInfo, ok bool) {
	mi, ok = mc.cacheByFN[name]
	return
}

// set model info to collection
func (mc *_modelCache) set(table string, mi *modelInfo) *modelInfo {
	mii := mc.cache[table]
	mc.cache[table] = mi
	mc.cacheByFN[mi.fullName] = mi
	if mii == nil {
		mc.orders = append(mc.orders, table)
	}
	return mii
}

// clean all model info.
func (mc *_modelCache) clean() {
	mc.orders = make([]string, 0)
	mc.cache = make(map[string]*modelInfo)
	mc.cacheByFN = make(map[string]*modelInfo)
	mc.done = false
}

// Clean model cache. Then you can re-RegisterModel.
// Common use this api for test case.
func ResetModelCache() {
	modelCache.clean()
}