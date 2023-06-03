package client

import (
	"github.com/pojntfx/go-nbd/pkg/client"
	"net"
	"os"
	"os/signal"
)

func NewClient(blockSize int64) *Client {
	return &Client{BlockSize: blockSize}
}

func (c *Client) List(address string) ([]string, error) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return nil, err
	}
	exports, err := client.List(conn)
	if err != nil {
		return nil, err
	}
	return exports, nil
}

func (c *Client) Get(address, name, pathTo string) error {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	f, err := os.Open(pathTo)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	go func() {
		for range sigCh {
			if err = client.Disconnect(f); err != nil {
				panic(err)
			}
			os.Exit(0)
		}
	}()

	if err = client.Connect(conn, f, &client.Options{
		ExportName: name,
		BlockSize:  uint32(c.BlockSize),
	}); err != nil {
		return err
	}
	return nil
}
