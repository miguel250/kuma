/*
Package http provides a wrapper of Go standard http module with recommended defaults values and context.

Get and Post requests

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := client.New(nil)
	resp, err := c.Get(ctx, 'http://example.com')
	...

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	c := client.Client(nil)
	resp, err := c.Post(ctx, 'http://example.com', "text/plain", &buf)
*/
package client
