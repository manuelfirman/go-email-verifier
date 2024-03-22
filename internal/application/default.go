package application

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/manuelfirman/go-email-verifier/internal/verifier"
)

// Run starts the application
func Run() {
	// Create a new scanner to read from the standard input
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter the domain names to check, one per line. Press Ctrl+C to finish.")
	fmt.Print("> ")

	// Read the input from the user
	for scanner.Scan() {
		// Call the CheckDomain function from the verifier package to check the domain
		verifier.CheckDomain(scanner.Text())
		fmt.Println("--------------------------------------------------")
		fmt.Print("> ")
	}

	// Check for errors during the scan
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
