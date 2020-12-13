package commands

import (
	"fmt"
	"log"
	"strings"
	azDevops "main/azDevops"
)

func getCommand(userInput string) string {

	trimmedInput := strings.TrimSpace(userInput)
	firstWord := strings.Split(trimmedInput, " ")[0]
	trimmedFirstWord := strings.ToLower(strings.TrimSpace(firstWord))
	log.Println("Parsed command:", trimmedFirstWord)

	return trimmedFirstWord
}

// HandleCommand handles user input string commands
func HandleCommand(userText string) string {
	log.Printf("Recieved: =>%s<=\n", userText)

	command := getCommand(userText)

	var response string
	switch command {
	case "listprs":
		response = azDevops.GetPRs()
	case "help":
		response = displayHelpMsg()
	case "hola":
		response = "Quiubo humano!"
	case "hello", "hi":
		response = "Hello human!"
	default:
		response = fmt.Sprintf("I did not understand the command: ```%s```. Type ```help``` for a list of available commands or ```help <command>``` for the command's documentation.", userText)
	}
	log.Println(response)
	return response
}
