package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

type quotes struct {
	Author string
	Quote  string
}

func main() {
	csvFile, _ := os.Open("quotesofgratitude.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	var listquotes []quotes
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		listquotes = append(listquotes, quotes{
			Author: line[1],
			Quote:  line[0],
		})
	}
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(listquotes)
	message := listquotes[n]
	fmt.Println(message.Quote, "-", message.Author)
}
