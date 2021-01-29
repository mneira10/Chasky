package commands

import "log"

// Command behavior
type Command interface {
	SetCommandInput(CommandInput CommandInput)
	GetDescription() string
	Execute()
	GetOutput() (string, bool)
}

// GetAvailableCommands returns the available commands
// mapped to an instance of a Command
func GetAvailableCommands() map[string]Command {

	return map[string]Command{
		"hello": helloCommand{},
		"hola":  helloCommand{},
	}
}

func commandFactory(commandInput CommandInput) Command {

	availableCommands := GetAvailableCommands()
	lowerCaseCommandName := commandInput.commandName
	log.Printf("lowercaseCommandName: =>%s<=", lowerCaseCommandName)

	if command, ok := availableCommands[lowerCaseCommandName]; ok {
		log.Println("Setting command input")
		command.SetCommandInput(commandInput)
		return command
	}

	// return initialized helpCommand struct
	var helpComm helpCommand
	helpComm.availableCommands = availableCommands
	helpComm.SetCommandInput(commandInput)
	return helpComm
}
