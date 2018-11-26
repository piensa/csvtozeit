package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	fileName := "data.csv"
	f, _ := os.Open(fileName)
	reader := csv.NewReader(f)
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Fprintf(w, "%s\n", record)
	}
}
