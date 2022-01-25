package main

import (
	"bytes"
	"context"
	"fmt"
	"time"

	consule_api "github.com/hashicorp/consul/api"
	"github.com/hashicorp/nomad/api"
)

func main() {
	cfg         := api.DefaultConfig()
	cfg.Address = "http://localhost:4646"
	client, err := api.NewClient(cfg)
	if err != nil {
		fmt.Printf(err.Error())
	}

	consulCfg         := consule_api.DefaultNonPooledConfig()
	consulCfg.Address = "http://localhost:8500"
	consulClient, err := consule_api.NewClient(consulCfg)
	if err != nil {
		fmt.Printf(err.Error())
	}

	// Get a handle to the KV
	kv := consulClient.KV()

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

					if e.Type == "JobRegistered" {
						for key, value := range job.Meta {
							writeToKV(kv, key, value)
						}
					} else if e.Type == "JobDeregistered" {
						for key, _ := range job.Meta {
							removeFromKV(kv, key)	
						}
					}
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

func writeToKV(kv *consule_api.KV, key, value string) {
	p := &consule_api.KVPair{Key: key, Value: []byte(value)}
	_, err := kv.Put(p, nil)
	if err != nil {
		panic(err)
	}
}

func removeFromKV(kv *consule_api.KV, key string) {
	_, err := kv.Delete(key, nil)
	if err != nil {
		panic(err)
	}
}