package main

import (
	"html/template"
	"log"
	"os"
	"strings"
)

func writeResult(data []OptionData, fn string) {
	ts := template.Must(template.ParseFiles("ui/index.html"))

	name := strings.TrimSuffix(fn, ".json")
	html_fn := name + ".html"
	output, err := os.Create("output/" + html_fn)
	if err != nil {
		log.Fatal(err)
	}

	err = ts.Execute(output, data)
	if err != nil {
		log.Fatal(err)
	}
	output.Close()
}
