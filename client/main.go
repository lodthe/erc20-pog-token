package main

import (
	"fmt"
	"log"
	"math/big"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("not enough arguments ('go run *)")
	}

	infuraURL := os.Getenv("INFURA_URL")
	privateKey := os.Getenv("PRIVATE_KEY")
	contractAddress := os.Getenv("CONTRACT_ADDRESS")
	if contractAddress == "" {
		contractAddress = "0xD727259569e631AD36d12b9C18cd754b69A8C4de"
	}

	client, err := NewClient(infuraURL, privateKey, contractAddress)
	if err != nil {
		log.Fatal("failed to initialize client:", err)
	}

	switch os.Args[2] {
	case "deploy":
		addr, tx, err := client.Deploy(big.NewInt(1_000_000_000))
		if err != nil {
			log.Fatal("failed to deploy:", err)
		}

		fmt.Printf("deployed successfully\ntx: \t%s\naddress: \t%s\n", tx.Hash(), addr)

	case "add":
		if len(os.Args) < 5 {
			log.Fatal("not enough arguments (provide address and username)")
		}

		tx, err := client.AddViewer(os.Args[3], os.Args[4])
		if err != nil {
			log.Fatal("failed to add viewer:", err)
		}

		fmt.Printf("viewer has been added: %s\n", tx.Hash())

	case "remove":
		if len(os.Args) < 4 {
			log.Fatal("not enough arguments (provide address)")
		}

		tx, err := client.RemoveViewer(os.Args[3])
		if err != nil {
			log.Fatal("failed to add viewer:", err)
		}

		fmt.Printf("viewer has been removed: %s\n", tx.Hash())

	case "logs":
		if len(os.Args) < 4 {
			log.Fatal("not enough arguments (provide 'add' or 'remove')")
		}

		if os.Args[3] == "add" {
			err = client.PrintLogsAdd()
		} else {
			err = client.PrintLogsRemove()
		}

		if err != nil {
			log.Fatal("failed to print logs:", err)
		}

	default:
		log.Fatal("unknown command", os.Args[2])
	}
}
