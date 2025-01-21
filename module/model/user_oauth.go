package model

import (
	"common/mysql"
	"module/entity"

	"gorm.io/gorm"
)

// UserOauthModel handles database operations for user_oauth table
type UserOauthModel struct {
	DB *gorm.DB
}

// Create creates a new user oauth record
func (m *UserOauthModel) Create(oauth *entity.UserOauth) error {
	return m.DB.Create(oauth).Error
}

// Update updates an existing user oauth record
func (m *UserOauthModel) Update(oauth *entity.UserOauth) error {
	return m.DB.Save(oauth).Error
}

// GetByID retrieves a user oauth record by ID
func (m *UserOauthModel) GetByID(id int64) (*entity.UserOauth, error) {
	var oauth entity.UserOauth
	err := m.DB.First(&oauth, id).Error
	if err != nil {
		return nil, err
	}
	return &oauth, nil
}

// GetByUID retrieves all oauth records for a user
func (m *UserOauthModel) GetByUID(uid string) ([]*entity.UserOauth, error) {
	var oauths []*entity.UserOauth
	err := m.DB.Where("uid = ?", uid).Find(&oauths).Error
	return oauths, err
}

// GetByPlatformAndOpenID retrieves a user oauth record by platform and openID
func (m *UserOauthModel) GetByPlatformAndOpenID(platform int8, openID string) (*entity.UserOauth, error) {
	var oauth entity.UserOauth
	err := m.DB.Where("platform = ? AND open_id = ?", platform, openID).First(&oauth).Error
	if err != nil {
		return nil, err
	}
	return &oauth, nil
}

// GetByUnionID retrieves user oauth records by unionID
func (m *UserOauthModel) GetByUnionID(unionID string) ([]*entity.UserOauth, error) {
	var oauths []*entity.UserOauth
	err := m.DB.Where("union_id = ?", unionID).Find(&oauths).Error
	return oauths, err
}

// UpdateLastLoginTime updates the last login time for a user oauth record
func (m *UserOauthModel) UpdateLastLoginTime(id int64, lastLoginTime int64) error {
	return m.DB.Model(&entity.UserOauth{}).
		Where("id = ?", id).
		Update("last_login_time", lastLoginTime).Error
}

// UpdateSessionKey updates the session key for a user oauth record
func (m *UserOauthModel) UpdateSessionKey(id int64, sessionKey string) error {
	return m.DB.Model(&entity.UserOauth{}).
		Where("id = ?", id).
		Update("session_key", sessionKey).Error
}

// UpdateTokenInfo updates token related information
func (m *UserOauthModel) UpdateTokenInfo(id int64, accessToken, refreshToken string, expireTime int64) error {
	return m.DB.Model(&entity.UserOauth{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"access_token":      accessToken,
			"refresh_token":     refreshToken,
			"token_expire_time": expireTime,
		}).Error
}

// NewUserOauthModel creates a new instance of UserOauthModel
func NewUserOauthModel() *UserOauthModel {
	return &UserOauthModel{
		DB: mysql.DB,
	}
}
