package main

//https://play.golang.org/p/EYtsbc2niPl

import (
	"encoding/csv"
	"fmt"
	"os"
)

const itemLevelSfs string = "payload.logisticsOffer.programEligibilities"
const osnLevelSfs string = "payload.logisticsOffer.offerShipNodes"

const eligibilityName string = "eligibilityName"
const eligibilityValue string = "eligibilityValue"
const programEligibilities string = "programEligibilities"
const osnType string = "offerShipNodeType"

func main() {
	offers, err := readFile()

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(len(offers))
	calls(offers)
}

func readFile() (record [][]string, err error) {
	file := "/Users/s0k02c9/Desktop/scripts/json_resources/SFS_Assortment_Latest_2.csv"
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close() // this needs to be after the err check

	lines, err := csv.NewReader(f).ReadAll()
	return lines, err
}

func calls(offers [][]string) {
	//client := &http.Client{}
	//fmt.Println(client)
	checkIndexs(len(offers), 10)
}

func checkIndexs(num, n int) {
	checkIndexForSlice.Split(num, n)
}
