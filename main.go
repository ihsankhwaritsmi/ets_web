package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Currency struct {
	Rates map[string]struct {
		Rate   float64 `json:"rate"`
		Status float64 `json:"status"`
	} `json:"rates"`
}

type AllCurrency struct {
	Results []string `json:"supportedPairs"`
}

func main() {

	apiUrl := "https://www.freeforexapi.com/api/live?pairs=EURUSD,USDJPY,NZDUSD"

	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println("Error making HTTP request", err)
		return
	}

	defer response.Body.Close()

	// var curr Rates
	var curr Currency

	err = json.NewDecoder(response.Body).Decode(&curr)
	if err != nil {
		fmt.Println("Error decoding JSON", err)
		return
	}

	// fmt.Println(curr.Rates["EURUSD"].Rate)
	fmt.Printf("EUR/USD Rate: %.2f\n", curr.Rates["EURUSD"].Rate)
	fmt.Printf("EUR/USD Rate: %.2f\n", curr.Rates["USDJPY"].Rate)
	fmt.Printf("EUR/USD Rate: %.2f\n", curr.Rates["NZDUSD"].Rate)
	// fmt.Println(curren.Results)

	f := func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.Execute(w, nil)
	}

	f1 := func(w http.ResponseWriter, r *http.Request) {
		htmlStr := fmt.Sprintf(`
					<h4 class="card-body text-info">EURUSD</h2>
					<p class="card-text">%f</p>
				`, curr.Rates["EURUSD"].Rate)

		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)
	}
	f2 := func(w http.ResponseWriter, r *http.Request) {
		htmlStr := fmt.Sprintf(`
					<h4 class="card-body text-info">USDJPY</h2>
					<p class="card-text">%f</p>
				`, curr.Rates["USDJPY"].Rate)

		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)
	}
	f3 := func(w http.ResponseWriter, r *http.Request) {
		htmlStr := fmt.Sprintf(`
					<h4 class="card-body text-info">NZDUSD</h2>
					<p class="card-text">%f</p>
				`, curr.Rates["NZDUSD"].Rate)

		tmpl, _ := template.New("t").Parse(htmlStr)
		tmpl.Execute(w, nil)
	}

	http.HandleFunc("/", f)
	http.HandleFunc("/action1", f1)
	http.HandleFunc("/action2", f2)
	http.HandleFunc("/action3", f3)

	fs := http.FileServer(http.Dir("style.css"))
	http.Handle("/css/", http.StripPrefix("/css", fs))
	log.Fatal(http.ListenAndServe(":8000", nil))

}
