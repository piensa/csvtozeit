package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func data() string {
	in := `first_name;last_name;username
"Rob";"Pike";rob
# lines beginning with a # character are ignored
Ken;Thompson;ken
"Robert";"Griesemer";"gri"
`
	return in
}

func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Reading the first line of a csv file</h1>")

	reader := csv.NewReader(strings.NewReader(data()))
	reader.Comma = ';'
	reader.Comment = '#'

	if reader != nil {
		fmt.Fprintf(w, "Got  file")
	}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		} else {

			fmt.Fprintf(w, "%s\n", record)
		}

	}
}
