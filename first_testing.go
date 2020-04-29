package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	testRegression("sqft_living")
}

func testRegression(key string) {
	// we open the csv file from the disk
	f, err := os.Open("./datasets/testing.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// we create a new csv reader specifying
	// the number of columns it has
	salesData := csv.NewReader(f)
	salesData.FieldsPerRecord = 21

	// we read all the records
	records, err := salesData.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var columnId = getColumnId(key, records[0])
	if columnId < 0 {
		log.Fatal("\nInvalid Column name ", key)
	}

	// by slicing the records we skip the header
	records = records[1:]
	// Loop over the test data predicting y
	observed := make([]float64, len(records))
	predicted := make([]float64, len(records))
	var sumObserved float64
	for i, record := range records {
		// Parse the house price, "y".
		price, err := strconv.ParseFloat(records[i][2], 64)
		if err != nil {
			log.Fatal(err)
		}
		observed[i] = price
		sumObserved += price

		// Parse the grade value.
		grade, err := strconv.ParseFloat(record[columnId], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Predict y with our trained model.
		predicted[i] = predict(grade)
	}

	mean := sumObserved / float64(len(observed))
	var observedCoefficient, predictedCoefficient float64
	for i := 0; i < len(observed); i++ {
		observedCoefficient += math.Pow((observed[i] - mean), 2)
		predictedCoefficient += math.Pow((predicted[i] - mean), 2)
	}
	rsquared := predictedCoefficient / observedCoefficient

	// Output the R-squared to standard out.
	fmt.Printf("R-squared = %0.2f\n\n", rsquared)
}

func predict(key float64) float64 {
	return -42392.2729 + key*280.5608
}

func getColumnId(columnName string, tableHeader []string) int {
	for i, column := range tableHeader {
		if column == columnName {
			return i
		}
	}

	return -1
}
