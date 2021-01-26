package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type SolrResponse struct {
	ResponseHeader struct {
		ZkConnected bool `json:"zkConnected"`
		Status      int  `json:"status"`
		QTime       int  `json:"QTime"`
		Params      struct {
			Q    string `json:"q"`
			Fl   string `json:"fl"`
			Rows string `json:"rows"`
		} `json:"params"`
	} `json:"responseHeader"`
	Response struct {
		NumFound int `json:"numFound"`
		Start    int `json:"start"`
		Docs     []struct {
			OfferID      string `json:"offerId"`
			LegacyItemID string `json:"legacyItemId"`
		} `json:"docs"`
	} `json:"response"`
}

const solrQueryString string = "http://solrcloud.prod-az-westus.limo-items-prod.ms-df-solrcloud.prod.us.walmart.net:8983/solr/limo_offer_setup/query?fl=offerId,%20legacyItemId&q=isSFSEligible:true&rows="

func main() {

	sfsCount := make(chan int)
	go getRecordCount(sfsCount)
	getResponse(<-sfsCount)
	close(sfsCount)

}

func getRecordCount(sfsCount chan int) {
	url := solrQueryString + "0"
	responseBody, responseBodyErr := getResponseByte(url)
	if responseBodyErr != nil {
		sfsCount <- 0
	}

	var countResponse SolrResponse
	jsonParseError := json.Unmarshal(responseBody, &countResponse)
	if jsonParseError != nil {
		fmt.Println(jsonParseError.Error())
	} else {
		sfsCount <- countResponse.Response.NumFound
	}
}

func getResponseByte(url string) ([]byte, error) {
	httpResponse, err := http.Get(url)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		defer httpResponse.Body.Close()
		responseBody, responseBodyErr := ioutil.ReadAll(httpResponse.Body)
		return responseBody, responseBodyErr
	}
	return nil, err
}

func getResponse(sfsCount int) bool {

	url := solrQueryString + strconv.Itoa(sfsCount)
	responseBody, responseBodyErr := getResponseByte(url)
	if responseBodyErr != nil {
		return false
	}
	var offerIds SolrResponse
	jsonParseError := json.Unmarshal(responseBody, &offerIds)

	if jsonParseError != nil {
		return false
	}
	return writeToFile(offerIds)
}

func writeToFile(offerIds SolrResponse) bool {
	var offerListSB strings.Builder
	file := "/Users/s0k02c9/Desktop/scripts/SfsTrue.csv"
	f, err := os.Create(file)
	if err != nil {
		return false
	}
	defer f.Close()

	writeDone := make(chan bool)
	offerListSB.WriteString("offer,itemId\n")

	for index, line := range offerIds.Response.Docs {
		offerListSB.WriteString(line.OfferID + "," + line.LegacyItemID + "\n")
		if (index+1)%42337 == 0 {
			fmt.Println("inside write")
			go writeOffersToFile(offerListSB.String(), writeDone, f)
			offerListSB.Reset()
		}
	}
	fmt.Println("loop done")
	fmt.Println(<-writeDone)
	close(writeDone)
	return true
}

func writeOffersToFile(offerList string, writeDone chan bool, file *os.File) {
	file.WriteString(offerList)
	file.Sync()
	writeDone <- true
}
