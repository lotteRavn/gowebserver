package main

import (
    "io"
    "fmt"
    "log"
    "net/http"
    "strconv"
)

func hello(w http.ResponseWriter, r *http.Request) {
    io.WriteString(w, "Hello world!")
}

func to_roman(input int)  string {
	var s string

	if input <1 || input > 3999 {
		return "Invalid Roman Number Value"}

	s = ""
	for input >= 1000 {
		s += "M"
		input -= 1000
	}
	for input >= 900 {
		s += "CM"
		input -= 900
	}
	for input >= 500 {
		s += "D"
		input -= 500
	}
	for input >= 400 {
		s += "CD"
		input -= 400
	}
	for input >= 100 {
		s += "C"
		input -= 100
	}
	for input >= 90 {
		s += "XC"
		input -= 90
	}
	for input >= 50 {
		s += "L"
		input -= 50
	}
	for input >= 40 {
		s += "XL"
		input -= 40
	}
	for input >= 10 {
		s += "X"
		input -= 10
	}
	for input >= 9 {
		s += "IX"
		input -= 9
	}
	for input >= 5 {
		s += "V"
		input -= 5
	}
	for input >= 4 {
		s += "IV"
		input -= 4	
	}
	for input >= 1 {
		s += "I"
		input -= 1
	}
	return s
}

type romanGenerator int
func (n romanGenerator) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    ascii_num := r.URL.Path[7:]
    i, err := strconv.Atoi(ascii_num)
    if err != nil {
        log.Print(err)
    }
    fmt.Fprintf(w, "Here's your number: %s\n", to_roman(i))
}



func main() {
    h := http.NewServeMux()

    h.Handle("/roman/", romanGenerator(1))
    h.HandleFunc("/", hello)

    err := http.ListenAndServe(":8000", h)
    log.Fatal(err)
}
