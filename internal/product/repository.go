package product

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(p *Product) error {

	query := `
		INSERT INTO products
		(title, description, price, size)
		VALUES (?, ?, ?, ?)
	`

	result, err := r.db.Exec(
		query,
		p.Title,
		p.Description,
		p.Price,
		p.Size,
	)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	p.ID = id

	return nil
}