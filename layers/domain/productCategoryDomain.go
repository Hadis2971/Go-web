package domain

import (
	"github.com/Hadis2971/go_web/layers/dataAccess"
	"github.com/Hadis2971/go_web/models"
)

type ProductCategoryDomain struct {
	productCategodyDataAccess *dataAccess.ProductCategoryDataAccess
}

func NewProductCategoryDomain(productCategodyDataAccess *dataAccess.ProductCategoryDataAccess) *ProductCategoryDomain {
	return &ProductCategoryDomain{productCategodyDataAccess: productCategodyDataAccess}
}

func (pcd ProductCategoryDomain) HandleGetAllProductCategory() ([]models.ProductCategory, error) {
	var productCategories []models.ProductCategory

	productCategories, err := pcd.productCategodyDataAccess.GetAllProductCategory()

	if err != nil {
		return nil, err
	}

	return productCategories, nil
}

func (pcd ProductCategoryDomain) HandleCreateProductCategory(productCategory models.NewProductCategoryRequst) error {


	err := pcd.productCategodyDataAccess.CreateProductCategory(productCategory)

	if err != nil {
		return err
	}

	return nil
}

func (pcd ProductCategoryDomain) HandleUpdateProductCategory(productCategory models.ProductCategory) error {

	err := pcd.productCategodyDataAccess.UpdateProductCategory(productCategory)

	if err != nil {
		return err
	}

	return nil
}

func (pcd ProductCategoryDomain) HandleDeleteProductCategory(id models.ProductCategoryId) error {

	err := pcd.productCategodyDataAccess.DeleteProductCategory(id)

	if err != nil {
		return err
	}

	return nil
}