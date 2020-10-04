package gomapper

import (
	"errors"
)

type Options map[string]interface{}

var ErrKeyNotFound = errors.New("option: key not found")

func (opt Options) Value(key string) (interface{}, bool) {
	if opt.Has(key) == false {
		return nil, false
	}
	return opt[key], true
}

func (opt Options) MustValue(key string) interface{} {
	if !opt.Has(key) {
		panic(ErrKeyNotFound)
	}
	return opt[key]
}

func (opt Options) Has(key string) bool {
	if _, ok := opt[key]; ok {
		return ok
	}
	return false
}

func (opt Options) DefaultInt(key string, def int) int {
	if opt.Has(key) {
		return opt[key].(int)
	}
	return def
}

func (opt Options) DefaultUint(key string, def uint) uint {
	if opt.Has(key) {
		return opt[key].(uint)
	}
	return def
}

func (opt Options) DefaultString(key string, def string) string {
	if opt.Has(key) {
		return opt[key].(string)
	}
	return def
}

func (opt Options) DefaultBool(key string, def bool) bool {
	if opt.Has(key) {
		return opt[key].(bool)
	}
	return def
}

func (opt Options) DefaultFloat(key string, def float64) float64 {
	if opt.Has(key) {
		return opt[key].(float64)
	}
	return def
}
