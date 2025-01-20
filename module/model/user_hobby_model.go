package model

import (
    "gorm.io/gorm"
    "module/entity"
)

// UserHobbyModel handles database operations for UserHobby
type UserHobbyModel struct {
    DB *gorm.DB
}

// CreateUserHobby creates a new user hobby record
func (m *UserHobbyModel) CreateUserHobby(hobby *entity.UserHobby) error {
    return m.DB.Create(hobby).Error
}

// GetUserHobby retrieves a user hobby by ID
func (m *UserHobbyModel) GetUserHobby(id int64) (*entity.UserHobby, error) {
    var hobby entity.UserHobby
    err := m.DB.First(&hobby, id).Error
    return &hobby, err
}

// UpdateUserHobby updates an existing user hobby record
func (m *UserHobbyModel) UpdateUserHobby(hobby *entity.UserHobby) error {
    return m.DB.Save(hobby).Error
}

// DeleteUserHobby deletes a user hobby by ID
func (m *UserHobbyModel) DeleteUserHobby(id int64) error {
    return m.DB.Delete(&entity.UserHobby{}, id).Error
}

// NewUserHobbyModel creates a new instance of UserHobbyModel
func NewUserHobbyModel(db *gorm.DB) *UserHobbyModel {
    return &UserHobbyModel{
        DB: db,
    }
} 