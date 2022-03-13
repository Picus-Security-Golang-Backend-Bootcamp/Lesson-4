package main

import (
	"fmt"

	"github.com/mehmetcantas/location-service/domain/city"
	"github.com/mehmetcantas/location-service/domain/country"
	"github.com/mehmetcantas/location-service/infrastructure"
)

var (
	repository        *city.CityRepository
	countryRepository *country.CountryRepository
)

func init() {
	db := infrastructure.NewMySQLDB("root:Password123!@tcp(127.0.0.1:3306)/location?parseTime=True&loc=Local")
	repository = city.NewCityRepository(db)
	countryRepository = country.NewCountryRepository(db)
	repository.Migration()
	//repository.InsertSampleData()
	//countryRepository.Migration()
	//countryRepository.InsertSampleData()
}

func main() {

	//go get -u gorm.io/gorm

	cities := repository.FindByName("ad")

	//cities := repository.GetById(1324)
	//countries, _ := countryRepository.GetAllCountriesWithoutCityInformation()
	//cities.ToString()

	for _, c := range cities {
		fmt.Println(c.ToString())
		fmt.Println()
	}
}
