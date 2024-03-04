package main

import (
	"log"
	"testing"
	"time"

	"github.com/hashicorp/mdns"
)

func TestMDNSDiscovery(t *testing.T) {
	serviceName := "_mqtt._tcp"

	entriesCh := make(chan *mdns.ServiceEntry, 10)

	if err := mdns.Lookup(serviceName, entriesCh); err != nil {
		log.Fatal(err)
	}

	time.Sleep(5 * time.Second)

	for entry := range entriesCh {
		log.Println("discovered service:", entry.Name)
		return
	}

	if len(entriesCh) == 0 {
		log.Println("no services discovered")
        t.FailNow()
    }
}