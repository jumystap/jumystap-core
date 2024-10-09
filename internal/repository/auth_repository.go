package repository

import (
	"database/sql"

	"github.com/jumystap/jumystap-core/internal/model"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
} 

func (r *AuthRepository) GetUserByEmail (email string) (*model.User, error) {
    user := &model.User{}

    query := "SELECT id, name, email, phone, password FROM users WHERE email = ?"
    
    err := r.db.QueryRow(query, email).Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Password)
    if err != nil {
        return nil, err
    }

    return user, nil
}
