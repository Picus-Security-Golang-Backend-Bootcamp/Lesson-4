package csv_helper

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/patika/file-processing/models"
)

func ReadCsv(filename string) ([]models.Customer, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	//reader.Comma = ';' // eğer virgül dışında farklı bir karakter ile verileri ayırıyorsanız
	lines, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var result []models.Customer

	for _, line := range lines[1:] {
		age, _ := strconv.Atoi(line[3])

		data := models.Customer{
			Name:        line[0],
			PhoneNumber: line[1],
			Email:       line[2],
			Age:         age,
		}

		result = append(result, data)
	}

	return result, nil
}
