package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Topup struct {
	ID         int       `gorm:"type: serial;primaryKey;" json:"id,omitempty"`
	ExternalID string    `gorm:"type: varchar(225);not null;" json:"external_id,omitempty"`
	UserID     string    `gorm:"type: uuid;not null;" json:"user_id,omitempty"`
	Status     string    `gorm:"type: varchar(225);default: ;" json:"status,omitempty"`
	CreatedAt  time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
}

func (t *Topup) AfterCreate(db *gorm.DB) error {
	var user User
	err := db.Where(&User{ID: t.UserID}).Take(&user).Error
	if err != nil {
		return err
	}
	activity := Activity{
		UserID:      t.UserID,
		Description: fmt.Sprintf("'%s' create new topup with topup id '%d'", user.Email, t.ID),
	}
	err = db.Create(&activity).Error
	if err != nil {
		return err
	}
	return nil
}

func (t *Topup) AfterUpdate(db *gorm.DB) error {
	var user User
	err := db.Where(&User{ID: t.UserID}).Take(&user).Error
	if err != nil {
		return err
	}
	activity := Activity{
		UserID:      t.UserID,
		Description: fmt.Sprintf("'%s' changed topup status to '%s' for topup id '%d'", user.Email, t.Status, t.ID),
	}
	err = db.Create(&activity).Error
	if err != nil {
		return err
	}
	return nil
}
