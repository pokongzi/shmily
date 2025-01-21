package model

import (
	"common/mysql"
	"module/entity"

	"gorm.io/gorm"
)

// UserLikeModel handles database operations for UserLike
type UserLikeModel struct {
	DB *gorm.DB
}

// CreateUserLike creates a new user like record
func (m *UserLikeModel) CreateUserLike(like *entity.UserLike) error {
	return m.DB.Create(like).Error
}

// GetUserLike retrieves a user like by ID
func (m *UserLikeModel) GetUserLike(id int64) (*entity.UserLike, error) {
	var like entity.UserLike
	err := m.DB.First(&like, id).Error
	return &like, err
}

// UpdateUserLike updates an existing user like record
func (m *UserLikeModel) UpdateUserLike(like *entity.UserLike) error {
	return m.DB.Save(like).Error
}

// DeleteUserLike deletes a user like by ID
func (m *UserLikeModel) DeleteUserLike(id int64) error {
	return m.DB.Delete(&entity.UserLike{}, id).Error
}

// NewUserLikeModel creates a new instance of UserLikeModel
func NewUserLikeModel() *UserLikeModel {
	return &UserLikeModel{
		DB: mysql.DB,
	}
}
