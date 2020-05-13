package db

import (
	"cozy-forum/models"
	"log"

	"errors"
)

// ErrAlreadyExists ...
var ErrAlreadyExists error = errors.New("already exists")

// CreateCategory ...
func CreateCategory(categoryname string) error {

	log.Printf("Creating category %v...\n", categoryname)
	createCategories, err := db.Prepare(`
		INSERT INTO categories
		(categoryname)
		VALUES (?);
	`)
	if err != nil {
		return err
	}

	_, err = createCategories.Exec(categoryname)
	if err != nil && err.Error() == "UNIQUE constraint failed: categories.categoryname" {
		return ErrAlreadyExists
	}

	return err
}

// GetAllCategories ...
func GetAllCategories() ([]string, error) {

	var categories []string

	row, err := db.Query(`
		SELECT categoryname
		FROM categories
	`)
	defer row.Close()

	if err != nil {
		return categories, err
	}

	for row.Next() {
		var category string
		row.Scan(&category)
		categories = append(categories, category)
	}

	return categories, nil
}

// GetPostCategories ...
func getPostCategories(post *models.Post) error {

	var categories []string

	row, err := db.Query(`
		SELECT categories.categoryname
		FROM postcategories
		INNER JOIN categories
		ON postcategories.categoryid = categories.categoryid
		WHERE postcategories.postid = ?
	`, post.PostID)
	defer row.Close()

	if err != nil {
		return err
	}

	for row.Next() {
		var category string
		row.Scan(&category)
		categories = append(categories, category)
	}

	post.Categories = categories

	return nil
}

// InsertPostIntoCategories ...
func insertPostIntoCategories(postid int, categories []string) error {
	insertPostIntoCategory, err := db.Prepare(`
		INSERT INTO postcategories
		(categoryid, postid)
		SELECT categoryid, ?
		FROM categories
		WHERE categoryname = ?
	`)
	if err != nil {
		return err
	}
	for _, category := range categories {
		_, err = insertPostIntoCategory.Exec(postid, category)
		if err != nil {
			return err
		}
	}
	return nil
}
