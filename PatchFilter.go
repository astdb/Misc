package main

import (
	"log"
	"encoding/csv"
	"io"
	"os"
	"strings"
	"fmt"
	"io/ioutil"
	"sort"
)

func main() {
	f, err := os.Open("file.csv")
	if err != nil {
		log.Fatal(err)
	}

	// map to hold results
	res := map[string][]string{}
	resKeys := []string{}

	// create new reader and iterate through CSV rows
	r := csv.NewReader(f)
	for {
		record, err := r.Read()

		// quit if EOF
		if err == io.EOF {
			// return res
			break
		}

		if err != nil {
			panic(err)
		}

		if len(record) >= 2 {
			server := strings.TrimSpace(record[0])
			update := strings.TrimSpace(record[1])

			_, exists := res[server]

			if exists {
				// add to update list
				res[server] = append(res[server], update)
			} else {
				// new server
				res[server] = []string{update}
				resKeys = append(resKeys, server)
			}
		}
	}
	
	// create output string
	var outputStr strings.Builder

	sort.Strings(resKeys)

	for _, server := range resKeys {
		updates := res[server]
		outputStr.WriteString(fmt.Sprintf("\n - %s: ", server))

		for i, update := range updates {
			if i == 0 {
				outputStr.WriteString(fmt.Sprintf("%s", update))
			} else {
				outputStr.WriteString(fmt.Sprintf(", %s", update))
			}
		}

		outputStr.WriteString("\n")
	}

	err = ioutil.WriteFile("updates.txt", []byte(outputStr.String()), 0644)
	if err != nil {
		log.Println(err)
	}
}
