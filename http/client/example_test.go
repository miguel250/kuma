package client_test

import (
	"context"
	"fmt"
	"io"
	"log"

	"github.com/miguel250/kuma/http/client"
)

func ExampleClient_Get() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := client.New()
	resp, err := c.Get(ctx, "http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		log.Fatalf("request failed with status code %d and\nbody: %s\n", resp.StatusCode, body)
	}

	fmt.Printf("%s", body)
}
