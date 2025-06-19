package dataAccess

import (
	"database/sql"
	"errors"

	"github.com/Hadis2971/go_web/models"
)

type ProductOrder struct {
	dbConnection *sql.DB
}

var (
	ErrorCreatingProductOrder = errors.New("Error Creating Product Order")
	ErrorUpdatingProductOrder = errors.New("Error Updating Product Order")
	ErrorGettingUserProductOrders = errors.New("Error Getting User Product Orders")
	ErrorGettingProductOrders = errors.New("Error Getting Product Orders")
)

func NewProductOrderDataAccess(dbConnection *sql.DB) *ProductOrder {
	return &ProductOrder{dbConnection: dbConnection}
}

func (po *ProductOrder) CreateProductOrder(productOrder models.ProductOrder) error {
	query := "INSERT INTO Product_Order (quantity, userId, productId, orderId) VALUES(?, ?, ?)"

	_, err := po.dbConnection.Exec(query, productOrder.Quantity, productOrder.UserId, productOrder.ProductId, productOrder.OrderId)

	if err != nil {
		return ErrorCreatingProductOrder
	}

	return nil
}

func (po *ProductOrder) GetOrdersByUserId(userId int) error {
	query := "SELECT * FROM Product_Order WHERE userId = ?"

	_, err := po.dbConnection.Exec(query, userId)

	if err != nil {
		return ErrorGettingUserProductOrders
	}

	return nil
}

func (po *ProductOrder) GetOrdersByOrderId(orderId int) error {
	query := "SELECT * FROM Product_Order WHERE orderId = ?"

	_, err := po.dbConnection.Exec(query, orderId)
	
	if err != nil {
		return ErrorUpdatingProductOrder
	}


	return nil
}

func (po *ProductOrder) UpdateProductOrder(productOrder models.ProductOrder) error {
	query := "UPDATE Product_Order SET quantity = ? WHERE id = ?"

	_, err := po.dbConnection.Exec(query, productOrder.Quantity, productOrder.ID)
	
	if err != nil {
		return ErrorUpdatingProductOrder
	}

	return nil
}