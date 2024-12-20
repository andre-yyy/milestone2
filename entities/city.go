package entity

type City struct {
	CityID     string `gorm:"type: varchar(100);primaryKey" json:"city_id,omitempty"`
	ProvinceID string `gorm:"type: varchar(100);" json:"province_id,omitempty"`
	Province   string `gorm:"type: varchar(225);" json:"province,omitempty"`
	Type       string `gorm:"type: varchar(225);" json:"type,omitempty"`
	CityName   string `gorm:"type: varchar(225);" json:"city_name,omitempty"`
	PostalCode string `gorm:"type: varchar(225);" json:"postal_code,omitempty"`
}
