package main

import (
	"fmt"
	"log"
	"weathercli/api"
)

func main() {
	dataset := api.ApiCall()

	res, err := api.PrettyStruct(dataset)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}
