package server

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/miguel250/kuma/http/client"
)

func TestServerStart(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(rw, "Hello World!")
	})

	server := New(nil, mux)
	server.Start()
	<-server.start

	client := client.New(nil)
	ctx := context.Background()

	resp, err := client.Get(ctx, server.Addr)

	if err != nil {
		t.Fatalf("failed to create http client with %s", err)
	}

	defer func() {
		err := resp.Body.Close()

		if err != nil {
			t.Errorf("failed to close response body with %s", err)
		}
	}()

	b, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fatalf("failed to read content of response with %s", err)
	}

	want := "Hello World!"
	got := string(b)
	if want != got {
		t.Errorf("Response body didn't match got: (%s), want: (%s)", got, want)
	}
}
