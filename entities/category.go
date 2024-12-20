package entity

type Category struct {
	ID   int    `gorm:"type: serial;primaryKey;" json:"id,omitempty"`
	Name string `gorm:"type: varchar(225);not null;unique;" json:"name,omitempty"`
}
