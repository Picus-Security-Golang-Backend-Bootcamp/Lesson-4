package country

import "gorm.io/gorm"

type CountryRepository struct {
	db *gorm.DB
}

func NewCountryRepository(db *gorm.DB) *CountryRepository {
	return &CountryRepository{
		db: db,
	}
}

func (r *CountryRepository) GetAllCountriesWithCityInformation() ([]Country, error) {
	var countries []Country
	result := r.db.Preload("Cities").Find(&countries)

	if result.Error != nil {
		return nil, result.Error
	}

	return countries, nil
}

func (r *CountryRepository) GetCountryWithName(name string) (*Country, error) {
	var country *Country
	result := r.db.Where(Country{Name: name}).Attrs(Country{Code: "NULL", Name: "NULL"}).FirstOrInit(&country) // Eğer sorgu sonucunda veri bulunursa Attrs kısmında yazılanlar ignore edilir.

	if result.Error != nil {
		return nil, result.Error
	}

	return country, nil
}

func (r *CountryRepository) GetCountryWithNameOrCreate(name string) (*Country, error) {
	var country *Country
	result := r.db.Where(Country{Name: name}).Attrs(Country{Code: "NULL", Name: name}).FirstOrCreate(&country) // Eğer sorgu sonucunda veri bulunursa veri oluşturulmaz.

	if result.Error != nil {
		return nil, result.Error
	}

	return country, nil
}

func (r *CountryRepository) GetAllCountriesWithoutCityInformation() ([]Country, error) {
	var countries []Country
	result := r.db.Find(&countries)

	if result.Error != nil {
		return nil, result.Error
	}

	return countries, nil
}

func (r *CountryRepository) Migration() {
	r.db.AutoMigrate(&Country{})
}

func (r *CountryRepository) InsertSampleData() {
	cities := []Country{
		{Name: "Türkiye", Code: "TR"},
		{Name: "Amerika", Code: "US"},
	}

	for _, c := range cities {
		r.db.Create(&c)
	}
}
