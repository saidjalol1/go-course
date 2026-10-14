package config


type LoginRequestBody struct{
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ProductCreate struct {
	Name  string 	`json:"name" binding:"required"`
	Price string 	`json:"price" binding:"required"`
}