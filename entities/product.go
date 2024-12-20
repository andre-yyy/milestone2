package entity

type Product struct {
	ID          int     `gorm:"type: serial;primaryKey;" json:"id,omitempty"`
	Name        string  `gorm:"type: varchar(225);not null;" json:"name,omitempty"`
	Stock       int     `gorm:"type: int;default:0;" json:"stock,omitempty"`
	RentalCost  float64 `gorm:"type: decimal(10, 2);not null;" json:"rental_cost,omitempty"`
	Daily       float64 `gorm:"type: decimal(10, 2);not null;" json:"daily,omitempty"`
	Weight      int     `gorm:"type: int;default:1;" json:"weight,omitempty"`
	CategoryID  int     `gorm:"type: int;not null;" json:"category_id,omitempty"`
	ImageUrl    string  `gorm:"type: text;unique;" json:"image_url,omitempty"`
	Description string  `gorm:"type: text;" json:"description,omitempty"`
}
