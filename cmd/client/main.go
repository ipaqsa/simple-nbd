package main

import (
	"github.com/ipaqsa/simple-nbd/pkg/client"
	"log"
)

func main() {
	c := client.NewClient(0)

	list, err := c.List(":7002")
	if err != nil {
		log.Fatal(err)
	}
	for _, e := range list {
		println(e)
	}

	//Use path as /dev/nbd*
	//c.Get(":7002", "disk", "/dev/nbd0")
}
