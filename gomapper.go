package gomapper

import (
	"errors"
	"fmt"
	"reflect"
)

type IObjectMapper interface {
	Register(converter IConverter)
	Map(src interface{}, destType interface{}, options *Options) (interface{}, bool)
	MustMap(src interface{}, destType interface{}, options *Options) interface{}
}

type objectMapperImpl struct {
	converters []IConverter
}

func New() IObjectMapper {
	return &objectMapperImpl{converters: make([]IConverter, 0)}
}

func (mapper *objectMapperImpl) Register(converter IConverter) {
	converter.SetMapper(mapper)
	mapper.converters = append(mapper.converters, converter)
}

func (mapper *objectMapperImpl) Map(src interface{}, destType interface{}, options *Options) (interface{}, bool) {
	if options == nil {
		options = &Options{}
	}
	for _, converter := range mapper.converters {
		if converter.Supports(src, destType, *options) {
			return converter.Convert(src, *options), true
		}
	}
	return src, false
}

func (mapper *objectMapperImpl) MustMap(src interface{}, destType interface{}, options *Options) interface{} {
	mapped, ok := mapper.Map(src, destType, options)
	if ok {
		return mapped
	}
	s := getType(reflect.TypeOf(src))
	d := getType(reflect.TypeOf(destType))
	panic(errors.New(fmt.Sprintf("gomapper: could not find converter [%s] -> [%s]", s, d)))
}
