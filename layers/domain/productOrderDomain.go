package domain

import (
	"github.com/Hadis2971/go_web/layers/dataAccess"
	"github.com/Hadis2971/go_web/models"
)

type ProductOrderDomain struct {
	productOrderDataAccess *dataAccess.ProductOrderDataAccess
}

func NewProductOrderDomain(productOrderDataAccess *dataAccess.ProductOrderDataAccess) *ProductOrderDomain {
	return &ProductOrderDomain{productOrderDataAccess: productOrderDataAccess}
}

func (pod ProductOrderDomain) HandleCreateProductOrder(productOrder models.ProductOrder) error {
	err := pod.productOrderDataAccess.CreateProductOrder(productOrder)

	if err != nil {
		return err
	}

	return nil
}

func (pod ProductOrderDomain) HandleCreateProductOrderWithMultipleProducts(productOrders []models.ProductOrder) error {
	err := pod.productOrderDataAccess.CreateProductOrderWithMultipleProducts(productOrders)

	if err != nil {
		return err
	}

	return nil
}

func (pod ProductOrderDomain) HandleGetProuctOrdersByUserId(userId models.UserId) ([]models.ProductOrder, error) {
	productOrders, err := pod.productOrderDataAccess.GetOrdersByUserId(userId)

	if err != nil {
		return nil, err
	}

	return productOrders, nil
}

func (pod ProductOrderDomain) HandleGetProductOrdersByOrderId(orderId models.OrderId) ([]models.ProductOrder, error) {
	productOrders, err := pod.productOrderDataAccess.GetOrdersByOrderId(orderId)

	if err != nil {
		return nil, err
	}

	return productOrders, nil
}

func (pod ProductOrderDomain) HandleGetProductOrdersByUserIdAndOrderId(userId models.UserId, orderId models.OrderId) ([]models.ProductAndUser, error) {
	productOrdersAndUser, err := pod.productOrderDataAccess.GetOrderByUserIdAndOrderId(userId, orderId)

	if err != nil {
		return nil, err
	}

	return productOrdersAndUser, nil
}

func (pod ProductOrderDomain) HandleUpdateProductOrder(productOrder models.ProductOrder) error {
	err := pod.productOrderDataAccess.UpdateProductOrder(productOrder)

	if err != nil {
		return err
	}

	return nil
}

func (pod ProductOrderDomain) HandleDeleteProductOrder(orderId models.OrderId) error {
	err := pod.productOrderDataAccess.DeleteProductOrder(orderId)

	if err != nil {
		return err
	}

	return nil
}