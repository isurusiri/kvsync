package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"time"

	consul_api "github.com/hashicorp/consul/api"
	"github.com/hashicorp/nomad/api"
)

func main() {

	nomadHostPtr  := flag.String("n", "http://localhost:4646", "Nomad host url")
	consulHostPtr := flag.String("c", "http://localhost:8500", "Consul host url")

	flag.Parse()


	fmt.Println("Starting kvsync...")
	fmt.Println("Connecting to Nomad host...")
	client := createNomadClient(nomadHostPtr)

	fmt.Println("Connecting to Consul host...")
	consulClient := createConsulClient(consulHostPtr)

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
	if err != nil {
		fmt.Printf("Error while listening to Nomad event stream")
	}

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
	
					fmt.Printf(*job.ID + " " + e.Type + "\n")
					// fmt.Printf(createKeyValuePairs(job.Meta))
					// fmt.Println()

					if e.Type == "JobRegistered" {
						for key, value := range job.Meta {
							writeToKV(kv, createKeyWithJobID(*job.ID, key), value)
						}
					} else if e.Type == "JobDeregistered" {
						for key := range job.Meta {
							removeFromKV(kv, createKeyWithJobID(*job.ID, key))	
						}
					}
				}
			}
		case <-time.After(120 * time.Second):
			fmt.Printf("... ")
		}
	}

}
 // Creates a client to connect with Nomad api identified by nomadHost
func createNomadClient(nomadHost *string) *api.Client {
	nomadCfg         := api.DefaultConfig()
	nomadCfg.Address = *nomadHost
	client, err := api.NewClient(nomadCfg)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Connected to Nomad host via %s ...\n", *nomadHost)
	}

	return client
}

// Creates a client to connect with Consul api identified by consulHost
func createConsulClient(consulHost *string) *consul_api.Client {
	consulCfg         := consul_api.DefaultNonPooledConfig()
	consulCfg.Address = *consulHost
	consulClient, err := consul_api.NewClient(consulCfg)
	if err != nil {
		fmt.Printf(err.Error())
	} else {
		fmt.Printf("Connected to Consul host via %s ...\n", *consulHost)
	}

	return consulClient
}

// Creates a string containing key and value in the key=value format
func createKeyValuePairs(m map[string]string) string {
    b := new(bytes.Buffer)
    for key, value := range m {
        fmt.Fprintf(b, "%s=\"%s\"\n", key, value)
    }
    return b.String()
}

// Writes a key value pair to Consul KV store
func writeToKV(kv *consul_api.KV, key, value string) {
	p := &consul_api.KVPair{Key: key, Value: []byte(value)}
	_, err := kv.Put(p, nil)
	if err != nil {
		panic(err)
	}
}

// Removes a key value pair identified by the key from Consul KV store
func removeFromKV(kv *consul_api.KV, key string) {
	_, err := kv.Delete(key, nil)
	if err != nil {
		panic(err)
	}
}

// Prepend the nomad job id to the key and return it as a single string
// in the job-id-key format
func createKeyWithJobID(jobID string, key string) string {
	return fmt.Sprintf("%s-%s", jobID, key)
}