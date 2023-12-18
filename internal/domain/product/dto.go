package product

import "net/http"

type (
	OrderRequest struct {
		Email  string `json:"email"`
		Amount int    `json:"amount"`
	}

	OrderResponse struct {
		CreatedAt string `json:"created_at"`
		ID        string `json:"id"`
	}

	LoginRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)

func (receiver *OrderRequest) Bind(r *http.Request) error {
	return nil
}

func (receiver *LoginRequest) Bind(r *http.Request) error {
	return nil
}
