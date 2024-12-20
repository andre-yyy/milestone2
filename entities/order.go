package entity

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID            int       `gorm:"type: serial;primaryKey;" json:"id,omitempty"`
	UserID        string    `gorm:"type: uuid;not null;" json:"user_id,omitempty"`
	ProductID     int       `gorm:"type: int;not null;" json:"product_id,omitempty"`
	StartDate     time.Time `gorm:"type: timestamptz;not null;" json:"start_date,omitempty"`
	RentDays      int       `gorm:"type: int;not null;default: ;" json:"rent_days,omitempty"`
	EndDate       time.Time `gorm:"type: timestamptz;not null;" json:"end_date,omitempty"`
	TotalPrice    float64   `gorm:"type: decimal(10, 2);not null;" json:"total_price,omitempty"`
	DestinationID string    `gorm:"type: varchar(100);not null;" json:"destination_id,omitempty"`
	Status        string    `gorm:"type: varchar(100);default:'pending'" json:"status,omitempty"`
	CreatedAt     time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt     time.Time `gorm:"autoCreateTime;autoUpdateTime" json:"updated_at,omitempty"`
}

func (o *Order) AfterCreate(db *gorm.DB) error {
	var user User
	var product Product
	err := db.Where(&User{ID: o.UserID}).Take(&user).Error
	if err != nil {
		return err
	}
	err = db.Where(&Product{ID: o.ProductID}).Take(&product).Error
	if err != nil {
		return err
	}
	user.Deposit = user.Deposit - o.TotalPrice
	product.Stock = product.Stock - 1
	err = db.Save(&user).Error
	err = db.Save(&product).Error
	if err != nil {
		return err
	}
	activity := Activity{
		UserID:      o.UserID,
		Description: fmt.Sprintf("'%s' create new order with order id '%d'", user.Email, o.ID),
	}
	err = db.Create(&activity).Error
	if err != nil {
		return err
	}
	return nil
}

func (o *Order) AfterUpdate(db *gorm.DB) error {
	var user User
	err := db.Where(&User{ID: o.UserID}).Take(&user).Error
	if err != nil {
		return err
	}
	activity := Activity{
		UserID:      o.UserID,
		Description: fmt.Sprintf("'%s' changed order status to '%s' for order id '%d'", user.Email, o.Status, o.ID),
	}
	err = db.Create(&activity).Error
	if err != nil {
		return err
	}
	return nil
}
