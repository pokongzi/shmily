package service

import (
	"errors"
	"shmily/module/entity"
	"time"

	"gorm.io/gorm"
)

func test() {

}

// 用户第三方登录
func ThirdPartyLogin(platform int8, code string) (*entity.User, error) {
	// 1. 获取平台用户信息
	oauthInfo, err := getPlatformUserInfo(platform, code)
	if err != nil {
		return nil, err
	}

	// 2. 查找是否已存在oauth记录
	existingOauth, err := userOauthModel.GetByPlatformAndOpenID(platform, oauthInfo.OpenID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if existingOauth != nil {
		// 3a. 已存在则更新信息
		existingOauth.LastLoginTime = time.Now().Unix()
		existingOauth.SessionKey = oauthInfo.SessionKey
		// ... 更新其他信息
		if err := userOauthModel.Update(existingOauth); err != nil {
			return nil, err
		}
		return userModel.GetByUID(existingOauth.UID)
	}

	// 3b. 不存在则创建新用户和oauth记录
	user := &entity.User{
		UID:  generateUID(),
		Name: oauthInfo.Nickname,
		// ... 设置其他用户信息
	}

	if err := userModel.Create(user); err != nil {
		return nil, err
	}

	oauth := &entity.UserOauth{
		UID:      user.UID,
		Platform: platform,
		OpenID:   oauthInfo.OpenID,
		// ... 设置其他oauth信息
	}

	if err := userOauthModel.Create(oauth); err != nil {
		return nil, err
	}

	return user, nil
}
