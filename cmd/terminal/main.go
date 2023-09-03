package main

import (
	"bufio"
	"flag"
	"fmt"
	"ggpt/pkg/ggpt"
	"os"
	"strings"

	"github.com/eiannone/keyboard"
)

func main() {
	role := flag.String("role", "", "The name of the role to use from your roles in the config.")
	flag.Parse()

	// create sessionHistory in memory
	sessionHistory := []string{}
	currentIdx := -1

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

		if input != "" {
			sessionHistory = append(sessionHistory, input)
			currentIdx = len(sessionHistory) - 1

			switch input {
			case "exit":
				break

			case "history", "hh":
				command, err := navigateHistory(currentIdx, sessionHistory)
				if err != nil {
					fmt.Println("Error navigating commands:", err)
					break
				}

				response, err := ggptAgent.SendGPTRequest(command)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}

				fmt.Println("Response:", response)

			default:

				response, err := ggptAgent.SendGPTRequest(input)
				if err != nil {
					fmt.Println("Error:", err)
					return
				}

				fmt.Println("Response:", response)
			}
		}
	}
}

func navigateHistory(currentIdx int, sessionHistory []string) (string, error) {
	if err := keyboard.Open(); err != nil {
		return "", err
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {
		_, key, err := keyboard.GetKey()
		if err != nil {
			return "", err
		}

		switch key {
		case keyboard.KeyArrowUp:
			if currentIdx > 0 {
				currentIdx--
				fmt.Printf("\r%s\033[K", sessionHistory[currentIdx])
			}
		case keyboard.KeyArrowDown:
			if currentIdx < len(sessionHistory)-1 {
				currentIdx++
				fmt.Printf("\r%s\033[K", sessionHistory[currentIdx])
			}
		case keyboard.KeyEsc, keyboard.KeyCtrlC, keyboard.KeyCtrlD:
			return "", nil
		case keyboard.KeyEnter:
			return sessionHistory[currentIdx], nil
		}
	}
}
