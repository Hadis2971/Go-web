package dataAccess

import (
	"database/sql"
	"errors"

	"github.com/Hadis2971/go_web/models"
)

type ProductDataAccess struct {
	dbConnection *sql.DB
}

var (
	ErrorCreateProduct = errors.New("Error Creating Product")
	ErrorCreateProductMissingFields = errors.New("Missing Fields!!!")
	ErrorGetAllProducts = errors.New("Error Getting Products!!!")
	ErrorGetProductById = errors.New("Error Getting Product!!!")
	ErrorGetProductByIdMissingId = errors.New("Missing ID Field!!!")
	ErrorDeleteProduct = errors.New("Error Delete Product!!!")
	ErrorDeleteProductMissingId = errors.New("Missing ID Field!!!")
	ErrorUpdateProduct = errors.New("Error Updating Product!!!")
	ErrorUpdateProductMissingFields = errors.New("Missing Fields!!!")
)

func NewProductDataAccess(dbConnection *sql.DB) *ProductDataAccess {
	return &ProductDataAccess{dbConnection: dbConnection}
}

func (pda *ProductDataAccess) CreateProduct(product *models.CreateProductReq)  error {

	query := "INSERT INTO Product (product_name, price, description, stock) VALUES (?, ?, ?, ?)"

	if (product.Name == "" || product.Price == 0 || product.Description == "" || product.Stock == 0) {
		return ErrorCreateProductMissingFields
	}

	_, err := pda.dbConnection.Exec(query, product.Name, product.Price, product.Description, product.Stock)


	if err != nil {
		return ErrorCreateProduct
	}


	return nil
}

func (pda *ProductDataAccess) GetAllProducts() ([]models.Product, error) {
	query := "SELECT * FROM Product"
	products := []models.Product{}
	var product models.Product

	rows, err := pda.dbConnection.Query(query)


	defer rows.Close()

	if err != nil {
		return nil, ErrorGetAllProducts
	}

	
	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Stock, &product.CreatedOn, &product.UpdatedOn)
		
		if err != nil {
			return nil, ErrorGetAllProducts
		}

		products = append(products, product)
	}


	return products, nil
}

func (pda *ProductDataAccess) GetProductById(id models.ProductId) (*models.Product, error) {
	query := "SELECT * FROM Product WHERE id = ?"
	var product models.Product

	if id == 0 {
		return nil, ErrorGetProductByIdMissingId
	}

	row := pda.dbConnection.QueryRow(query, id)

	err := row.Scan(&product.ID, &product.Name, &product.Price, &product.Description, &product.Stock, &product.CreatedOn, &product.UpdatedOn)

	if err != nil {
		return nil, ErrorGetProductById
	}

	return &product, nil
}

func (pda *ProductDataAccess) DeleteProduct(id models.ProductId) error {
	query := "DELETE FROM Product Where id = ?"

	if id == 0 {
		return ErrorDeleteProductMissingId
	}

	_, err := pda.dbConnection.Exec(query, id)

	if err != nil {
		return ErrorDeleteProduct
	}

	return nil
}

func (pda *ProductDataAccess) UpdateProduct(product models.Product) error {
	query := "UPDATE Product SET product_name = ?, price = ?, description = ?, stock = ? WHERE id = ?"

	if (product.ID == 0 || product.Name == "" || product.Price == 0 || product.Description == "") {
		return ErrorUpdateProductMissingFields
	}

	_, err := pda.dbConnection.Exec(query, product.Name, product.Price, product.Description, product.Stock, product.ID)

	if err != nil {
		return ErrorUpdateProduct
	}

	return nil
}