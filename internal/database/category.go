package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	DB          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCategory(db *sql.DB) *Category {
	return &Category{DB: db}
}

func (c *Category) CreateCategory(name, description string) (Category, error) {
	c.ID = uuid.New().String()
	c.Name = name
	c.Description = description

	query := `INSERT INTO categories (id, name, description) VALUES (?, ?, ?)`
	_, err := c.DB.Exec(query, c.ID, c.Name, c.Description)
	if err != nil {
		return Category{}, err
	}
	return Category{DB: c.DB, ID: c.ID, Name: c.Name, Description: c.Description}, nil
}

func (c *Category) FindAll() ([]Category, error) {
	query := `SELECT id, name, description FROM categories`
	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return categories, nil
}

func (c *Category) FindByCourseID(courseID string) (Category, error) {
	query := `SELECT c.id, c.name, c.description FROM categories c
		JOIN courses co ON c.id = co.category_id WHERE co.id = ?`
	rows, err := c.DB.Query(query, courseID)
	if err != nil {
		return Category{}, err
	}
	defer rows.Close()
	var category Category
	if rows.Next() {
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return Category{}, err
		}
	}
	if err := rows.Err(); err != nil {
		return Category{}, err
	}
	if category.ID == "" {
		return Category{}, sql.ErrNoRows
	}
	return category, nil
}
