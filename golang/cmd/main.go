package main

import (
	"fmt"
	"os"

	"example.com/algo"
)

func main() {
	if err := algo.Run(); err != nil {
		fmt.Printf("Run algorithm failure, err:%v\n", err)
		os.Exit(1)
	}
}
