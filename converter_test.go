package gomapper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type userDto struct {
	Username string
	Age int
}

type userModel struct {
	ModelID uint
	Name string
	Age int
}

// userDto -> userModel
type userDtoModelConverter struct {
	BaseConverter
}

func (c *userDtoModelConverter) Supports(src interface{}, dest interface{}, options Options) bool {
	return c.IsType(src, userDto{}) && c.IsType(dest, userModel{})
}

func (c *userDtoModelConverter) Convert(obj interface{}, options Options) interface{} {
	user := obj.(userDto)
	return userModel{
		ModelID: options.DefaultUint("id", 0),
		Name:    user.Username,
		Age:     user.Age,
	}
}

func TestConverter_Convert(t *testing.T) {
	converter := userDtoModelConverter{}
	dto := userDto{
		Username: "Test",
		Age:      36,
	}

	f := converter.Supports(dto, userModel{}, Options{})
	assert.Equal(t, true, f)

	mapped := converter.Convert(dto, Options{})
	assert.IsType(t, userModel{}, mapped)

	assert.Equal(t, "Test", mapped.(userModel).Name)
	assert.Equal(t, 36, mapped.(userModel).Age)
}
