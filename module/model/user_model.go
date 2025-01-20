package model

import (
    "gorm.io/gorm"
    "module/entity"
)

// UserModel handles database operations for User
// ... existing code ...
type UserModel struct {
    DB *gorm.DB
}

// CreateUser creates a new user record
func (m *UserModel) CreateUser(user *entity.User) error {
    return m.DB.Create(user).Error
}

// GetUser retrieves a user by ID
func (m *UserModel) GetUser(id int64) (*entity.User, error) {
    var user entity.User
    err := m.DB.First(&user, id).Error
    return &user, err
}

// UpdateUser updates an existing user record
func (m *UserModel) UpdateUser(user *entity.User) error {
    return m.DB.Save(user).Error
}

// DeleteUser deletes a user by ID
func (m *UserModel) DeleteUser(id int64) error {
    return m.DB.Delete(&entity.User{}, id).Error
}

// NewUserModel creates a new instance of UserModel
func NewUserModel(db *gorm.DB) *UserModel {
    return &UserModel{
        DB: db,
    }
} 