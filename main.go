package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"
	"text/template"
)

type President struct {
	Number     uint   `json:"number"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	Party      string `json:"party"`
	TookOffice string `json:"tookOffice"`
	LeftOffice string `json:"leftOffice"`
}

func exitOnErr(err error) {
	if err != nil {
		log.Fatalf("error encountered exiting: %s", err.Error())
	}
}

func main() {
	// Read data into slice of struct
	presidentsJSONFile, err := os.Open(path.Join("data", "the-presidents.json"))
	exitOnErr(err)
	byteValue, _ := ioutil.ReadAll(presidentsJSONFile)
	var presidents []President
	json.Unmarshal(byteValue, &presidents)

	// Create new file
	file, err := os.Create("public/index.html")
	exitOnErr(err)
	defer file.Close()

	// apply the template to the vars map and write the result to file.
	tmpl, err := template.New("index.tmpl.html").Funcs(template.FuncMap{
		"last": func(x int, y int) bool {
			return x == y-1
		},
		"partyColor": func(party string) string {
			switch party {
			case "Republican":
				return "bg-red-500"
			case "Democratic":
				return "bg-blue-500"
			case "Democratic-Republican":
				return "bg-teal-500"
			case "Whig":
				return "bg-orange-500"
			case "Federalist":
				return "bg-purple-500"
			default:
				return "bg-slate-500"
			}
		},
	}).ParseFiles(path.Join("src", "templates", "index.tmpl.html"))
	exitOnErr(err)

	// Apply data to template
	err = tmpl.Execute(file, presidents)
	exitOnErr(err)

}
