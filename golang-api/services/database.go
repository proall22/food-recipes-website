package services

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"recipehub/models"
)

type DatabaseService struct {
	db *sql.DB
}

func NewDatabaseService() *DatabaseService {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgrespassword@localhost:5432/recipehub?sslmode=disable"
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("Failed to ping database: %v", err))
	}

	return &DatabaseService{db: db}
}

func (s *DatabaseService) CreateUser(user *models.User) error {
	query := `
		INSERT INTO users (email, username, first_name, last_name, password_hash, bio, avatar)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRow(
		query,
		user.Email,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Password,
		user.Bio,
		user.Avatar,
	).Scan(&user.ID, &user.CreatedAt, &user.UpdatedAt)

	return err
}

func (s *DatabaseService) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, username, first_name, last_name, password_hash, bio, avatar, 
		       is_verified, is_active, email_verified_at, created_at, updated_at
		FROM users 
		WHERE email = $1
	`

	err := s.db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Bio,
		&user.Avatar,
		&user.IsVerified,
		&user.IsActive,
		&user.EmailVerifiedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *DatabaseService) GetUserByUsername(username string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, username, first_name, last_name, password_hash, bio, avatar, 
		       is_verified, is_active, email_verified_at, created_at, updated_at
		FROM users 
		WHERE username = $1
	`

	err := s.db.QueryRow(query, username).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Bio,
		&user.Avatar,
		&user.IsVerified,
		&user.IsActive,
		&user.EmailVerifiedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *DatabaseService) GetUserByID(id string) (*models.User, error) {
	user := &models.User{}
	query := `
		SELECT id, email, username, first_name, last_name, password_hash, bio, avatar, 
		       is_verified, is_active, email_verified_at, created_at, updated_at
		FROM users 
		WHERE id = $1
	`

	err := s.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Email,
		&user.Username,
		&user.FirstName,
		&user.LastName,
		&user.Password,
		&user.Bio,
		&user.Avatar,
		&user.IsVerified,
		&user.IsActive,
		&user.EmailVerifiedAt,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *DatabaseService) SavePasswordResetToken(userID, token string) error {
	// Create password_reset_tokens table if it doesn't exist
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS password_reset_tokens (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			token VARCHAR(255) NOT NULL,
			expires_at TIMESTAMP NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		)
	`
	
	if _, err := s.db.Exec(createTableQuery); err != nil {
		return err
	}

	// Delete any existing tokens for this user
	deleteQuery := `DELETE FROM password_reset_tokens WHERE user_id = $1`
	if _, err := s.db.Exec(deleteQuery, userID); err != nil {
		return err
	}

	// Insert new token (expires in 1 hour)
	insertQuery := `
		INSERT INTO password_reset_tokens (user_id, token, expires_at)
		VALUES ($1, $2, $3)
	`
	
	expiresAt := time.Now().Add(time.Hour)
	_, err := s.db.Exec(insertQuery, userID, token, expiresAt)
	return err
}

func (s *DatabaseService) VerifyPasswordResetToken(token string) (string, error) {
	var userID string
	query := `
		SELECT user_id FROM password_reset_tokens 
		WHERE token = $1 AND expires_at > NOW()
	`

	err := s.db.QueryRow(query, token).Scan(&userID)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (s *DatabaseService) UpdateUserPassword(userID, hashedPassword string) error {
	query := `UPDATE users SET password_hash = $1, updated_at = NOW() WHERE id = $2`
	_, err := s.db.Exec(query, hashedPassword, userID)
	return err
}

func (s *DatabaseService) DeletePasswordResetToken(token string) error {
	query := `DELETE FROM password_reset_tokens WHERE token = $1`
	_, err := s.db.Exec(query, token)
	return err
}

func (s *DatabaseService) VerifyEmailToken(token string) (string, error) {
	// Create email_verification_tokens table if it doesn't exist
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS email_verification_tokens (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
			token VARCHAR(255) NOT NULL,
			expires_at TIMESTAMP NOT NULL,
			created_at TIMESTAMP DEFAULT NOW()
		)
	`
	
	if _, err := s.db.Exec(createTableQuery); err != nil {
		return "", err
	}

	var userID string
	query := `
		SELECT user_id FROM email_verification_tokens 
		WHERE token = $1 AND expires_at > NOW()
	`

	err := s.db.QueryRow(query, token).Scan(&userID)
	if err != nil {
		return "", err
	}

	return userID, nil
}

func (s *DatabaseService) MarkEmailAsVerified(userID string) error {
	query := `UPDATE users SET is_verified = true, email_verified_at = NOW() WHERE id = $1`
	_, err := s.db.Exec(query, userID)
	return err
}

func (s *DatabaseService) CreateRecipePurchase(purchase *models.RecipePurchase) error {
	query := `
		INSERT INTO recipe_purchases (recipe_id, user_id, amount, payment_method, payment_reference, status)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id, created_at, updated_at
	`

	err := s.db.QueryRow(
		query,
		purchase.RecipeID,
		purchase.UserID,
		purchase.Amount,
		purchase.PaymentMethod,
		purchase.PaymentReference,
		purchase.Status,
	).Scan(&purchase.ID, &purchase.CreatedAt, &purchase.UpdatedAt)

	return err
}

func (s *DatabaseService) UpdateRecipePurchaseStatus(purchaseID, status string) error {
	query := `UPDATE recipe_purchases SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := s.db.Exec(query, status, purchaseID)
	return err
}

func (s *DatabaseService) GetRecipePurchaseByReference(reference string) (*models.RecipePurchase, error) {
	purchase := &models.RecipePurchase{}
	query := `
		SELECT id, recipe_id, user_id, amount, payment_method, payment_reference, status, created_at, updated_at
		FROM recipe_purchases 
		WHERE payment_reference = $1
	`

	err := s.db.QueryRow(query, reference).Scan(
		&purchase.ID,
		&purchase.RecipeID,
		&purchase.UserID,
		&purchase.Amount,
		&purchase.PaymentMethod,
		&purchase.PaymentReference,
		&purchase.Status,
		&purchase.CreatedAt,
		&purchase.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return purchase, nil
}

func (s *DatabaseService) TrackRecipeView(recipeID, userID, ipAddress, userAgent string) error {
	query := `
		INSERT INTO recipe_views (recipe_id, user_id, ip_address, user_agent)
		VALUES ($1, $2, $3, $4)
	`

	var userIDPtr *string
	if userID != "" {
		userIDPtr = &userID
	}

	_, err := s.db.Exec(query, recipeID, userIDPtr, ipAddress, userAgent)
	return err
}
