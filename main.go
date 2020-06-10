package main

import (
	"fmt"
	"log"

	"citest/container"

	"github.com/kelseyhightower/envconfig"
)

type Environment struct {
	Port string `envconfig:"PORT" default:"5000"`
}

func main() {
	var vars Environment
	if err := envconfig.Process("", &vars); err != nil {
		log.Fatalf("[Error] Failed to process env var: %s", err)
	}

	c := container.New()

	c.Append(container.Baggage{Name: "test1", Content: "test-content1"})
	c.Append(container.Baggage{Name: "test2", Content: "test-content2"})
	c.Append(container.Baggage{Name: "test3", Content: "test-content3"})
	fmt.Println(c)

	b := c.Find("test2")
	fmt.Println(b)
}
