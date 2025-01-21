package entity

// UserOauth represents third-party platform user information
type UserOauth struct {
	ID              int64  `json:"id"`
	UID             string `json:"uid"`
	Platform        int8   `json:"platform"`
	OpenID          string `json:"open_id"`
	UnionID         string `json:"union_id"`
	Nickname        string `json:"nickname"`
	Avatar          string `json:"avatar"`
	Gender          int8   `json:"gender"`
	SessionKey      string `json:"session_key"`
	AccessToken     string `json:"access_token"`
	RefreshToken    string `json:"refresh_token"`
	TokenExpireTime int64  `json:"token_expire_time"`
	LastLoginTime   int64  `json:"last_login_time"`
	Status          int8   `json:"status"`
	CreateTime      int64  `json:"create_time"`
	UpdateTime      int64  `json:"update_time"`
}

func (uo *UserOauth) TableName() string {
	return "user_oauth"
}
