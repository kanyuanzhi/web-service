package model

type DefaultFields struct {
	ID        uint  `json:"id" gorm:"primary_key"`
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime:milli"` // gorm自动使用当前时间戳的秒数填充
	UpdatedAt int   `json:"updated_at" gorm:"autoUpdateTime:milli"`
}
