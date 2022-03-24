package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Stat struct {
	Date       time.Time
	TotalCases float64
	NewCases   float64
}

func main() {
	file, err := os.Open("owid-covid-data.csv")
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(file)

	myData := make(map[string][]Stat)

	_, _ = csvReader.Read()
	for {
		line, err := csvReader.Read()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			panic(err)
		}

		date, err := time.Parse("2006-01-02", line[3])
		if err != nil {
			panic(err)
		}

		totalCases := float64(0)
		if line[4] != "" {
			totalCases, err = strconv.ParseFloat(line[4], 16)
			if err != nil {
				panic(err)
			}
		}

		newCases := float64(0)
		if line[5] != "" {
			newCases, err = strconv.ParseFloat(line[5], 16)
			if err != nil {
				panic(err)
			}
		}

		myData[line[2]] = append(myData[line[2]], Stat{
			Date:       date,
			TotalCases: totalCases,
			NewCases:   newCases,
		})
	}

	fmt.Println(len(myData))
}
