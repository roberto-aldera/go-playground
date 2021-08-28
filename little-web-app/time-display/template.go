package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type PageVariables struct {
	Date string
	Time string
}

func main() {
	http.HandleFunc("/", HomePage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func HomePage(w http.ResponseWriter, r *http.Request) {

	now := time.Now()
	HomePageVars := PageVariables{
		// Go has weird formatting... try remember Jan 02 15:04:05 2006 MST (1234567, get it? No.)
		Date: now.Format("02 Jan 2006"),
		Time: now.Format("15:04:05"),
	}

	// Parse the html file homepage.html
	t, err := template.ParseFiles("homepage.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	// Execute the template and pass it the HomePageVars struct to fill in the gaps
	err = t.Execute(w, HomePageVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}
}
