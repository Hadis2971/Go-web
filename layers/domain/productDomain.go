package domain

import (
	"errors"

	"github.com/Hadis2971/go_web/layers/dataAccess"
	"github.com/Hadis2971/go_web/models"
)

type ProductDomain struct {
	productDataAccess *dataAccess.ProductDataAccess
}

func NewProductDomain(productDataAccess *dataAccess.ProductDataAccess) *ProductDomain {
	return &ProductDomain{productDataAccess: productDataAccess}
}

func (pd *ProductDomain) HandleCreateProduct(product models.ProductReqPayload) (*models.Product, error) {
	sqlResult, err := pd.productDataAccess.CreateProduct(&product)

	id, _ := sqlResult.LastInsertId()

	newProduct := models.Product{
		ID: models.ProductId(id),
		Name: product.Name,
		Description: product.Description,
		Price: float64(product.Price),
		Stock: product.Stock,
	}

	if errors.Is(err, dataAccess.ErrorCreateProduct) {
		return nil, dataAccess.ErrorCreateProduct
	}

	if errors.Is(err, dataAccess.ErrorCreateProductMissingFields) {
		return nil, dataAccess.ErrorCreateProductMissingFields
	}

	return &newProduct, nil
}

func (pd *ProductDomain) HandleGetAllProducts() ([]models.Product, error) {
	products, err := pd.productDataAccess.GetAllProducts()

	if errors.Is(err, dataAccess.ErrorGetAllProducts) {
		return nil, dataAccess.ErrorGetAllProducts
	}

	return products, nil
}

func (pd *ProductDomain) HandleGetProductById(id models.ProductId) (*models.Product, error) {
	product, err := pd.productDataAccess.GetProductById(id)

	if errors.Is(err, dataAccess.ErrorGetProductById) {
		return nil, dataAccess.ErrorGetProductById
	}

	if errors.Is(err, dataAccess.ErrorGetProductByIdMissingId) {
		return nil, dataAccess.ErrorGetProductByIdMissingId
	}

	return product, nil
}

func (pd *ProductDomain) HandleDeleteProduct(id models.ProductId) error {
	err := pd.productDataAccess.DeleteProduct(id)



	if errors.Is(err, dataAccess.ErrorDeleteProduct) {
		return dataAccess.ErrorDeleteProduct
	}

	if errors.Is(err, dataAccess.ErrorDeleteProductMissingId) {
		return dataAccess.ErrorDeleteProductMissingId
	}

	return nil
}

func (pd *ProductDomain) HandleUpdateProduct(product models.ProductReqPayload) (*models.Product, error) {
	sqlResult, err := pd.productDataAccess.UpdateProduct(product)

	id, _ := sqlResult.LastInsertId()

	newProduct := models.Product{
		ID: models.ProductId(id),
		Name: product.Name,
		Description: product.Description,
		Price: float64(product.Price),
		Stock: product.Stock,
	    ProductCategory: product.ProductCategory,
	}

	if errors.Is(err, dataAccess.ErrorUpdateProduct) {
		return nil, dataAccess.ErrorUpdateProduct
	}

	if errors.Is(err, dataAccess.ErrorUpdateProductMissingFields) {
		return nil,dataAccess.ErrorUpdateProductMissingFields
	}

	return &newProduct, nil
}