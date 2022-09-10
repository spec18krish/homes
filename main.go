package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"homes.co.nz/helper"
	model "homes.co.nz/property/model"
)

var allRecords []model.ValuationModel

func GetRecordFromArray(lineRecord []string) model.ValuationModel {
	record := model.ValuationModel{
		Id:            helper.ParseInt64(lineRecord[0]),
		StreetAddress: helper.GetString(lineRecord[1]),
		Town:          helper.GetString(lineRecord[2]),
		ValuationDate: helper.ParseDate(lineRecord[3]),
		Value:         helper.ParseFloat64(lineRecord[4]),
	}

	return record
}

func GetValuationsFromFile() {
	f, err := os.Open("properties.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		s := strings.Split(scanner.Text(), "	")
		if len(s) >= 5 {
			result := GetRecordFromArray(s)
			allRecords = append(allRecords, result)
		}
	}

	fmt.Println(allRecords)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	GetValuationsFromFile()
}
