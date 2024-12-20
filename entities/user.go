package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        string    `gorm:"type: uuid;primaryKey;default:gen_random_uuid()" json:"id,omitempty"`
	FullName  string    `gorm:"type: varchar(225);not null;" json:"full_name,omitempty"`
	Role      string    `gorm:"type: varchar(100);default:'user';" json:"role,omitempty"`
	Email     string    `gorm:"type: varchar(225);not null;unique;" json:"email,omitempty"`
	Password  string    `gorm:"type: varchar(225);not null;" json:"-"`
	Address   string    `gorm:"type: varchar(225);" json:"address,omitempty"`
	CityID    string    `gorm:"type: varchar(100);" json:"city_id,omitempty"`
	Deposit   float64   `gorm:"type: decimal(10, 2);default:0.0;" json:"deposit"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updated_at,omitempty"`
}

/*

func (u *User) AfterUpdate(db *gorm.DB) error {
	activity := Activity{
		UserID:      u.ID,
		Description: fmt.Sprintf("'%s' update address '%s' and city id '%s'", u.Email, u.Address, u.CityID),
	}
	err := db.Create(&activity).Error
	if err != nil {
		return err
	}
	return nil
}

*/

func (u *User) AfterUpdate(db *gorm.DB) error {
	activity := Activity{
		UserID:      u.ID,
		Description: fmt.Sprintf("'%s' - coins : '%v'", u.Email, u.Deposit),
	}
	err := db.Create(&activity).Error
	if err != nil {
		return err
	}
	return nil
}
