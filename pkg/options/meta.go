package options

import (
	"database/sql"
	"time"
)

type ObjMeta struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`
}

type ListMeta struct {
	TotalCount int64 `json:"total_count"`
}
