package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Province struct {
		ID   string `json:"id"`
		Name string `json:"nama"`
	}
	City struct {
		ID   string `json:"id"`
		Name string `json:"nama"`
	}
)

func ValidateProvince(provinceID string) (bool, error) {
	res, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	var provinces []Province
	if err := json.NewDecoder(res.Body).Decode(&provinces); err != nil {
		return false, err
	}

	for _, p := range provinces {
		if p.ID == provinceID {
			return true, nil
		}
	}
	return false, nil
}

func ValidateCity(provinceID, cityID string) (bool, error) {
	url := fmt.Sprintf("https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%s.json", provinceID)
	res, err := http.Get(url)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	var cities []City
	if err := json.NewDecoder(res.Body).Decode(&cities); err != nil {
		return false, err
	}

	for _, c := range cities {
		if c.ID == cityID {
			return true, nil
		}
	}
	return false, nil
}