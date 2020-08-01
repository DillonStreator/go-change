package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"./change"
)

func main() {
	var paid int
	paidHelp := "amount that is paid"
	flag.IntVar(&paid, "paid", 0, paidHelp)
	flag.IntVar(&paid, "p", 0, paidHelp+" (shorthand)")

	var owed int
	owedHelp := "amount that is owed"
	flag.IntVar(&owed, "owed", 0, owedHelp)
	flag.IntVar(&owed, "o", 0, owedHelp+" (shorthand)")

	var drawerString string
	drawerHelp := "the drawer"
	flag.StringVar(&drawerString, "drawer", "", drawerHelp)
	flag.StringVar(&drawerString, "d", "", drawerHelp+" (shorthand)")

	var jsonInput string
	jsonInputHelp := "json input with keys paid, owed, drawer"
	flag.StringVar(&jsonInput, "json", "", jsonInputHelp)
	flag.StringVar(&jsonInput, "j", "", jsonInputHelp+" (shorthand)")

	var runAsAPI bool
	runAsAPIHelp := "run as api"
	flag.BoolVar(&runAsAPI, "api", false, runAsAPIHelp)
	flag.BoolVar(&runAsAPI, "a", false, runAsAPIHelp+" (shorthand)")

	flag.Parse()

	if runAsAPI {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != "POST" {
				fmt.Fprint(w, "send POST with json body containing keys `paid`, `owed`, `drawer`")
				return
			}

			var input change.Input
			err := json.NewDecoder(r.Body).Decode(&input)
			if err != nil {
				fmt.Fprint(w, err)
				return
			}

			result, err := change.Calculate(input)
			if err != nil {
				fmt.Fprint(w, err)
				return
			}
			bytes, err := json.Marshal(result)
			if err != nil {
				fmt.Fprint(w, err)
				return
			}

			fmt.Fprint(w, string(bytes))
		})
		port := os.Getenv("PORT")
		if len(port) == 0 {
			port = "8080"
		}
		http.ListenAndServe(fmt.Sprintf(":%s", port), nil)

	} else {
		if drawerString == "" && jsonInput == "" {
			panic("must provide drawer or json input")
		}

		var input change.Input

		if jsonInput != "" {
			err := json.Unmarshal([]byte(jsonInput), &input)
			if err != nil {
				fmt.Println("issue parsing json")
				panic(err)
			}
		} else {
			var drawer []int
			for _, denominationString := range strings.Split(drawerString, ",") {
				denomination, err := strconv.Atoi(denominationString)
				if err != nil {
					panic(fmt.Sprintf("issue parsing %s into integer", denominationString))
				}
				drawer = append(drawer, denomination)
			}

			input = change.Input{
				Owed:   owed,
				Paid:   paid,
				Drawer: drawer,
			}
		}

		result, err := change.Calculate(input)
		if err != nil {
			panic(err)
		}

		fmt.Println(result)
	}
}
