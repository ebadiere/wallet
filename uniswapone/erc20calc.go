package uniswapone

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"os"
)

type Exchange struct {
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Id       string `json:"id"`
	MakerFee string `json:"maker_fee"`
	TakerFee string `json:"taker_fee"`
}

func LoadResponse() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	jsonFile, err := os.Open(os.Getenv("uniswapOneResp"))
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	exchanges := make(map[string]Exchange)
	json.Unmarshal([]byte(byteValue), &exchanges)

	count := 0
	for k, v := range exchanges {
		count++
		fmt.Println("token address:", k)
		fmt.Println("token Name:", v.Name)
		fmt.Println("token Symbol:", v.Symbol)
		fmt.Println("Exchange Address :", v.Id)
		fmt.Println("maker fee:", v.MakerFee)
		fmt.Println("taker fee:", v.TakerFee)
		fmt.Println("Exchange count: ", count)

	}
}
