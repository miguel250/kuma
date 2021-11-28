package client

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestMethods(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		bodyb, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("ReadAll on req.Body: %v", err)
		}
		fmt.Fprintf(w, "CONTENT=%s METHOD=%s CONTENT=%s", r.Header.Get("content-type"), r.Method, string(bodyb))
	}))

	defer ts.Close()

	for _, test := range []struct {
		method, payload, expected string
	}{
		{"GET", "", "CONTENT= METHOD=GET CONTENT="},
		{"POST", "content", "CONTENT=text/plain METHOD=POST CONTENT=content"},
	} {
		t.Run(fmt.Sprintf("METHOD: %s", test.method), func(t *testing.T) {
			client := New()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			switch test.method {
			case "GET":
				resp, err := client.Get(ctx, ts.URL)

				if err != nil {
					t.Fatalf("Failed to make request with (%v)", err)
				}

				defer resp.Body.Close()

				bodyb, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Fatalf("ReadAll on req.Body: %v", err)
				}

				if g := string(bodyb); g != test.expected {
					t.Errorf("got body %q, want %q", g, test.expected)
				}
			case "POST":
				b := strings.NewReader(test.payload)
				resp, err := client.Post(ctx, ts.URL, "text/plain", b)

				if err != nil {
					t.Fatalf("Failed to make request with (%v)", err)
				}

				defer resp.Body.Close()
				bodyb, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					t.Fatalf("ReadAll on req.Body: %v", err)
				}

				if g := string(bodyb); g != test.expected {
					t.Errorf("got body %q, want %q", g, test.expected)
				}
			default:
				t.Errorf("Method not supported %s", test.method)
			}
		})
	}
}

func TestDeadlineContext(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		bodyb, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatalf("ReadAll on req.Body: %v", err)
		}
		time.Sleep(50 * time.Millisecond)
		fmt.Fprintf(w, "CONTENT=%s METHOD=%s CONTENT=%s", r.Header.Get("content-type"), r.Method, string(bodyb))
	}))

	defer ts.Close()

	client := New()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Millisecond)

	defer cancel()
	errChan := make(chan error, 1)
	go func() {
		select {
		case <-time.After(50 * time.Millisecond):
			errChan <- errors.New("context should have timed out")
		case <-ctx.Done():
			errChan <- ctx.Err()
		}
	}()

	resp, err := client.Get(ctx, ts.URL)

	if err == nil {
		defer resp.Body.Close()
		t.Fatal("Expecting for this to return an error")
	}

	err = <-errChan
	if err != context.DeadlineExceeded {
		t.Errorf("Context failed with an error %v", err)
	}

}

func TestTLSClient(t *testing.T) {
	ts := httptest.NewTLSServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Success")
	}))
	defer ts.Close()

	cert, err := x509.ParseCertificate(ts.TLS.Certificates[0].Certificate[0])
	if err != nil {
		t.Fatalf("Failed to parse Certificates %v", err)
	}

	certpool := x509.NewCertPool()
	certpool.AddCert(cert)

	tlsOption := WithTLSConfig(
		&tls.Config{
			RootCAs: certpool,
		},
	)

	client := New(tlsOption)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	resp, err := client.Get(ctx, ts.URL)

	if err != nil {
		t.Fatalf("Failed to make request with (%v)", err)
	}

	defer resp.Body.Close()

	bodyb, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ReadAll on req.Body: %v", err)
	}

	expected := "Success"
	if g := string(bodyb); g != expected {
		t.Errorf("got body %q, want %q", g, expected)
	}
}
