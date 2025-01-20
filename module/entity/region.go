package entity

// Region represents a region entity
type Region struct {
    ID         int64  `json:"id"`
    Name       string `json:"name"`
    ParentID   int64  `json:"parent_id"`
    Level      int    `json:"level"`
    SortOrder  int    `json:"sort_order"`
    Status     int    `json:"status"`
    CreateTime int64  `json:"create_time"`
    UpdateTime int64  `json:"update_time"`
} 