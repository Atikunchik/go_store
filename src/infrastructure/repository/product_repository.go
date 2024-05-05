package repository

import (
	model "GolangwithFrame/src/domain/model"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)


type ProductRepository interface {
	CreateProduct(product model.Product)
	UpdateProduct(product model.Product) error
	DeleteProduct(product model.Product) error
	FindAllProducts() []model.Product
	GetProduct(id int) (model.Product, error)
}


func (db *Database) GetProduct(id int) (model.Product, error) {
	product := model.Product{}
	err := db.Connection.Where("id = ?", id).First(&product).Error
	if err != nil {
		return product, err
	} else {
		return product, nil
	}
}

func (db *Database) CreateProduct(product model.Product) {
	db.Connection.Create(&product)
}
func (db *Database) UpdateProduct(product model.Product) error {
	currentProduct := model.Product{}

	err := db.Connection.Where("id = ?", product.Id).First(&currentProduct).Error
	if err != nil {
		return db.Connection.Where("id = ?", product.Id).First(&currentProduct).Error
	}
	db.Connection.Save(&product)
	return nil

}
func (db *Database) DeleteProduct(product model.Product) error {
	err := db.Connection.Where("id = ?", product.Id).First(&product).Error
	db.Connection.Delete(&product)
	return err
}
func (db *Database) FindAllProducts() []model.Product {
	var products []model.Product
	db.Connection.Set("gorm:auto_preload", true).Order("id").Find(&products)
	return products
}
