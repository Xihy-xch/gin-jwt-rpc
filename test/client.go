package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	httpDo()
}

func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/login", strings.NewReader("username=admin&password=admin"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "jwtCookie=Snd0VHlwZTpKd3QsQWxnb3JpdGhtOkhTMjU2LA==.VXNlcm5hbWU6YWRtaW4s.f8e1c8b25bcf72aad7d7dc8d4011cebfd84f2b764501eaea10285368c38c033b")

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}
