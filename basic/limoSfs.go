package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

const itemLevel string = "payload.logisticsOffer.programEligibilities"
const osnLevel string = "payload.logisticsOffer.offerShipNodes"

func main() {
	anotherChan := make(chan bool)
	offerId := "5DD75EF95CE4480E9629C22E7CAFA0AD"
	strUrl := "http://limo-nsf-app.prod.cp.prod.walmart.com/limo/services/v1/offers/" + offerId
	//url, _ := url.Parse(strUrl)
	fmt.Println(strUrl)
	go makeAsyncRestCall(strUrl, anotherChan)
	res := <-anotherChan
	if res == false {
		fmt.Println(offerId)
	}
	close(anotherChan)
}

func makeAsyncRestCall(url string, anotherChan chan bool) {
	client := &http.Client{}
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
		response, errBody := ioutil.ReadAll(res.Body)
		//fmt.Println(response)
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
			eligiblityName := gjson.Get(line.String(), "eligibilityName")
			eligiblityValue := gjson.Get(line.String(), "eligibilityValue")
			if eligiblityName.String() == "SFS" && eligiblityValue.Bool() {
				return true
			}
		}
	}
	return false
}

// programEligibilities
func osnParseJson(reponse []byte, err error, sfsPath string) bool {
	if err == nil {
		value := gjson.Get(string(reponse), sfsPath)
		fmt.Println(sfsPath)
		sfsCount := 0
		for _, line := range value.Array() {
			programEligibilities := gjson.Get(line.String(), "programEligibilities")
			for _, elli := range programEligibilities.Array() {
				eligiblityName := gjson.Get(elli.String(), "eligibilityName")
				eligiblityValue := gjson.Get(elli.String(), "eligibilityValue")
				if eligiblityName.String() == "SFS" && eligiblityValue.Bool() {
					//fmt.Println(eligiblityName, eligiblityValue)
					sfsCount += 1
				}
			}
		}
		fmt.Println(sfsCount)
		if sfsCount == 327 {
			return true
		}
	}
	return false
}
