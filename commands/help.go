package commands

import (
	"log"
)

// HelpCommand struct
type helpCommand struct {
	availableCommands map[string]Command
	commInput         CommandInput
}

func (command helpCommand) GetDescription() string {
	return "List the available commands and their descriptions"
}

func (command helpCommand) SetCommandInput(commInput CommandInput) {
	command.commInput = commInput
	log.Println("Set command input ")
}

func (command helpCommand) Execute() {
	log.Println("executing help command")
	command.availableCommands = GetAvailableCommands()
}

func (command helpCommand) GetOutput() (string, bool) {
	// for commandName, command := range command.availableCommands {
	// 	description := command.GetDescription()
	// }
	responseTxt := `{
		"$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
		"type": "AdaptiveCard",
		"version": "1.0",
		"body": [
			{
				"type": "Container",
				"items": [
					{
						"type": "TextBlock",
						"text": "Help",
						"weight": "Bolder",
						"size": "Medium"
					}
				]
			},
			{
				"type": "Container",
				"items": [
					{
						"type": "TextBlock",
						"text": "List of available commands",
						"weight": "Bolder",
						"wrap": true
					}
				]
			}
		]
	}`

	return responseTxt, true
}

func (command helpCommand) Help() (string, bool) {
	return "some help", false
}
