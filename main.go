package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Currency struct {
	Rate   string `json:"rate"`
	Status string `json:"status"`
}

type AllCurrency struct {
	Results []string `json:"supportedPairs"`
}

func main() {

	apiUrl := "https://www.freeforexapi.com/api/live"

	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error making HTTP request", err)
		return
	}

	defer response.Body.Close()

	var curr AllCurrency
	// var curren AllCurrency

	err = json.NewDecoder(response.Body).Decode(&curr)
	if err != nil {
		fmt.Println("Error decoding JSON", err)
		return
	}
	fmt.Println(curr.Results)

	f1 := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}

	f2 := func(w http.ResponseWriter, r *http.Request) {
		for _, str := range curr.Results {
			// fmt.Println(str)
			htm := str
			htmlStr := fmt.Sprintf(`<div class="card m-sm-3 border-danger m-4" style="width: 18rem;">
            <div class="card-body" >
                <h4 class="card-body">%s</h2>
                <p class="card-text">Some quick example text to build on the card title and make up the bulk of the card's content.</p>
            </div></div>`, htm)
			tmpl, _ := template.New("t").Parse(htmlStr)
			tmpl.Execute(w, nil)
		}
	}
	http.HandleFunc("/", f1)
	http.HandleFunc("action", f2)
	log.Fatal(http.ListenAndServe(":7000", nil))

}
