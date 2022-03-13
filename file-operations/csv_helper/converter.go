package csv_helper

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/patika/file-processing/models"
)

func ConvertJSONToCSV(source, destination string) error {
	sourceFile, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	var customers []models.Customer
	err = json.NewDecoder(sourceFile).Decode(&customers)

	if err != nil {
		return err
	}

	output, err := os.Create(destination)
	if err != nil {
		return err
	}

	defer output.Close()

	csvWriter := csv.NewWriter(output)

	headers := []string{"Name", "PhoneNumber", "Email", "Age"}

	err = csvWriter.Write(headers)

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, c := range customers {
		var row []string
		row = append(row, c.Name, c.PhoneNumber, c.Email, fmt.Sprint(c.Age))
		err = csvWriter.Write(row)
		if err != nil {
			return err
		}
	}
	csvWriter.Flush()
	log.Println("JSON to CSV convertion completed...")

	return nil
}
