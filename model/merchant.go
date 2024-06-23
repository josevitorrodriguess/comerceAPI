package model

type Merchant struct {
	ID          int     `json:"id_merchant"`
	Name        string  `json:"name"`
	TypeProduct string `json:"product_type"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Products []Product  `json:products`
}
