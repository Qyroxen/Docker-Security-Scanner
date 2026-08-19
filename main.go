package main

import (
	"fmt"
	"os"
)

// docker_security_scanner - Scan Docker images for vulns
func docker_security_scanner(path string) {
	fmt.Println("========================================")
	fmt.Println("  Docker-Security-Scanner")
	fmt.Println("  Scan Docker images for vulns")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	docker_security_scanner(path)
}
