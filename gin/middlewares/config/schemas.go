package config


type RequestBody struct{
	Integer int64 `json:"integer" binding:"required"`
}