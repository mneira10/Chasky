package commands

import (
	"fmt"
	"strings"
)

// HelpCommand struct
type helpCommand struct {
	availableCommands map[string]Command
	commInput         CommandInput
}

func (command helpCommand) getDescription() string {
	return "Listar los comandos disponibles."
}

func (command helpCommand) setCommandInput(commInput CommandInput) {
	command.commInput = commInput
}

func (command helpCommand) execute() {}

func (command helpCommand) getOutput() (string, bool) {

	var commandStrings []string

	commandContainer := `{
		"type": "Container",
		"items": [
			{
				"type": "ColumnSet",
				"columns": [
					{
						"type": "Column",
						"width": "stretch",
						"items": [
							{
								"type": "TextBlock",
								"text": "%s",
								"wrap": true,
								"horizontalAlignment": "Center",
								"separator": true,
								"spacing": "None",
								"height": "stretch"
							}
						]
					},
					{
						"type": "Column",
						"width": "stretch",
						"items": [
							{
								"type": "TextBlock",
								"text": "%s",
								"wrap": true
							}
						]
					}
				]
			}
		]
	}`

	availableCommands := GetAvailableCommands()

	availableCommands["help"] = helpCommand{}

	for commandName, externalCommand := range availableCommands {
		description := externalCommand.getDescription()

		commandContainerWithData := fmt.Sprintf(commandContainer, commandName, description)
		commandStrings = append(commandStrings, commandContainerWithData)
	}

	commandsText := strings.Join(commandStrings, ",")
	responseTxt := fmt.Sprintf(`{
		"type": "AdaptiveCard",
		"$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
		"version": "1.0",
		"body": [
			{
				"type": "TextBlock",
				"text": "Help",
				"wrap": true,
				"weight": "Bolder",
				"size": "Large"
			},
			{
				"type": "TextBlock",
				"text": "Estos son los comandos displonibles:",
				"wrap": true
			},
			%s
		]
	}`, commandsText)

	return responseTxt, true
}
