package entity

// User represents a user entity
type User struct {
	ID           int64  `json:"id"`
	UID          string `json:"uid"`
	Name         string `json:"name"`
	Birthday     string `json:"birthday"`
	Phone        string `json:"phone"`
	BirthdayType int    `json:"birthday_type"`
	Icon         string `json:"icon"`
	Sex          int    `json:"sex"`
	City         int    `json:"city"`
	Signature    string `json:"signature"`
	Status       int    `json:"status"`
	CreateTime   int64  `json:"create_time"`
	UpdateTime   int64  `json:"update_time"`
}

func (u *User) TableName() string {
	return "user"
}
