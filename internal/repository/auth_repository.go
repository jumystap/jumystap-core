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

func (r *AuthRepository) StoreUser (user *model.User) (int, error) {
    var id int
    query := `
        INSERT INTO users (name, role_id, email, phone, password, description) 
        VALUES ( ?, 4, ?, ?, ?, 'admin')
        RETURNING id;
    `  
    err := r.db.QueryRow(query, user.Name, user.Email, user.Phone, user.Password).Scan(&id)
    if err != nil {
        return 0, err
    }
    
    return id, nil
}
