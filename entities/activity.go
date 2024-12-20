package entity

import "time"

type Activity struct {
	ID          int        `gorm:"type: serial;primaryKey;autoIncrement;" json:"id,omitempty"`
	UserID      string     `gorm:"type: uuid;not null;" json:"user_id,omitempty"`
	Description string     `gorm:"type: text;not null;" json:"description,omitempty"`
	CreatedAt   *time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
}
