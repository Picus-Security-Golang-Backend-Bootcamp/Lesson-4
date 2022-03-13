package city

import (
	"fmt"

	"gorm.io/gorm"
)

type City struct {
	gorm.Model
	Name        string
	Code        string
	CountryCode string `gorm:"type:varchar(100);column:CountryCode"`
}

// equals
/*type City struct {
	ID          uint `gorm:"primaryKey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Name        string
	Code        string
	CountryCode string
}*/

func (City) TableName() string {
	return "City"
}

func (c *City) ToString() string {
	return fmt.Sprintf("ID : %d, Name : %s, Code : %s, CountryCode : %s,CreatedAt : %s", c.ID, c.Name, c.Code, c.CountryCode, c.CreatedAt.Format("2006-01-02 15:04:05"))
}

func (c *City) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Printf("City (%s) deleting...", c.Name)
	return nil
}
