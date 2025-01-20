package entity

// UserPhoto represents a user photo entity
type UserPhoto struct {
    ID         int64  `json:"id"`
    UID        int64  `json:"uid"`
    PhotoURL   string `json:"photo_url"`
    SortOrder  int    `json:"sort_order"`
    PhotoType  int    `json:"photo_type"`
    Status     int    `json:"status"`
    CreateTime int64  `json:"create_time"`
    UpdateTime int64  `json:"update_time"`
} 