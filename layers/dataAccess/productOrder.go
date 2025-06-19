package dataAccess

import (
	"database/sql"
	"errors"

	"github.com/Hadis2971/go_web/models"
)

type ProductOrderDataAccess struct {
	dbConnection *sql.DB
}

var (
	ErrorCreatingProductOrder = errors.New("Error Creating Product Order")
	ErrorProductOutOfStock = errors.New("Error Procut Out Of Stock")
	ErrorUpdatingProductOrder = errors.New("Error Updating Product Order")
	ErrorGettingUserProductOrders = errors.New("Error Getting User Product Orders")
	ErrorGettingProductOrders = errors.New("Error Getting Product Orders")
	ErrorDeletingProductOrders = errors.New("Error Deleting Product Orders")
)

func NewProductOrderDataAccess(dbConnection *sql.DB) *ProductOrderDataAccess {
	return &ProductOrderDataAccess{dbConnection: dbConnection}
}

func (po *ProductOrderDataAccess) CreateProductOrder(productOrder models.ProductOrder) error {
	queryCheckStock := "SELECT stock FROM Product WHERE id = ?"
	
	type QueryCheckStockResponse struct {
		Stock int
	}

	queryCheckStockResponse := QueryCheckStockResponse{}
	
	row:= po.dbConnection.QueryRow(queryCheckStock, productOrder.ProductId)

	row.Scan(queryCheckStockResponse.Stock)

	if queryCheckStockResponse.Stock == 0 {
		return ErrorProductOutOfStock
	}

	queryCreate := "INSERT INTO Product_Order (quantity, user_id, product_id, order_id) VALUES(?, ?, ?)"

	_, err := po.dbConnection.Exec(queryCreate, productOrder.Quantity, productOrder.UserId, productOrder.ProductId, productOrder.OrderId)

	if err != nil {
		return ErrorCreatingProductOrder
	}

	return nil
}

func (po *ProductOrderDataAccess) GetOrdersByUserId(userId models.UserId) ([]models.ProductOrder, error) {
	query := "SELECT * FROM Product_Order WHERE user_id = ?"
	var productOrder models.ProductOrder
	productOrders := []models.ProductOrder{}

	rows, err := po.dbConnection.Query(query, userId)

	for rows.Next() {
		rows.Scan(productOrder.ID, productOrder.Quantity, productOrder.UserId, productOrder.ProductId, productOrder.OrderId, productOrder.CreatedOn, productOrder.UpdatedOn)

		productOrders = append(productOrders, productOrder)
	}

	if err != nil {
		return nil, ErrorGettingUserProductOrders
	}

	return productOrders, nil
}

func (po *ProductOrderDataAccess) GetOrdersByOrderId(orderId models.OrderId) ([]models.ProductOrder, error) {
	query := "SELECT * FROM Product_Order WHERE order_id = ?"
	var productOrder models.ProductOrder
	productOrders := []models.ProductOrder{}

	rows, err := po.dbConnection.Query(query, orderId)

	for rows.Next() {
		rows.Scan(productOrder.ID, productOrder.Quantity, productOrder.UserId, productOrder.ProductId, productOrder.OrderId, productOrder.CreatedOn, productOrder.UpdatedOn)

		productOrders = append(productOrders, productOrder)
	}
	
	if err != nil {
		return nil, ErrorGettingProductOrders
	}


	return productOrders, nil
}

func (po *ProductOrderDataAccess) UpdateProductOrder(productOrder models.ProductOrder) error {
	query := "UPDATE Product_Order SET quantity = ? WHERE id = ?"

	_, err := po.dbConnection.Exec(query, productOrder.Quantity, productOrder.ID)
	
	if err != nil {
		return ErrorUpdatingProductOrder
	}

	return nil
}

func (po *ProductOrderDataAccess) DeleteProductOrder(productOrderId models.ProductOrderId) error {
	query := "DELETE Product_Order WHERE order_id = ?"

	_, err := po.dbConnection.Exec(query, productOrderId)

	if err != nil {
		return ErrorDeletingProductOrders
	}

	return nil
}