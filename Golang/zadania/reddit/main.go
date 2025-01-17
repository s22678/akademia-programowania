package main

import (
	"context"
	"io"
	"log"
	"os"
	"reddit/fetcher"
)

var (
	file *os.File
	err  error
)

func init() {
	LOG_FILE := "./app.log"
	file, err = os.OpenFile(LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0660)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("Initialzing reddit fetcher app")
}

func main() {
	var f fetcher.RedditFetcher // do not change
	var w io.Writer             // do not change

	f = &fetcher.Response{}
	w, err = os.Create("output/my_output")
	if err != nil {
		log.Fatalln("Failed to create my_output file")
	}
	defer file.Close()

	ctx := context.Background()

	err := f.Fetch(ctx)
	if err != nil {
		log.Fatal("Fetching data unsuccessful")
	}

	err = f.Save(w)
	if err != nil {
		log.Fatal("Saving data to file unsuccessful")
	}

	w = os.Stdout
	err = f.Save(w)
	if err != nil {
		log.Fatal("Saving data to stdout unsuccessful")
	}
}
