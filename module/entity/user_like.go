package entity

// UserLike represents a user like entity
type UserLike struct {
    ID         int64  `json:"id"`
    UID        int64  `json:"uid"`
    LikedUID   int64  `json:"liked_uid"`
    LikeType   int    `json:"like_type"`
    Status     int    `json:"status"`
    CreateTime int64  `json:"create_time"`
    UpdateTime int64  `json:"update_time"`
} 