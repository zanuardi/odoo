package main

import (
	"fmt"
	"log"

	"github.com/kolo/xmlrpc"
)

func main() {
	url := "https://demo6.odoo.com"
	// db := "demo6"
	// username := "admin"
	// password := "admin"

	client, err := xmlrpc.NewClient(fmt.Sprintf("%s/xmlrpc/2/common", url), nil)
	if err != nil {
		log.Fatal(err)
	}
	common := map[string]any{}

	if err := client.Call("version", nil, &common); err != nil {
		fmt.Println("TES version ", &common, err)
		log.Fatal(err)
	}

	// client.Call("Bugzilla.version", nil, &result)
	// fmt.Printf("Version: %s\n", &result, result.Version) // Version: 4.2.7+
}
