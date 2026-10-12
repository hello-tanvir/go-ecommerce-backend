package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type User struct {
	ID          int    `db:"id" json:"id"`
	FirstName   string `db:"first_name" json:"first_name"`
	LastName    string `db:"last_name" json:"last_name"`
	Email       string `db:"email" json:"email"`
	Password    string `db:"password" json:"password"`
	IsShopOwner bool   `db:"is_shop_owner" json:"is_shop_owner"`
}

type UserRepo interface {
	Create(user User) (*User, error)
	Find(email, pass string) (*User, error)
}

type userRepo struct {
	dbCon *sqlx.DB
}

func NewUserRepo(dbCon *sqlx.DB) UserRepo {
	return &userRepo{
		dbCon: dbCon,
	}

}

func (r *userRepo) Create(user User) (*User, error) {
	query := `
		INSERT INTO users (
			first_name,
			last_name,
			email,
			password,
			is_shop_owner
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id
	`
	err := r.dbCon.QueryRow(
		query,
		user.FirstName,
		user.LastName,
		user.Email,
		user.Password,
		user.IsShopOwner,
	).Scan(&user.ID)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) Find(email, pass string) (*User, error) {
	var user User

	query := `
        SELECT id, first_name, last_name, email, password, is_shop_owner
        FROM users
        WHERE email = $1 AND password = $2
    `

	err := r.dbCon.Get(&user, query, email, pass)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &user, nil
}
