package gomapper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptions_Has(t *testing.T) {
	opts := Options{
		"testOption": "optionValue",
	}
	assert.Equal(t, true, opts.Has("testOption"))
}


func TestOptions_MustGet(t *testing.T) {
	opts := Options{
		"age": 36,
	}
	assert.Equal(t, 36, opts.MustValue("age").(int))
}

func TestOptions_MustGet_Panic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			if r != ErrKeyNotFound {
				t.Error(r)
			}
		}
	}()
	opts := Options{}
	opts.MustValue("opt")
}

func TestOptions_DefaultBool(t *testing.T) {
	opts := Options{
		"flag": false,
	}
	assert.Equal(t, false, opts.DefaultBool("flag", true))
}

func TestOptions_DefaultInt(t *testing.T) {
	opts := Options{
		"number": 123,
	}
	assert.Equal(t, 123, opts.DefaultInt("number", 0))
}
