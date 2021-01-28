package commands

import (
	"errors"
	"log"
	"strings"
)

// CommandInput struct
type CommandInput struct {
	commandName string
	args        []string
}

func (commandInput CommandInput) String() string {
	completeCommand := append([]string{commandInput.commandName}, commandInput.args...)
	return strings.Join(completeCommand, " ")
}

func (commandInput *CommandInput) parseInput(userInput string) error {

	trimmedInput := strings.TrimSpace(userInput)
	splitCommand := strings.Split(trimmedInput, " ")

	if len(splitCommand) == 0 {
		return errors.New("Empty input, received: " + userInput)
	}

	commandInput.commandName = strings.ToLower(splitCommand[0])

	if len(splitCommand) > 1 {
		commandInput.args = splitCommand[1:]
	}

	log.Println("Parsed command:", commandInput)

	return nil
}

/*
	HandleCommand recieves the user input
	and returns the string to be
	returned to the user.
*/
func HandleCommand(userText string) (string, bool) {
	log.Printf("Recieved: =>%s<=\n", userText)

	var commandInput CommandInput
	error := commandInput.parseInput(userText)

	if error != nil {
		log.Fatal(error)
	}

	command := commandFactory(commandInput)

	command.Execute()

	log.Println("Getting command output")
	response, isJSONText := command.GetOutput()
	log.Println("Retrieved command output")
	return response, isJSONText
}
