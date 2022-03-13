package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/patika/file-processing/csv_helper"
)

const fileName string = "sample.txt"

func main() {
	//CreateEmptyFile()
	// f, err := os.OpenFile(fileName, os.O_RDWR, 0755) // READONLY
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// WriteFile(f)
	//GetFileInfo()
	//ReadFileLines()
	//ReadFileWords()

	// customer, _ := csv_helper.ReadCsv("customer.csv")

	// fmt.Println(customer)

	// err := csv_helper.ConvertJSONToCSV("customer.json", "customer_output.csv")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	csv_helper.ReadCustomerWithWorkerPool("customer.csv")
	fmt.Scanln()

	// -------------------- KAYNAKLAR -------------------------
	// https://medium.com/swlh/processing-16gb-file-in-seconds-go-lang-3982c235dfa2
	// https://medium.com/@timhigins/parallelizing-golang-file-io-7c5d0cacfe5a
	// http://mihirkelkar.ca/csvapione.html
	// https://medium.com/golicious/comparing-ioutil-readfile-and-bufio-scanner-ddd8d6f18463
	// https://medium.com/golangspec/introduction-to-bufio-package-in-golang-ad7d1877f762
}

func CreateEmptyFile() *os.File {
	// info, err := os.Stat(fileName)
	// if os.IsNotExist(err) == false {
	// 	log.Fatal(info.Name())
	// }
	myFile, err := os.Create(fileName)

	if err != nil {
		log.Fatal("ERROR! ", err)
	}
	log.Println("Empty file created successfully. ", myFile)

	return myFile
}

func GetFileInfo() {
	fileInfo, err := os.Stat(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("File Name:", fileInfo.Name())
	fmt.Println("Size ", fileInfo.Size(), " bytes")
	fmt.Println("Permissions:", fileInfo.Mode())
	fmt.Println("Last modified:", fileInfo.ModTime())
	fmt.Println("Is Directory: ", fileInfo.IsDir())
}

func ReadFileLines() {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		// do something with a line
		fmt.Printf("line: %s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func ReadFileWords() {
	// open file
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file word by word using scanner
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		// do something with a word
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func WriteFile(file *os.File) {
	defer file.Close()
	w := bufio.NewWriter(file)
	w.WriteString("line 1\n")
	w.WriteString("line 2\n")
	w.WriteString("line 3\n")
	w.WriteString("line 4\n")
	w.Flush()

}
