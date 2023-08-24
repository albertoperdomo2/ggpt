package main

import (
	"bufio"
	"flag"
	"fmt"
	"ggpt/pkg/ggpt"
	"os"
	"strings"
)

func main() {
	role := flag.String("role", "", "The name of the role to use from your roles in the config.")
	flag.Parse()

	// bufio reader to read user input
	reader := bufio.NewReader(os.Stdin)

	// create GGPT agent
	ggptAgent, err := ggpt.NewAgent(*role)

	if err != nil {
		fmt.Println("Error creating GGPT agent: %s", err)
		return
	}

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		// trim whitespace and newline characters
		input = strings.TrimSpace(input)

		response, err := ggptAgent.SendGPTRequest(input)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		fmt.Println("Response:", response)

		if input == "exit" {
			break
		}
	}
}
