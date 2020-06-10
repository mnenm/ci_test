package main

import (
	"fmt"

	"citest/container"
)

func main() {
	c := container.New()

	c.Append(container.Baggage{Name: "test1", Content: "test-content1"})
	c.Append(container.Baggage{Name: "test2", Content: "test-content2"})
	c.Append(container.Baggage{Name: "test3", Content: "test-content3"})
	fmt.Println(c)

	b := c.Find("test2")
	fmt.Println(b)
}
