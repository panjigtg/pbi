package models


type Province struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type City struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

type ProvinceResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CityResponse struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}

