package dataAccess

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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
	ErrorProductNotFound = errors.New("Error Product Not Found!!!")
)

func NewProductOrderDataAccess(dbConnection *sql.DB) *ProductOrderDataAccess {
	return &ProductOrderDataAccess{dbConnection: dbConnection}
}

func (po *ProductOrderDataAccess) CreateProductOrder(productOrder models.ProductOrder) error {
	ctx := context.Background()

	tx, _ := po.dbConnection.BeginTx(ctx, nil)

	defer tx.Rollback()

	fmt.Println(productOrder.ProductId, productOrder.Quantity)

	queryCheckStock := "SELECT (stock > ?) FROM Product WHERE id = ?"
	var inStock bool

	err := tx.QueryRowContext(ctx, queryCheckStock, productOrder.Quantity, productOrder.ProductId).Scan(&inStock)

	if err == sql.ErrNoRows {
		return ErrorProductNotFound
	}

	if  !inStock {
		return ErrorProductOutOfStock
	}


	queryUpdateStock := "UPDATE Product SET stock = stock - ? WHERE id = ?"

	_, err = tx.ExecContext(ctx, queryUpdateStock, productOrder.Quantity, productOrder.ProductId)

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

func (po *ProductOrderDataAccess) CreateProductOrderWithMultipleProducts(productOrders []models.ProductOrder) error {
	tx, err := po.dbConnection.Begin()

	if err != nil {
		return err
	}

	defer tx.Rollback()

	var inStock bool

	checkIfInStockQuery := "SELECT (stock >= ?) FROM Product WHERE id = ?"

	for _, productOrder := range productOrders {
		err := tx.QueryRow(checkIfInStockQuery, productOrder.Quantity, productOrder.ProductId).Scan(&inStock)

		if err == sql.ErrNoRows {
			return ErrorProductNotFound
		}

		if !inStock {
			return ErrorProductOutOfStock
		}
	}

	updateStockQuery := "UPDATE Product SET stock = stock - ? WHERE id = ?"

	for _, productOrder := range productOrders {
		_, err := tx.Exec(updateStockQuery, productOrder.Quantity, productOrder.ProductId)

		if err != nil {
			return err
		}
	}


	createProuctOrderQuery := "INSERT INTO Product_Order (user_id, product_id, quantity, order_id) VALUES(?,?,?,?)"

	for _, productOrder := range productOrders {
		_, err := tx.Exec(createProuctOrderQuery, productOrder.UserId, productOrder.ProductId, productOrder.Quantity, productOrder.OrderId)

		if err != nil {
			return err
		}
	}

	err = tx.Commit();

	if err != nil {
		return err
	}

	return nil
}

func (po *ProductOrderDataAccess) GetOrdersByUserId(userId models.UserId) ([]models.ProductOrder, error) {
	query := "SELECT * FROM Product_Order WHERE user_id = ?"
	var productOrder models.ProductOrder
	productOrders := []models.ProductOrder{}

	rows, err := po.dbConnection.Query(query, userId)

	for rows.Next() {
		rows.Scan(&productOrder.ID, &productOrder.Quantity, &productOrder.UserId, &productOrder.ProductId, &productOrder.OrderId, &productOrder.CreatedOn, &productOrder.UpdatedOn)

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
		rows.Scan(&productOrder.ID, &productOrder.Quantity, &productOrder.UserId, &productOrder.ProductId, &productOrder.CreatedOn, &productOrder.UpdatedOn, &productOrder.OrderId)

		productOrders = append(productOrders, productOrder)
	}
	
	if err != nil {
		return nil, ErrorGettingProductOrders
	}


	return productOrders, nil
}

func (po *ProductOrderDataAccess) GetOrderByUserIdAndOrderId(userId models.UserId, orderId models.OrderId) ([]models.ProductAndUser, error) {
	query := `SELECT User.id AS user_id, User.username, Product_Order.quantity, Product_Order.product_id AS product_id, Product_Order.created_on AS order_created, Product_Order.updated_on AS order_updated FROM
	User JOIN Product_Order ON Product_Order.user_id = User.id WHERE Product_Order.user_id = ? AND Product_Order.order_id = ?`

	var productOrderAndUser models.ProductAndUser
	productOrdersAndUser := []models.ProductAndUser{}


	rows, err := po.dbConnection.Query(query, userId, orderId);

	if err != nil {
		return nil, ErrorGettingUserProductOrdersByUserIdAndOrderId
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&productOrderAndUser.UserId, &productOrderAndUser.Username, &productOrderAndUser.Quantity, &productOrderAndUser.ProductId, &productOrderAndUser.OrderCreated, &productOrderAndUser.OrderUpdated)

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

func (po *ProductOrderDataAccess) DeleteProductOrder(orderId models.OrderId) error {
	tx, err := po.dbConnection.Begin()

	defer tx.Rollback()

	if err != nil {
		return err
	}

	productOrders, err := po.GetOrdersByOrderId(orderId)

	if err != nil {
		return err;
	}

	updateProductStockQuery := "UPDATE Product Set stock = stock + ? WHERE id = ?"

	for _, productOrder := range productOrders {
		_, err := tx.Exec(updateProductStockQuery, productOrder.Quantity, productOrder.ID)

		if err != nil {
			return ErrorDeletingProductOrders
		}
	}

	deleteProductOrderQuery := "DELETE Product_Order WHERE order_id = ?"

	_, err = po.dbConnection.Exec(deleteProductOrderQuery, orderId)

	if err != nil {
		return ErrorDeletingProductOrders
	}

	err = tx.Commit()


	if err != nil {
		return ErrorDeletingProductOrders
	}

	return nil
}