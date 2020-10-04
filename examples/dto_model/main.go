package main

import (
	"fmt"
	"github.com/paveldanilin/gomapper"
)

// DTO
type orderDto struct {
	ID uint
}

type userDto struct {
	Username string
	Orders []orderDto
}

// Model
type orderModel struct {
	ID uint
	Title string
}

type userModel struct {
	ID uint
	Name string
	Orders []orderModel
}

// Repository
type orderRepository struct {

}
func (r *orderRepository) FindByID(orderID uint) orderModel {
	if orderID == 1 {
		return orderModel{
			ID: 1,
			Title: "First order",
		}
	}
	return orderModel{
		ID:    2,
		Title: "Second order",
	}
}

// Converter
type userDtoConverter struct {
	gomapper.BaseConverter
	orders orderRepository
}

func (c *userDtoConverter) Supports(src interface{}, dest interface{}, options gomapper.Options) bool {
	return c.IsType(src, userDto{}) && c.IsType(dest, userModel{})
}

func (c *userDtoConverter) Convert(obj interface{}, options gomapper.Options) interface{} {
	// Resolve by means of DB repository
	ordersModel := make([]orderModel, 0)
	for _, order := range obj.(userDto).Orders {
		ordersModel = append(ordersModel, c.orders.FindByID(order.ID))
	}

	return userModel{
		ID:   options.DefaultUint("id", 0),
		Name: obj.(userDto).Username,
		Orders: ordersModel,
	}
}

func main() {
	mapper := gomapper.New()
	mapper.Register(&userDtoConverter{
		orders: orderRepository{}, // Resolve by means of DB repository
	})

	dto := userDto{Username: "TestUser", Orders: []orderDto{{ID: 1}, {ID: 2}}}

	model := mapper.MustMap(dto, userModel{}, &gomapper.Options{"id": uint(123)})

	fmt.Printf("%v", model)
}
