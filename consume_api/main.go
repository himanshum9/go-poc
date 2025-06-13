package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetData() {
	url := "https://pokeapi.co/api/v2/pokemon"
	data, err := http.Get(url)
	if err != nil {
		fmt.Println("something is wrong in URL:", err)
		return
	}
	defer data.Body.Close()
	if data.StatusCode == http.StatusOK {
		body, err := io.ReadAll(data.Body)
		if err != nil {
			fmt.Println("faild To fetch body:", err)
			return
		}
		var prettyJSON map[string]any
		if err := json.Unmarshal(body, &prettyJSON); err != nil {
			fmt.Println("json Decode failed:", err)
			return
		}
		formatted, _ := json.MarshalIndent(prettyJSON, "", "  ")
		fmt.Println("Formatted Response:\n", string(formatted))
	}

}

func main() {

	GetData()

}
