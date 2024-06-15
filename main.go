package main

import (
	"fmt"
	"github.com/31-rat4/gmail-cli/cmd/ascess"
	"log"
	"os"
)

func main() {

	Ascess := ascess.NewAscessHandler("test2")
	fmt.Println(Ascess.GetConfigFromFile())
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	fmt.Println(b)
}
