package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"os/exec"

	"gopkg.in/yaml.v3"
)

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {

	if len(os.Args) != 2 {
		log.Fatal("incorrect arguments")
	}

	originalCommand := os.Getenv("SSH_ORIGINAL_COMMAND")

	if originalCommand == "" {
		log.Fatal("SSH_ORIGINAL_COMMAND environment variable missing")
	}

	ipB64 := os.Args[1]

	ip, err := base64.StdEncoding.DecodeString(ipB64)
	if err != nil {
		log.Fatal(err)
	}

	type ipS struct {
		AllowedCommands []string `yaml:"allowedCommands"`
	}

	var ac ipS

	err = yaml.Unmarshal(ip, &ac)

	if err != nil {
		log.Fatal(err)
	}

	if contains(ac.AllowedCommands, originalCommand) {

		cmd := exec.Command(originalCommand)
		stdout, err := cmd.Output()

		if err != nil {
			log.Fatal(err.Error())
		}
		// Print the output
		fmt.Println(string(stdout))
	} else {
		fmt.Println("unauthorized")
	}

}
