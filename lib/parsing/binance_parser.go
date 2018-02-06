package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

var filename string

var filetype string

type Quote struct {
	Symbol      string `json:symbol`
	BidPrice    string `json:bidPrice`
	BidQuantity string `json:bidQty`
	AskPrice    string `json:askPrice`
	AskQuantity string `json:askQty`
}

func main() {

	flag.StringVar(&filename, "f", "", "filename of json file to load")
	flag.StringVar(&filetype, "t", "", "file type trades, quotes, or book")
	flag.Parse()

	if filename == "" {
		help()
		os.Exit(0)
	}

	if filetype == "" || (filetype != "trades" && filetype != "quotes" && filetype != "book") {
		help()
		os.Exit(0)
	}

	if filetype == "quotes" {
		quotes, err := read_json_quotes(filename)

		if err != nil {
			fmt.Printf("Error reading json file of type %s: %s\n", filetype, err)
		}

		fmt.Printf("%v\n", quotes)

	} else {

		data, err := read_json(filename)

		if err != nil {
			fmt.Printf("Error reading json file of type %s: %s\n", filetype, err)
		}

		fmt.Printf("JSON:\n%v\n", data)
	}

	os.Exit(0)

}

func read_json(filename string) (interface{}, error) {

	rawbytes, _ := ioutil.ReadFile(filename)

	var data interface{}

	err := json.Unmarshal(rawbytes, &data)

	return data, err

}

func read_json_quotes(filename string) ([]Quote, error) {

	rawbytes, _ := ioutil.ReadFile(filename)

	var data []Quote

	err := json.Unmarshal(rawbytes, &data)

	return data, err

}

func help() {
	fmt.Printf("parse sample json files from binance exchange\nusage binance_parser -f <filename> -t <[trades|quotes|book]>\n")
}

func print_quotes(data interface{}) {

	ia, ok := data.([]interface{})

	if ok {

		for _, v := range ia {

			im, ok := v.(map[string]interface{})

			if ok {
				fmt.Printf("%T: %v\n", im, im)
			}
		}

	}

}
