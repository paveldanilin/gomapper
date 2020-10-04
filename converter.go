package gomapper

import (
	"reflect"
)

type IConverter interface {
	SetMapper(mapper IObjectMapper)
	GetMapper() IObjectMapper
	Supports(src interface{}, dest interface{}, options Options) bool
	Convert(obj interface{}, options Options) interface{}
}

type BaseConverter struct {
	mapper IObjectMapper
}

func (base *BaseConverter) SetMapper(mapper IObjectMapper) {
	base.mapper = mapper
}

func (base *BaseConverter) GetMapper() IObjectMapper {
	return base.mapper
}

func (base *BaseConverter) GetType(v interface{}) string {
	return getType(reflect.TypeOf(v))
}

func (base *BaseConverter) IsArray(v interface{}) bool {
	t := reflect.TypeOf(v).Kind()
	return t == reflect.Array || t == reflect.Slice
}

func (base *BaseConverter) IsMap(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Map
}

func (base *BaseConverter) IsStruct(v interface{}) bool {
	return reflect.TypeOf(v).Kind() == reflect.Struct
}

func (base *BaseConverter) IsType(v interface{}, t interface{}) bool {
	if base.IsStruct(v) && base.IsStruct(t) {
		return base.GetType(v) == base.GetType(t)
	}
	if base.IsArray(v) && base.IsArray(t) {
		return base.GetType(v) == base.GetType(t)
	}
	if base.IsMap(v) && base.IsMap(t) {
		return base.GetType(v) == base.GetType(t)
	}
	return false
}

func getType(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice:
		elType := t.PkgPath() + "." + getType(t.Elem())
		return elType
	case reflect.Struct:
		return t.PkgPath() + "." + t.Name()
	}
	return "undefined"
}
