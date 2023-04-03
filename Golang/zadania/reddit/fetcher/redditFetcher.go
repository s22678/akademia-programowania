package fetcher

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type Response struct {
	Data struct {
		Children []struct {
			Data struct {
				Title string `json:"title"`
				URL   string `json:"url"`
			} `json:"data"`
		} `json:"children"`
	} `json:"data"`
}

type RedditFetcher interface {
	Fetch(context.Context) error
	Save(io.Writer) error
}

func (r *Response) Fetch(ctx context.Context) error {

	url := "https://www.reddit.com/r/golang.json"
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "myclient")
	req.Header.Add("Content-Type", "application/json")

	res, getErr := http.DefaultClient.Do(req)
	if getErr != nil {
		log.Fatalln(getErr)
	}
	log.Println("Connected to:", url, "status:", res.Status)
	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	jsonErr := json.Unmarshal(body, r)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	log.Println("Fetched", url, "data decoded")

	return nil
}

func (r *Response) Save(w io.Writer) error {
	var b []byte
	for _, val := range r.Data.Children {
		b = append(b, []byte(val.Data.Title+"\n")...)
		b = append(b, []byte(val.Data.URL+"\n")...)
	}

	_, err := w.Write(b)
	if err != nil {
		log.Fatalln("Error writing data")
		return err
	}
	log.Println("Fetched data saved")
	return nil
}
