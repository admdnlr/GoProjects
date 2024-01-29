package main

import (
	"fmt"
	"os"
)

func main() {
	// Getenv returns the value of the environment variable named by the key.

	uName := os.Getenv("USERNAME")
	uDomain := os.Getenv("USERDOMAIN")
	processorArchitecture := os.Getenv("PROCESSOR_ARCHITECTURE")
	processorIdentifier := os.Getenv("PROCESSOR_IDENTIFIER")
	processorLevel := os.Getenv("PROCESSOR_LEVEL")
	goRoot := os.Getenv("GOROOT")
	goPath := os.Getenv("GOPATH")
	homePath := os.Getenv("HOMEPATH")

	fmt.Println("USERNAME:", uName)
	fmt.Println("USERDOMAIN:", uDomain)
	fmt.Println("PROCESSOR_ARCHITECTURE:", processorArchitecture)
	fmt.Println("PROCESSOR_IDENTIFIER:", processorIdentifier)
	fmt.Println("PROCESSOR_LEVEL:", processorLevel)
	fmt.Println("GOROOT:", goRoot)
	fmt.Println("GOPATH:", goPath)
	fmt.Println("HOMEPATH:", homePath)

}
