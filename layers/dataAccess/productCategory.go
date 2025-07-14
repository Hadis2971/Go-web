package dataAccess

import (
	"database/sql"
	"errors"

	"github.com/Hadis2971/go_web/models"
)


var (
	ErrorGettingAllProductCategories = errors.New("Error Getting All Product Categories!!!")
	ErrorCreatingNewProductCategory = errors.New("Error Creating New Product Category!!!")
	ErrorUpdatingProductCategory = errors.New("Error Updating Product Category!!!")
	ErrorDeletingProductCategory = errors.New("Error Deleting Product Category!!!")
)

type ProductCategoryDataAccess struct {
	dbConnection *sql.DB
}

func NewProductCategoryDataAccess(dbConnection *sql.DB) *ProductCategoryDataAccess {
	return &ProductCategoryDataAccess{dbConnection: dbConnection}
}


func (pcda ProductCategoryDataAccess) GetAllProductCategory() ([]models.ProductCategory, error) {
	query := "SELECT * FROM Product_Category"
	var productCategories []models.ProductCategory
	var productCategory models.ProductCategory

	rows, err := pcda.dbConnection.Query(query)

	if err != nil {
		return nil, ErrorGettingAllProductCategories
	}

	defer rows.Close()

	for rows.Next() {
		rows.Scan(&productCategory.ID, &productCategory.Name, &productCategory.Category)

		productCategories = append(productCategories, productCategory)
	}

	return productCategories, nil
}

func (pcda ProductCategoryDataAccess) CreateProductCategory(productCategory models.NewProductCategoryRequst) error {
	query := "INSERT INTO Product_Category (name, category) VALUES(?, ?)"

	_, err := pcda.dbConnection.Exec(query, productCategory.Name, productCategory.Category)

	if err != nil {
		return ErrorCreatingNewProductCategory
	}

	return nil
}

func (pcda ProductCategoryDataAccess) UpdateProductCategory(productCategory models.ProductCategory) error {
	query := "UPDATE Product_Category SET name = ?, SET category = ? WHERE id = ?"

	_, err := pcda.dbConnection.Exec(query, productCategory.Name, productCategory.Category, productCategory.ID)

	if err != nil {
		return ErrorUpdatingProductCategory
	}

	return nil
}

func (pcda ProductCategoryDataAccess) DeleteProductCategory(id models.ProductCategoryId) error {
	query := "DELETE FROM Product_Category WHERE id = ?"

	_, err := pcda.dbConnection.Exec(query, id)

	if err != nil {
		return ErrorDeletingProductCategory
	}

	return nil
}