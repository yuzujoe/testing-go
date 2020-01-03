package main

import (
	"testing-go/api/providers"

	"fmt"
)

func main() {
	country, err := providers.GetCoutry("AR")
	fmt.Println(err)
	fmt.Println(country)
}
