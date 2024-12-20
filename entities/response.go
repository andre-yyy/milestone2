package entity

type ResponseData interface {
	[]User | User |
		[]Category | Category |
		[]Product | Product |
		[]City | City |
		[]Coin | Coin |
		[]Topup | Topup |
		[]Order | Order |
		[]Activity | Activity |
		Invoice | Notification
}

type ServerResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ServerResponseData[T ResponseData] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

type LoginSuccess[T ResponseData] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
	Token   string `json:"token,omitempty"`
}

type TopupInvoice[T ResponseData] struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	Data    T       `json:"data"`
	Invoice Invoice `json:"invoice,omitempty"`
}

type OngkirResponse[T ResponseData] struct {
	RajaOngkir RajaOngkir[T] `json:"rajaongkir"`
}
