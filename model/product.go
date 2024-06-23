package model


type Product struct {
	ID 			int 	`json:"id_product"`
	Name 		string	`json:"name"`
	Price 		float64	`json:"price"`
	Description string  `json:"description"`
	MerchantID  int     `json:"merchant_id"`
}

type ProductOutput struct {
	ID          int     `json:"id_product"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Supplier    string  `json:"supplier"`
}

func ConvertToProductOutput(input Product, supplierName string) ProductOutput {
	return ProductOutput{
		Name:        input.Name,
		Price:       input.Price,
		Description: input.Description,
		Supplier:    supplierName,
	}
}