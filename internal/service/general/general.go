package general

import (
	"context"
	"errors"
	"store-back/internal/domain/product"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func (s *Service) SelectProducts(ctx context.Context) (res []product.Product, err error) {
	res, err = s.productRepository.SelectProducts(ctx)

	return
}

func (s *Service) SelectOrdersByClient(ctx context.Context, email string) (res []product.Order, err error) {
	res, err = s.productRepository.SelectOrdersByClient(ctx, email)

	return
}

func (s *Service) GetOrderByID(ctx context.Context, id string) (res product.Order, err error) {
	res, err = s.productRepository.GetOrderByID(ctx, id)

	return
}

func (s *Service) CreateOrder(ctx context.Context, req product.OrderRequest) (res product.Order, err error) {
	res, err = s.productRepository.CreateOrder(ctx, req)

	return
}

func (s *Service) Login(ctx context.Context, req product.LoginRequest) (res product.LoginResponse, err error) {
	data, err := s.productRepository.Login(ctx, req)
	if err != nil {
		return
	}

	if data.Password != req.Password {
		err = errors.New("invalid password")
		return
	}

	res.AccessToken, err = GenerateJWT(data.Email)

	return
}

func (s *Service) Register(ctx context.Context, req product.LoginRequest) (res product.User, err error) {
	res, err = s.productRepository.Register(ctx, req)

	return
}

func GenerateJWT(email string) (string, error) {
	secretKey := []byte("secret")
	// Create a new JWT token with the email as a claim
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	})

	// Sign the token with your secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
