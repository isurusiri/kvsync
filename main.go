package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/isurusiri/funnel/models"
)

func createConsulKVRecord() {
	putBody, _ := json.Marshal("isuru")
	responseBody := bytes.NewBuffer(putBody)

	req, err := http.NewRequest(http.MethodPut, "http://localhost:8500/v1/kv/meta-sync", responseBody)
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

func receiveStreamedEvents() {
	resp, err := http.Get("http://localhost:4646/v1/event/stream?topic=Evaluation")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(resp.Body)

	var events models.EventStream

	for {
		events = printEvents(reader)
		if len(events.Events) > 0 {
			queryJobByJobID(events.Events[0].Payload["Evaluation"].JobID)
			fmt.Println("")
		}
		// line, err := reader.ReadString('\n')
		// if err != nil {
		// 	panic(err)
		// }
		
		// fmt.Println(line)
		// fmt.Println("---")
	}
}

func queryJobByJobID(jobID string) {
	resp, err := http.Get(fmt.Sprintf("http://localhost:4646/v1/job/%s", jobID))
	if err != nil {
		panic(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf(string(body))
}

func printEvents(body io.Reader) models.EventStream{
	var events models.EventStream

	err := json.NewDecoder(body).Decode(&events)
	if err != nil {
        log.Fatal(err)
    }

	return events
}

func main() {
	fmt.Println("Meta sync")
	// createConsulKVRecord()

	listNomadServers()
	fmt.Println("")
	// listNomadEvaluations()
	// receiveEvents()

	fmt.Println("")
	fmt.Println("**** Stream ****")
	fmt.Println("")
	receiveStreamedEvents()
}