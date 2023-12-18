package product

import "time"

type Product struct {
	ID       string  `json:"id" db:"id"`
	Title    string  `json:"title" db:"title"`
	Image    string  `json:"image" db:"image"`
	Price    float64 `json:"price" db:"price"`
	Category string  `json:"category" db:"category"`
}

type Order struct {
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ID        string    `json:"id" db:"id"`
	UserID    string    `json:"user_id" db:"user_id"`
	Amount    int       `json:"amount" db:"amount"`
}

type User struct {
	ID       string `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}
