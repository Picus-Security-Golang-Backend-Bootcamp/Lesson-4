package csv_helper

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/patika/file-processing/models"
)

func ReadCustomerWithWorkerPool(path string) error {

	jobs := make(chan []string, 5)
	results := make(chan models.Customer)

	wg := new(sync.WaitGroup)

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go convertToCustomerStruct(jobs, results, wg)
	}

	go func() {
		f, _ := os.Open(path)
		defer f.Close()
		lines, _ := csv.NewReader(f).ReadAll()
		isFirstRow := true
		for _, line := range lines {
			if isFirstRow {
				isFirstRow = false
				continue
			}

			jobs <- line
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()

		close(results)
	}()

	for v := range results {
		fmt.Println(v)
	}

	return nil
}

func convertToCustomerStruct(jobs <-chan []string, results chan<- models.Customer, wg *sync.WaitGroup) {
	defer wg.Done()
	// eventually I want to have a []string channel to work on a chunk of lines not just one line of text
	for j := range jobs {
		age, _ := strconv.Atoi(j[3])
		customer := models.Customer{
			Name:        j[0],
			PhoneNumber: j[1],
			Email:       j[2],
			Age:         age,
		}
		results <- customer
	}
}
