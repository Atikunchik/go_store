package service

import "GolangwithFrame/src/domain/model"

type CartService interface {
	CreateCart(cart model.Cart)
	DeleteCategory(login string) error
	FindAllCarts() []model.Cart
	GetUserCart(login string) (model.Cart, error)
}

func (service *Service) CreateCart(cart model.Cart) model.Cart {
	service.Repository.CreateCart(cart)
	return cart
}

func (service *Service) FindAllCarts() []model.Cart {
	return service.Repository.FindAllCarts()
}

func (service *Service) DeleteCart(login string) error {
	return service.Repository.DeleteCartByLogin(login)

}

func (service *Service) GetUserCart(login string) ([]model.Cart, error) {
	return service.Repository.GetUserCart(login)
}
