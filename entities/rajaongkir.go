package entity

type RajaOngkir[T ResponseData] struct {
	Results T `json:"results"`
}
