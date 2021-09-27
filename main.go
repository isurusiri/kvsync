package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func createConsulKVRecord() {
	putBody, _ := json.Marshal("isuru")
	responseBody := bytes.NewBuffer(putBody)

	req, err := http.NewRequest(http.MethodPut, "http://consul.service.galactica.consul:8500/v1/kv/meta-sync", responseBody)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	client := &http.Client{}

	resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }

    fmt.Println(resp.StatusCode)
	fmt.Println(resp.Body)
	fmt.Println("")

	resp.Body.Close()
}

func listNomadServers() {
	resp, err := http.Get("http://localhost:4646/v1/nodes")
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(body))

}

func listNomadEvaluations() {
	resp, err := http.Get("http://localhost:4646/v1/evaluations")
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(body))
}

func receiveEvents() {
	resp, err := http.Get("http://localhost:4646/v1/event/stream?topic=Evaluation")
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}


	fmt.Printf(string(body))
}

func main() {
	fmt.Println("Meta sync")
	// createConsulKVRecord()

	listNomadServers()
	fmt.Println("")
	// listNomadEvaluations()
	receiveEvents()
}