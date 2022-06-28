package main

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	s := gocron.NewScheduler(time.UTC)

	s.Every(10).Seconds().Do(func() {
		response, error := http.PostForm("https://eofze3br6mxmqb5.m.pipedream.net", url.Values{"key": {"Value"}, "id": {"123"}})
		if error != nil {
			// Error calling from this server code
			fmt.Println("HTTP call failed:", error)
		} else {
			defer response.Body.Close()
			if response.StatusCode != http.StatusOK {
				// Error from api
				fmt.Println("Non-OK HTTP status:", response.StatusCode)
			} else {
				// OK
				fmt.Println(response.Body)
			}
		}
	})
	s.StartBlocking()
}
