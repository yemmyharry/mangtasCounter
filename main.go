package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"sort"
)

type sorter struct {
	val int
	key string
}

type book struct {
	Content string	`json:"content"`
}

func counter(text string) []sorter {
	var s []sorter
	arrayText := regexp.MustCompile("[^0-9a-zA-Z]+").Split(text, -1)

	m := map[string]int{}
	for _, word := range arrayText {
		m[word]++
	}
	for k, v := range m {
		s = append(s, sorter{v, k})
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].val > s[j].val
	})
	total := len(s)
	if total > 10 {
		total = 10
	}
	fmt.Println()
	return s[:total]
}

func performPostJsonRequest(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}
	b := book{}
	json.Unmarshal(body, &b)

	res := counter(b.Content)
	var resultString = ""
	for _, v := range res {
		resultString += fmt.Sprintf("%s: %d\n", v.key, v.val)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(resultString))

}


func main() {

	http.HandleFunc("/hello", performPostJsonRequest)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
