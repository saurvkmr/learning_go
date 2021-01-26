package main

//https://play.golang.org/p/EYtsbc2niPl

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/tidwall/gjson"
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
	} else {
		fmt.Println(len(offers))
		calls(offers)
	}

}

func readFile() (record [][]string, err error) {
	file := "/Users/s0k02c9/Desktop/scripts/json_resources/SFS_Assortment_Latest_2.csv"
	f, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close() // this needs to be after the err check

	lines, err := csv.NewReader(f).ReadAll()
	return lines, err
}

func calls(offers [][]string) {
	var nonSfsList strings.Builder
	strUrl := ""
	anotherChan1 := make(chan bool)
	anotherChan2 := make(chan bool)
	anotherChan3 := make(chan bool)
	anotherChan4 := make(chan bool)
	anotherChan5 := make(chan bool)
	anotherChan6 := make(chan bool)
	anotherChan7 := make(chan bool)
	anotherChan8 := make(chan bool)

	file := "/Users/s0k02c9/Desktop/scripts/json_resources/NonSfs.csv"
	f, err := os.Create(file)

	if err != nil {
		fmt.Println(err)
	} else {
		startTime := time.Now()
		//for index, lines := range offers {
		for index := 1; index < len(offers); index += 8 {
			offerId1 := string(offers[index][0])
			offerId2 := string(offers[index+1][0])
			offerId3 := string(offers[index+2][0])
			offerId4 := string(offers[index+3][0])
			offerId5 := string(offers[index+4][0])
			offerId6 := string(offers[index+5][0])
			offerId7 := string(offers[index+6][0])
			offerId8 := string(offers[index+7][0])

			//fmt.Println(offerId1, offerId2, offerId3, offerId4, offerId5, offerId6, offerId7, offerId8)

			go makeAsyncRestCall(strUrl+offerId1, anotherChan1)
			go makeAsyncRestCall(strUrl+offerId2, anotherChan2)
			go makeAsyncRestCall(strUrl+offerId3, anotherChan3)
			go makeAsyncRestCall(strUrl+offerId4, anotherChan4)
			go makeAsyncRestCall(strUrl+offerId5, anotherChan5)
			go makeAsyncRestCall(strUrl+offerId6, anotherChan6)
			go makeAsyncRestCall(strUrl+offerId7, anotherChan7)
			go makeAsyncRestCall(strUrl+offerId8, anotherChan8)

			if false == <-anotherChan1 {
				nonSfsList.WriteString(offerId1 + "\n")
			}
			if false == <-anotherChan2 {
				nonSfsList.WriteString(offerId2 + "\n")
			}
			if false == <-anotherChan3 {
				nonSfsList.WriteString(offerId3 + "\n")
			}
			if false == <-anotherChan4 {
				nonSfsList.WriteString(offerId4 + "\n")
			}
			if false == <-anotherChan5 {
				nonSfsList.WriteString(offerId5 + "\n")
			}
			if false == <-anotherChan6 {
				nonSfsList.WriteString(offerId6 + "\n")
			}
			if false == <-anotherChan7 {
				nonSfsList.WriteString(offerId7 + "\n")
			}
			if false == <-anotherChan8 {
				nonSfsList.WriteString(offerId8 + "\n")
			}

			if (index-1)%240 == 0 {
				f.WriteString(nonSfsList.String())
				nonSfsList.Reset()
				f.Sync()
				fmt.Println("timeElasped :", time.Since(startTime).Seconds(), " index :", index)
			}
		}
	}
	close(anotherChan1)
	close(anotherChan2)
	close(anotherChan3)
	close(anotherChan4)
	close(anotherChan5)
	close(anotherChan6)
	close(anotherChan7)
	close(anotherChan8)

}

// create Transport instead of creating http client multiple times
	client := &http.Client{}
func makeAsyncRestCall(url string, anotherChan chan bool) {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("WM_CONSUMER.ID", "0b02d5f8-93d2-4e22-af26-3a47ecb857b0")
	req.Header.Set("WM_SVC.VERSION", "0.0.20")
	req.Header.Set("WM_SVC.ENV", "prod")
	req.Header.Set("WM_SVC.NAME", "limo-service")
	req.Header.Set("WM_BU_ID", "0")
	req.Header.Set("WM_QOS.CORRELATION_ID", "123")
	req.Header.Set("WM_CONSUMER.INTIMESTAMP", "1434667328446")
	req.Header.Set("WM_SEC.AUTH_SIGNATURE", "UrLb7h58h+cq6eLYD+u/d0ttW5l2OX+fV3vjiv1aEcLWfXppvBw+L2YTwGW0LzQ91oUsNe5zN44mhwn5qk1x8+KXDHN+jUbfMoE5IHJSl2ouPLQkSs7KnCuRSPBD1fo3vJQBe18Y3neak6XW4ITD7BNSwm7CeM1wmcsKJSdRc8g=")
	req.Header.Set("WM_SEC.KEY_VERSION", "2")
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Errors", err.Error())
		fmt.Println(err.Error())
		anotherChan <- false
	} else {
		defer res.Body.Close()
		response, errBody := ioutil.ReadAll(res.Body)
		if errBody == nil {
			if parseJson(response, errBody, itemLevelSfs) && osnParseJson(response, errBody, osnLevelSfs) {
				anotherChan <- true
			} else {
				anotherChan <- false
			}
		}
	}
}

func parseJson(reponse []byte, err error, sfsPath string) bool {
	if err == nil {
		value := gjson.Get(string(reponse), sfsPath)
		for _, line := range value.Array() {
			eligiblityName := gjson.Get(line.String(), eligibilityName)
			eligiblityValue := gjson.Get(line.String(), eligibilityValue)
			if eligiblityName.String() == "SFS" && eligiblityValue.Bool() {
				return true
			}
		}
	}
	return false
}

func osnParseJson(reponse []byte, err error, sfsPath string) bool {
	if err == nil {
		value := gjson.Get(string(reponse), sfsPath)
		//fmt.Println(sfsPath)
		sfsCount := 0
		for _, line := range value.Array() {
			programEligibilities := gjson.Get(line.String(), programEligibilities)
			for _, elli := range programEligibilities.Array() {
				eligiblityName := gjson.Get(elli.String(), eligibilityName)
				eligiblityValue := gjson.Get(elli.String(), eligibilityValue)
				if eligiblityName.String() == "SFS" && eligiblityValue.Bool() {
					//fmt.Println(eligiblityName, eligiblityValue)
					sfsCount += 1
				}
			}
		}
		//fmt.Println(sfsCount)
		if sfsCount == 327 {
			return true
		}
	}
	return false
}
