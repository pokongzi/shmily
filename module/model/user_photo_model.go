package model

import (
    "gorm.io/gorm"
    "module/entity"
    "common/mysql"
)

// UserPhotoModel handles database operations for UserPhoto
type UserPhotoModel struct {
    DB *gorm.DB
}

// CreateUserPhoto creates a new user photo record
func (m *UserPhotoModel) CreateUserPhoto(photo *entity.UserPhoto) error {
    return m.DB.Create(photo).Error
}

// GetUserPhoto retrieves a user photo by ID
func (m *UserPhotoModel) GetUserPhoto(id int64) (*entity.UserPhoto, error) {
    var photo entity.UserPhoto
    err := m.DB.First(&photo, id).Error
    return &photo, err
}

// UpdateUserPhoto updates an existing user photo record
func (m *UserPhotoModel) UpdateUserPhoto(photo *entity.UserPhoto) error {
    return m.DB.Save(photo).Error
}

// DeleteUserPhoto deletes a user photo by ID
func (m *UserPhotoModel) DeleteUserPhoto(id int64) error {
    return m.DB.Delete(&entity.UserPhoto{}, id).Error
}

// NewUserPhotoModel creates a new instance of UserPhotoModel
func NewUserPhotoModel() *UserPhotoModel {
    return &UserPhotoModel{
        DB: mysql.db,
    }
} 