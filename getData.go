package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Pools struct {
	Data struct {
		Pools []struct {
			ID      string `json:"id"`
			TxCount string `json:"txCount"`
			FeesUSD string `json:"feesUSD"`
		} `json:"pools"`
	} `json:"data"`
}

func main() {
	fmt.Printf("%v, %T \n", getData(), getData())

}

func getData() *Pools {
	const api = "https://api.thegraph.com/subgraphs/name/ianlapham/uniswap-v3-subgraph"

	jsonData := map[string]string{
		"query": `
			{
				pools{
					id
					txCount
					feesUSD
				}
			}
		`,
	}

	jsonValue, err := json.Marshal(jsonData)
	if err != nil {
		panic(err)
	}

	request, err := http.NewRequest("POST", api, bytes.NewBuffer(jsonValue))
	if err != nil {
		panic(err)
	}

	//? Telling the server we're sending a json type request
	request.Header.Add("content-type", "application/json")

	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var pool Pools
	if err := json.Unmarshal(data, &pool); err != nil {
		fmt.Printf("err = %v\n", err)

	}

	return &pool

}
