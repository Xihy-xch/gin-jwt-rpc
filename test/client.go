package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

var cookie string

func main() {
	for i := 0; i < 2; i++ {
		fmt.Println(cookie)
		httpDo()
		time.Sleep(time.Duration(2) * time.Second)
	}

}

func httpDo() {
	client := &http.Client{}

	req, err := http.NewRequest("POST", "http://127.0.0.1:8080/login", strings.NewReader("username=admin&password=admin"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	cookie1 := &http.Cookie{Name: "jwtCookie", Value: cookie, HttpOnly: true}
	req.AddCookie(cookie1)

	resp, err := client.Do(req)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	var tmp map[string]interface{}
	json.Unmarshal(body, &tmp)
	cookie = tmp["token"].(string)
}
