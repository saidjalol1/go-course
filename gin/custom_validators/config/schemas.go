package config


type Profile struct{
	Name string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Bio string `json:"bio"`
}