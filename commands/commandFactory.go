package commands

import "log"

// Command behavior
type Command interface {
	setCommandInput(commInput CommandInput)
	getDescription() string
	execute()
	getOutput() (string, bool)
}

// GetAvailableCommands returns the available commands
// mapped to an instance of a Command
func GetAvailableCommands() map[string]Command {

	return map[string]Command{
		"hello":   helloCommand{},
		"hola":    helloCommand{},
		"listprs": listPRsCommand{},
	}
}

func commandFactory(commandInput CommandInput) Command {

	availableCommands := GetAvailableCommands()
	lowerCaseCommandName := commandInput.commandName
	log.Printf("lowercaseCommandName: =>%s<=", lowerCaseCommandName)

	if command, ok := availableCommands[lowerCaseCommandName]; ok {
		log.Println("Setting command input")
		command.setCommandInput(commandInput)
		return command
	}

	// return initialized helpCommand struct
	var helpComm helpCommand
	helpComm.availableCommands = availableCommands
	helpComm.setCommandInput(commandInput)
	return helpComm
}
