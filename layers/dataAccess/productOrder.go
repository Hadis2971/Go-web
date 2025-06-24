package dataAccess

import (
	"context"
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
	ErrorGettingUserProductOrdersByUserIdAndOrderId = errors.New("Error Getting Product Orders By User ID and Order ID")
)

func NewProductOrderDataAccess(dbConnection *sql.DB) *ProductOrderDataAccess {
	return &ProductOrderDataAccess{dbConnection: dbConnection}
}

func (po *ProductOrderDataAccess) CreateProductOrder(productOrder models.ProductOrder) error {
	ctx := context.Background()

	tx, _ := po.dbConnection.BeginTx(ctx, nil)

	defer tx.Rollback()

	queryCheckStock := "SELECT (stock >= ?) WHERE id = ?"
	var inStock bool

	err := tx.QueryRowContext(ctx, queryCheckStock, productOrder.Quantity, productOrder.ID).Scan(&inStock)

	if err == sql.ErrNoRows || !inStock {
		return ErrorProductOutOfStock
	}


	queryUpdateStock := "UPDATE Product SET stock = stock - ? WHERE id = ?"

	_, err = tx.ExecContext(ctx, queryUpdateStock, productOrder.Quantity, productOrder.ID)

	if err != nil {
		return err
	}

	queryCreate := "INSERT INTO Product_Order (quantity, user_id, product_id, order_id) VALUES(?, ?, ?, ?)"

	_, err = tx.ExecContext(ctx, queryCreate, productOrder.Quantity, productOrder.UserId, productOrder.ProductId, productOrder.OrderId)

	if err != nil {
		return ErrorCreatingProductOrder
	}

	err = tx.Commit()

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

func (po *ProductOrderDataAccess) GetOrderByUserIdAndOrderId(userId models.UserId, orderId models.OrderId) ([]models.ProductAndUser, error) {
	query := `SELECT User.id AS user_id, User.username, Product_Order.quantity, Product_Order.created_on AS order_created, Product_Order.updated_on AS order_updated FROM User JOIN Product_Order ON Product_Order.
	user_id = ? AND Product_Order.order_id = ?;`

	var productOrderAndUser models.ProductAndUser
	productOrdersAndUser := []models.ProductAndUser{}


	rows, err := po.dbConnection.Query(query, userId, orderId);

	if err != nil {
		return nil, ErrorGettingUserProductOrdersByUserIdAndOrderId
	}

	for rows.Next() {
		rows.Scan(productOrderAndUser.UserId, productOrderAndUser.Username, productOrderAndUser.Quantity, productOrderAndUser.OrderCreated, productOrderAndUser.OrderUpdated)

		productOrdersAndUser = append(productOrdersAndUser, productOrderAndUser)
	}

	return productOrdersAndUser, nil
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