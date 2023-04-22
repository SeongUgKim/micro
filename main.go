package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/SeongUgKim/micro/client"
	"log"
)

func runClient() {
	client := client.New("http://localhost:3000")
	price, err := client.FetchPrice(context.Background(), "ETH")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", price)
}

func runServer(listenAddr string) {
	svc := NewLoggingService(NewMetricService(&priceFetcher{}))
	server := NewJSONAPIServer(listenAddr, svc)
	server.Run()
}

func main() {
	listenAddr := flag.String("listenaddr", ":3000", "listen address the service is running")
	mode := flag.String("mode", "server", "")
	flag.Parse()
	if *mode == "client" {
		runClient()
	} else {
		runServer(*listenAddr)
	}
}
