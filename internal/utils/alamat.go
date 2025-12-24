package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

// utils/emsifa.go
type EmsifaProvince struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type EmsifaCity struct {
	ID         string `json:"id"`
	ProvinceID string `json:"province_id"`
	Name       string `json:"name"`
}


func ValidateProvince(provinceID string) (bool, error) {
	res, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	var provinces []EmsifaProvince
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

	var cities []EmsifaCity
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

func GetProvince(provinceID string) (*EmsifaProvince, error) {
	res, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var provinces []EmsifaProvince
	if err := json.Unmarshal(body, &provinces); err != nil {
		return nil, err
	}

	for _, p := range provinces {
		if p.ID == provinceID {
			return &p, nil
		}
	}
	return nil, errors.New("province not found")
}


func GetCity(provinceID, cityID string) (*EmsifaCity, error) {
	url := fmt.Sprintf(
		"https://www.emsifa.com/api-wilayah-indonesia/api/regencies/%s.json",
		provinceID,
	)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)

	var cities []EmsifaCity
	if err := json.Unmarshal(body, &cities); err != nil {
		return nil, err
	}

	for _, c := range cities {
		if c.ID == cityID {
			return &c, nil
		}
	}
	return nil, errors.New("city not found")
}

