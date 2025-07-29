package database

import (
	"fmt"

	"github.com/rest_api_motivation/internal/models"
)

func (d *DB) GetAllPost() ([]models.Post, error) {
	const op = "database.GetAllPost"
	var posts []models.Post

	result := d.db.Find(&posts)
	if result.Error != nil {
		return []models.Post{}, fmt.Errorf("%s: %w", op, result.Error)
	}

	return posts, nil
}

func (d *DB) GetPost(user string) (models.Post, error) {
	const op = "database.GetPost"
	var post models.Post

	result := d.db.Where("user = ?", user).Find(&post)
	if result.Error != nil {
		return models.Post{}, fmt.Errorf("%s: %w", op, result.Error)
	}

	return post, nil
}

func (d *DB) CreatePost(post models.Post) (int, error) {
	const op = "database.CreatePost"

	result := d.db.Create(&post)
	if result.Error != nil {
		return 0, fmt.Errorf("%s: %w", op, result.Error)
	}

	return int(post.ID), nil
}

func (d *DB) UpdatePost(post models.Post) error {
	const op = "database.UpdatePost"

	var oldPost models.Post

	d.db.Where("user = ? AND email = ?", post.User, post.Email).First(&oldPost)

	oldPost.Category = post.Category
	oldPost.Text = post.Text

	result := d.db.Save(&oldPost)
	if result.Error != nil {
		return fmt.Errorf("%s: %w", op, result.Error)
	}

	return nil
}

func (d *DB) DeletePost(user string, email string) (models.Post, error) {

	const op = "database.DeletePost"
	var post models.Post
	result := d.db.Where("user = ? AND email = ?", user, email).Delete(&post)
	if result.Error != nil {
		return models.Post{}, fmt.Errorf("%s: %w", op, result.Error)
	}

	return post, nil
}
