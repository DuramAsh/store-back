package postgres

import (
	"context"
	"store-back/internal/domain/product"

	"github.com/jmoiron/sqlx"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) SelectProducts(ctx context.Context) (res []product.Product, err error) {
	query := `
		SELECT id, title, image, price, category
		FROM products`

	err = r.db.SelectContext(ctx, &res, query)

	return
}

func (r *ProductRepository) SelectOrdersByClient(ctx context.Context, email string) (res []product.Order, err error) {
	query := `
		SELECT created_at, id
		FROM orders
		WHERE user_id=$1`

	args := []any{email}

	err = r.db.SelectContext(ctx, &res, query, args...)

	return
}

func (r *ProductRepository) GetOrderByID(ctx context.Context, id string) (res product.Order, err error) {
	query := `
		SELECT created_at, id
		FROM orders
		WHERE id=$1`

	args := []any{id}

	err = r.db.GetContext(ctx, &res, query, args...)

	return
}

func (r *ProductRepository) CreateOrder(ctx context.Context, req product.OrderRequest) (res product.Order, err error) {
	query := `
		INSERT INTO orders (user_id, amount)
		VALUES ($1, $2)
		RETURNING created_at, id`	

	args := []any{req.Email, req.Amount}

	err = r.db.GetContext(ctx, &res, query, args...)

	return
}

func (r *ProductRepository) Login(ctx context.Context, req product.LoginRequest) (res product.User, err error) {
	query := `
		SELECT id, email, password
		FROM users
		WHERE email=$1`

	args := []any{req.Email}

	err = r.db.GetContext(ctx, &res, query, args...)

	return
}

func (r *ProductRepository) Register(ctx context.Context, req product.LoginRequest) (res product.User, err error) {
	query := `
		INSERT INTO users (email, password)
		VALUES ($1, $2)
		RETURNING id, email, password`

	args := []any{req.Email, req.Password}

	err = r.db.GetContext(ctx, &res, query, args...)

	return
}

// func (r *ProductRepository) Add(ctx context.Context, data author.Entity) (id string, err error) {
// 	query := `
// 		INSERT INTO authors (full_name, pseudonym, specialty)
// 		VALUES ($1, $2, $3)
// 		RETURNING id`

// 	args := []any{data.FullName, data.Pseudonym, data.Specialty}

// 	err = r.db.QueryRowContext(ctx, query, args...).Scan(&id)
// 	if err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			err = store.ErrorNotFound
// 		}
// 	}

// 	return
// }

// func (r *ProductRepository) Get(ctx context.Context, id string) (dest author.Entity, err error) {
// 	query := `
// 		SELECT id, full_name, pseudonym, specialty
// 		FROM authors
// 		WHERE id=$1`

// 	args := []any{id}

// 	if err = r.db.GetContext(ctx, &dest, query, args...); err != nil {
// 		if errors.Is(err, sql.ErrNoRows) {
// 			err = store.ErrorNotFound
// 		}
// 	}

// 	return
// }
