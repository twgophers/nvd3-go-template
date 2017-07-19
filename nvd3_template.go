package main

import (
	"html/template"
	"os"

	"github.com/rpinheiroalmeida/browser"
)

var t *template.Template

type Data_type struct {
	X int
	Y float64
}

var data []Data_type

func init() {
	t = template.Must(template.ParseFiles("templates/index.html",
		"templates/js.tmpl"))
	data = []Data_type{
		Data_type{X: 1950, Y: 300.2},
		Data_type{X: 1950, Y: 300.2},
		Data_type{X: 1960, Y: 543.3},
		Data_type{X: 1970, Y: 1075.9},
		Data_type{X: 1980, Y: 2862.5},
		Data_type{X: 1990, Y: 5979.6},
		Data_type{X: 2000, Y: 10289.7},
		Data_type{X: 2010, Y: 14958.3},
	}
}

func render(path string) {
	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	t.ExecuteTemplate(f, "index.html", data)
}

func main() {
	path := os.Args[1]
	render(path)
	browser.Open(path)
}
