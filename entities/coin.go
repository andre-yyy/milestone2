package entity

type Coin struct {
	ID          int     `gorm:"type: serial;primaryKey;" json:"id,omitempty"`
	Name        string  `gorm:"type: varchar(225);not null;" json:"name,omitempty"`
	Description string  `gorm:"type: text;not null;" json:"description,omitempty"`
	Price       float64 `gorm:"type: decimal(10, 2);not null;" json:"price,omitempty"`
}
