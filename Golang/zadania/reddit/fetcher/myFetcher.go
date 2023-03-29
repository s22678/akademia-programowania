package fetcher

import (
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

func (r *Response) Fetch() error {
	// url := []string {
	// 	"https://www.reddit.com/r/golang.json",
	// }
	url := "https://www.reddit.com/r/golang.json"
	client := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent", "myclient")

	res, getErr := client.Do(req)
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
	return nil
}
