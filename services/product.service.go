package services

import (
	"fmt"
	"go-rest-api/models"

	"gorm.io/gorm"
)

type ProductService struct {
	DB *gorm.DB
}

func (ps *ProductService) CreateProduct(product *models.Product) (*models.Product, error) {
	if err := ps.DB.Create(product).Error; err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *ProductService) GetProductById(id uint) (*models.Product, error) {
	var product models.Product
	if err := ps.DB.First(&product, id).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (ps *ProductService) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	if err := ps.DB.Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (ps *ProductService) UpdateProduct(id uint, product *models.Product) (*models.Product, error) {
	var existingProduct models.Product

	if err := ps.DB.First(&existingProduct, id).Error; err != nil {
		return nil, fmt.Errorf("product with id %d not found", id)
	}

	if err := ps.DB.Model(&existingProduct).Updates(product).Error; err != nil {
		return nil, err
	}

	return &existingProduct, nil
}

func (ps *ProductService) DeleteProduct(id uint) error {
	if err := ps.DB.Delete(&models.Product{}, id).Error; err != nil {
		return err
	}

	return nil
}
