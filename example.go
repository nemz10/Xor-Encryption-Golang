package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func xor(input string) string {
	var result string
	for _, char := range input {
		result += fmt.Sprintf("%x", char^5)
	}
	return result
}

func main() {
	url := "https://api.tiktok.com/login" // note: placeholder api 

	body, err := json.Marshal(map[string]string{
		"username": xor("username"),
		"password": xor("password"),
	})
	if err != nil {
		panic(err) 
	}

	payload := bytes.NewBuffer(body)

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		panic(err)
	}

	req.Header.Set("header1", "")
	req.Header.Set("header2", "")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// convert response body to string and print it
	sb := string(response)
	fmt.Println(sb) // output the response to the cmd
}
