package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Quote struct {
	Author string `json:"author"`
	Text   string `json:"text"`
}

type AuthorToQuoteMap map[string][]string

func parseQuotesToMap(quotes *[]Quote) AuthorToQuoteMap {
	mapToReturn := make(AuthorToQuoteMap)
	for _, quote := range *quotes {
		mapToReturn[quote.Author] = append(mapToReturn[quote.Author], quote.Text)
	}
	return mapToReturn
}

func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseStringArray(quoteArray []string) []string {
	reversedStringArray := new([]string)

	for _, val := range quoteArray {
		*reversedStringArray = append(*reversedStringArray, ReverseString(val))
	}
	return *reversedStringArray
}

func reverseQuotes(authorToQuoteMap AuthorToQuoteMap) {
	for key, value := range authorToQuoteMap {
		authorToQuoteMap[key] = reverseStringArray(value)
	}
}

func fetch() *[]Quote {
	quotes := new([]Quote)
	resp, err := http.Get("https://type.fit/api/quotes")
	if err != nil {
		log.Fatalf("Error fetching quotes")
	}
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, quotes); err != nil {
		log.Fatalf("Error processing JSON data")
	}
	return quotes
}

func main() {
	quotes := fetch()
	parsedQuotes := parseQuotesToMap(quotes)
	reverseQuotes(parsedQuotes)
	str, _ := json.Marshal(parsedQuotes)
	fmt.Println(string(str))
}
