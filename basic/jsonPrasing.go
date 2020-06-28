package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/tidwall/gjson"
)

// https://github.com/public-apis/public-apis
func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	if err == nil {
		responseStr, err := ioutil.ReadAll(response.Body)
		if err == nil {
			value := gjson.Get(string(responseStr), "descriptions")
			//fmt.Println(reflect.TypeOf(value))
			//fmt.Println(reflect.ValueOf(value))
			fmt.Println(value)
			//fmt.Println(string(responseStr))
		}
	}
}
