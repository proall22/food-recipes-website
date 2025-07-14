package models

import (
	"database/sql/driver"
	"time"
)

type User struct {
	ID              string     `json:"id" db:"id"`
	Email           string     `json:"email" db:"email"`
	Username        string     `json:"username" db:"username"`
	FirstName       string     `json:"first_name" db:"first_name"`
	LastName        string     `json:"last_name" db:"last_name"`
	Password        string     `json:"-" db:"password_hash"`
	Bio             string     `json:"bio" db:"bio"`
	Avatar          string     `json:"avatar" db:"avatar"`
	IsVerified      bool       `json:"is_verified" db:"is_verified"`
	IsActive        bool       `json:"is_active" db:"is_active"`
	EmailVerifiedAt *time.Time `json:"email_verified_at" db:"email_verified_at"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time  `json:"updated_at" db:"updated_at"`
}

type Recipe struct {
	ID           string    `json:"id" db:"id"`
	Title        string    `json:"title" db:"title"`
	Slug         string    `json:"slug" db:"slug"`
	Description  string    `json:"description" db:"description"`
	FeaturedImage string   `json:"featured_image" db:"featured_image"`
	PrepTime     int       `json:"prep_time" db:"prep_time"`
	CookTime     int       `json:"cook_time" db:"cook_time"`
	TotalTime    int       `json:"total_time" db:"total_time"`
	Servings     int       `json:"servings" db:"servings"`
	Difficulty   string    `json:"difficulty" db:"difficulty"`
	CuisineType  string    `json:"cuisine_type" db:"cuisine_type"`
	Price        float64   `json:"price" db:"price"`
	IsPremium    bool      `json:"is_premium" db:"is_premium"`
	Status       string    `json:"status" db:"status"`
	AuthorID     string    `json:"author_id" db:"author_id"`
	CategoryID   string    `json:"category_id" db:"category_id"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type RecipePurchase struct {
	ID               string    `json:"id" db:"id"`
	RecipeID         string    `json:"recipe_id" db:"recipe_id"`
	UserID           string    `json:"user_id" db:"user_id"`
	Amount           float64   `json:"amount" db:"amount"`
	PaymentMethod    string    `json:"payment_method" db:"payment_method"`
	PaymentReference string    `json:"payment_reference" db:"payment_reference"`
	Status           string    `json:"status" db:"status"`
	CreatedAt        time.Time `json:"created_at" db:"created_at"`
	UpdatedAt        time.Time `json:"updated_at" db:"updated_at"`
}

type NullString struct {
	String string
	Valid  bool
}

func (ns *NullString) Scan(value interface{}) error {
	if value == nil {
		ns.String, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	ns.String = value.(string)
	return nil
}

func (ns NullString) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.String, nil
}
