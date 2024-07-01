package model

type Merchant struct {
	ID          int     `json:"id_merchant"`
	Name        string  `json:"name"`
	CNPJ        string  `json:"cnpj"`
	TypeProduct string  `json:"product_type"`
	Email       string  `json:"email"`
	Password    string  `json:"password"`
	Products []string   `json:products`
}
