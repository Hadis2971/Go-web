package models


const (
	Phone ProductCategoryId = iota + 1
	Laptop 
	DesktopComputer 
	Appliance
	Game
	TV
)

type ProductCategoryId int
type CategoryId int

type ProductCategory struct {
	ID ProductCategoryId `json:"id"`
	Name string `json:"name"`
	Category CategoryId
}

type NewProductCategoryRequst struct {
	Name string `json:"name"`
	Category int `json:"category"`
}