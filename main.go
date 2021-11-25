package main

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/nomad/api"
)

func main() {
	cfg := api.DefaultConfig()
	
	cfg.Address = "http://localhost:4646"

	// handdle "https://" in Address if HTTPS is used

	// timeout, err := strconv.Atoi("30")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// cfg.HttpClient.Timeout = time.Duration(timeout) * time.Second

	client, err := api.NewClient(cfg)
	if err != nil {
		fmt.Printf(err.Error())
	}

	// build event stream request
	events := client.EventStream()
	q := &api.QueryOptions{}
	topics := map[api.Topic][]string{
		api.TopicNode: {"*"},
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	streamCh, err := events.Stream(ctx, topics, 0, q)

	select {
	case event := <-streamCh:
		if event.Err != nil {
			fmt.Printf(event.Err.Error())
		}
		for _, e := range event.Events {
			// verify that we get a node
			n, err := e.Node()
			if err != nil {
				fmt.Printf(err.Error())
			}
			
			// eval, err := e.Evaluation()
			// if err != nil {
			// 	fmt.Printf(err.Error())
			// }

			fmt.Printf(n.Name)
			// fmt.Printf(eval.ID)
		}
	case <-time.After(5 * time.Second):
		fmt.Printf("failed waiting for event stream event")
	}

}