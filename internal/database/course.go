package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Course struct {
	DB          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{DB: db}
}

func (c *Course) CreateCourse(name, description, categoryID string) (Course, error) {
	c.ID = uuid.New().String()
	c.Name = name
	c.Description = description
	c.CategoryID = categoryID

	query := `INSERT INTO courses (id, name, description, category_id) VALUES (?, ?, ?, ?)`
	_, err := c.DB.Exec(query, c.ID, c.Name, c.Description, c.CategoryID)
	if err != nil {
		return Course{}, err
	}
	return Course{DB: c.DB, ID: c.ID, Name: c.Name, Description: c.Description, CategoryID: c.CategoryID}, nil
}
func (c *Course) FindAll() ([]Course, error) {
	query := `SELECT id, name, description, category_id FROM courses`
	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.Name, &course.Description, &course.CategoryID); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return courses, nil
}

func (c *Course) FindByCategoryID(id string) ([]Course, error) {
	query := `SELECT id, name, description FROM courses WHERE category_id = ?`
	rows, err := c.DB.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var courses []Course
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.ID, &course.Name, &course.Description); err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return courses, nil
}
