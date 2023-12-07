package product

import "context"

type Repository interface {
	SelectProducts(ctx context.Context) (res []Product, err error)

	SelectOrdersByClient(ctx context.Context, email string) (res []Order, err error)
	GetOrderByID(ctx context.Context, id string) (res Order, err error)
	CreateOrder(ctx context.Context, req OrderRequest) (res Order, err error)
	
	Login(ctx context.Context, req LoginRequest) (res User, err error)
	Register(ctx context.Context, req LoginRequest) (res User, err error)
}
