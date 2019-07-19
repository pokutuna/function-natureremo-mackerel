package main

import (
	"context"
	"log"

	function "github.com/pokutuna/function-natureremo-mackerel"
)

func main() {
	err := function.RemoToMackerel(context.Background(), struct{}{})
	if err != nil {
		log.Fatal(err)
	}
}
