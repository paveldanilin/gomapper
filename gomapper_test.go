package gomapper

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

type str struct {
	val string
}

type num struct {
	val int
}

type valConverter struct {
	BaseConverter
}

func (c *valConverter) Supports(src interface{}, dest interface{}, options Options) bool {
	return c.IsType(src, str{}) && c.IsType(dest, num{})
}

func (c *valConverter) Convert(obj interface{}, options Options) interface{} {
	strVal := obj.(str).val
	numVal, _ := strconv.ParseInt(strVal, 10, 64)
	return num{
		val: int(numVal),
	}
}

func TestObjectMapperImpl_Map(t *testing.T) {
	mapper := New()
	mapper.Register(&valConverter{})

	n := mapper.MustMap(str{val: "554"}, num{}, nil)

	assert.Equal(t, 554, n.(num).val)
}
