package main

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/nomad/api"
)

func main() {
	cfg := api.DefaultConfig()
	
	cfg.Address = "http://localhost:4646"

	client, err := api.NewClient(cfg)
	if err != nil {
		fmt.Printf(err.Error())
	}

	// build event stream request
	events := client.EventStream()
	q := &api.QueryOptions{}
	topics := map[api.Topic][]string{
		api.TopicJob: {"*"},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	streamCh, err := events.Stream(ctx, topics, 0, q)

	for {
		select {
		case event := <-streamCh:
			if event.Err != nil {
				fmt.Printf(event.Err.Error())
			}
			for _, e := range event.Events {

				if e.Type == "JobRegistered" || e.Type == "JobDeregistered" {
					job, err := e.Job()
					if err != nil {
						fmt.Printf(err.Error())
					}
	
					fmt.Printf(*job.ID + "\n")
					fmt.Printf(createKeyValuePairs(job.Meta))
					fmt.Println()
				}
	
				fmt.Printf(e.Type + "\n")
			}
		case <-time.After(120 * time.Second):
			fmt.Printf("... ")
		}
	}

}

func createKeyValuePairs(m map[string]string) string {
    b := new(bytes.Buffer)
    for key, value := range m {
        fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
    }
    return b.String()
}