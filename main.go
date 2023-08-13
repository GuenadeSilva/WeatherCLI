package main

import (
	"fmt"
	"weathercli/api"
)

func main() {
	fmt.Println(string(api.ApiCall()))
}
