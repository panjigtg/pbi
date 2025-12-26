package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"pbi/internal/pkg/models"
)

type AddressRepository interface {
	GetProvinces() ([]*models.Province, error)
	GetCities(provinceID string) ([]*models.City, error)
	GetProvinceByID(provinceID string) (*models.Province, error)
	GetCityByID(provinceID, cityID string) (*models.City, error)
}

type addressRepositoryImpl struct {
	baseURL string
}

func NewAddressRepository() *addressRepositoryImpl{
	return &addressRepositoryImpl{
		baseURL: "https://www.emsifa.com/api-wilayah-indonesia/api",
	}
}

func (ad *addressRepositoryImpl) GetProvinces() ([]*models.Province, error) {
	res, err := http.Get("https://www.emsifa.com/api-wilayah-indonesia/api/provinces.json")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var provinces []*models.Province
	if err := json.NewDecoder(res.Body).Decode(&provinces); err != nil {
		return nil, err
	}
	return provinces, nil
}


func (ad *addressRepositoryImpl) GetCities(provinceID string) ([]*models.City, error) {
	url := fmt.Sprintf("%s/regencies/%s.json", ad.baseURL, provinceID)

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var cities []*models.City
	if err := json.NewDecoder(res.Body).Decode(&cities); err != nil {
		return nil, err
	}

	return cities, nil
}

func (ad *addressRepositoryImpl) GetProvinceByID(provinceID string) (*models.Province, error) {
	provinces, err := ad.GetProvinces()
	if err != nil {
		return nil, err
	}

	for _, p := range provinces {
		if p.ID == provinceID {
			return p, nil
		}
	}
	return nil, errors.New("province not found")
}

func (ad *addressRepositoryImpl)GetCityByID(provinceID, cityID string) (*models.City, error) {
	cities, err := ad.GetCities(provinceID)
	if err != nil {
		return nil, err
	}

	for _, c := range cities {
		if c.ID == cityID {
			return c, nil
		}
	}
	return nil, errors.New("city not found")
}