package repo

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Product struct {
	Id         int     `json:"id" db:"id"`
	Title      string  `json:"title"  db:"title"`
	Dscription string  `json:"description"  db:"description"`
	Price      float64 `json:"price"  db:"price"`
	ImgUrl     string  `json:"image"  db:"image"`
}
type ProductRepo interface {
	Create(p Product) (*Product, error)
	Get(productID int) (*Product, error)
	List() ([]*Product, error)
	Delete(productID int) error
	Update(product Product) (*Product, error)
}

type productRepo struct {
	dbCon *sqlx.DB
}

func NewProductRepo(dbCon *sqlx.DB) ProductRepo {
	return &productRepo{
		dbCon: dbCon,
	}
}
func (r *productRepo) Create(p Product) (*Product, error) {
	query := `
		INSERT INTO products (
			title,
			description,
			price,
			image
		)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`
	err := r.dbCon.QueryRow(
		query,
		p.Title,
		p.Dscription,
		p.Price,
		p.ImgUrl,
	).Scan(&p.Id)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &p, nil
}
func (r *productRepo) Get(productID int) (*Product, error) {
	query := `
		SELECT id, title, description, price, image
		FROM products
		WHERE id = $1
	`

	var product Product

	err := r.dbCon.Get(&product, query, productID)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
func (r *productRepo) List() ([]*Product, error) {
	query := `
		SELECT id, title, description, price, image
		FROM products
		ORDER BY id
	`

	var products []*Product

	err := r.dbCon.Select(&products, query)
	if err != nil {
		return nil, err
	}

	return products, nil
}
func (r *productRepo) Delete(productID int) error {
	query := `
		DELETE FROM products
		WHERE id = $1
	`

	_, err := r.dbCon.Exec(query, productID)

	if err != nil {
		return err
	}

	return nil
}
func (r *productRepo) Update(product Product) (*Product, error) {
	query := `
		UPDATE products
		SET
			title = $1,
			description = $2,
			price = $3,
			image = $4
		WHERE id = $5
		RETURNING id
	`

	err := r.dbCon.QueryRow(
		query,
		product.Title,
		product.Dscription,
		product.Price,
		product.ImgUrl,
		product.Id,
	).Scan(&product.Id)

	if err != nil {
		return nil, err
	}

	return &product, nil
}
