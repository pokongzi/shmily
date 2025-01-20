package entity

// Relation represents a user relation entity
type Relation struct {
    ID           int64  `json:"id"`
    RequesterUID string `json:"requester_uid"`
    RequestedUID string `json:"requested_uid"`
    Status       int    `json:"status"`
    RequestTime  int64  `json:"request_time"`
    ResponseTime int64  `json:"response_time"`
    StartTime    int64  `json:"start_time"`
    EndTime      int64  `json:"end_time"`
    RequestMsg   string `json:"request_msg"`
    RejectReason string `json:"reject_reason"`
    CreateTime   int64  `json:"create_time"`
    UpdateTime   int64  `json:"update_time"`
} 