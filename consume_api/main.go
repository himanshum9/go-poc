package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func GetData(w http.ResponseWriter, r *http.Request) {
	url := "https://pokeapi.co/api/v2/pokemon"
	data, err := http.Get(url)
	if err != nil {
		http.Error(w, "Faild To fetch data", http.StatusInternalServerError)
		return
	}
	defer data.Body.Close()
	if data.StatusCode == http.StatusOK {
		body, err := io.ReadAll(data.Body)
		if err != nil {
			http.Error(w, "Faild To fetch body", http.StatusInternalServerError)
			return
		}
		var prettyJSON map[string]interface{}
		if err := json.Unmarshal(body, &prettyJSON); err != nil {
			http.Error(w, "Json Decode failed", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(200)
		formatted, _ := json.MarshalIndent(prettyJSON, "", "  ")
		w.Write(formatted)
		// fmt.Println("Formatted Response:\n", string(formatted))
	}

}

func main() {

	http.HandleFunc("/getPokeData", GetData)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}

}
