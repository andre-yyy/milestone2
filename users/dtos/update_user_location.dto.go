package dto

type UpdateUserLocationDto struct {
	Address string `json:"address"`
	CityID  string `json:"city_id"`
}
