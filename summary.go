package main

import (
	"fmt"
	"log"
	"os"

	"github.com/kniren/gota/dataframe"
)

func main() {
	// we open the csv file from the disk
	f, err := os.Open("./datasets/kc_house_data.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	df := dataframe.ReadCSV(f)

	fmt.Println(df.Describe())
}
