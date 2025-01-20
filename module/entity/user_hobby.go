package entity

// UserHobby represents a user hobby entity
type UserHobby struct {
    ID          int64  `json:"id"`
    UID         int64  `json:"uid"`
    Name        string `json:"name"`
    HobbyType   int    `json:"hobby_type"`
    Description string `json:"description"`
    Status      int    `json:"status"`
    CreateTime  int64  `json:"create_time"`
    UpdateTime  int64  `json:"update_time"`
} 