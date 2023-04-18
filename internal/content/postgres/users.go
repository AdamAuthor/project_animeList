package postgres

import (
	"animeList/internal/content"
	"animeList/internal/models"
	"context"
	"github.com/jmoiron/sqlx"
)

func (D *Database) User() content.UserRepo {
	if D.user == nil {
		D.user = NewUserRepo(D.conn)
	}

	return D.user
}

type UserRepo struct {
	conn *sqlx.DB
}

func NewUserRepo(conn *sqlx.DB) content.UserRepo {
	return &UserRepo{conn: conn}
}

func (u UserRepo) Create(ctx context.Context, user *models.User) error {
	_, err := u.conn.Exec("INSERT INTO users(email, password, name, gender, age) VALUES ($1, $2, $3, $4, $5)", user.Email, user.Password, user.Name, user.Gender, user.Age)
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepo) Parse(ctx context.Context, email string, password string) ([]*models.User, error) {
	var user []*models.User
	err := u.conn.Select(&user, `SELECT id, gender, email, age, name, password
FROM users
WHERE email = $1 AND password = $2`, email, password)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserRepo) ResetPassword(ctx context.Context, email string, password string) error {
	_, err := u.conn.Exec("UPDATE users SET password = $1 WHERE email = $2", password, email)

	if err != nil {
		return err
	}
	return nil
}
