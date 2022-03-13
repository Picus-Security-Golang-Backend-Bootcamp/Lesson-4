package country

import (
	"github.com/mehmetcantas/location-service/domain/city"
	"gorm.io/gorm"
)

type Country struct {
	gorm.Model
	Name   string
	Code   string
	Cities []city.City `gorm:"foreignKey:CountryCode;references:Code"`
}
