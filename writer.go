package main

import (
	"html/template"
	"log"
	//"bytes"
	"os"
)

func writeResult(data []OptionData) {
	ts := template.Must(template.ParseFiles("ui/index.html"))
	output, err := os.Create("output/test2.html")
	if err != nil {
		log.Fatal(err)
	}

	//buf := new(bytes.Buffer)
	err = ts.Execute(output, data)
	if err != nil {
		log.Fatal(err)
	}
	//fileToWrite := "output/test.html"
	//err = os.WriteFile(fileToWrite, buf.Bytes(), 0644)
	//if err != nil {
	//	log.Fatal(err)
	//}
	output.Close()

}
