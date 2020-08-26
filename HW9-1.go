package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)
func main() {
	resp, err := http.Get("http://localhost:8081/log")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
	requestBody, err := json.Marshal(map[string]string{
		"login": "Marry",
		"password": "18",
	})
	resp2, err := http.Post("http://localhost:8080/log","application/json",bytes.NewBuffer(requestBody) )
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp2.Body.Close()
	io.Copy(os.Stdout, resp2.Body)
	/*var result map[string]interface{}
	json.NewDecoder(resp2.Body).Decode(&result)
	log.Print(result["form"])*/
}

