package model

type Client struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	CPF      string `json:"cpf"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
