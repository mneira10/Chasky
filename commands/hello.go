package commands

import (
	"fmt"
	nasaAPI "main/nasaAPI"
)

type helloCommand struct {
}

func (c helloCommand) Execute() {}

func (c helloCommand) GetDescription() string {
	return "Saludar!"
}

func (c helloCommand) SetCommandInput(command CommandInput) {}

func (c helloCommand) GetOutput() (string, bool) {

	apodImageURL := nasaAPI.GetAPODURL()

	response := `
	{
		"$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
		"type": "AdaptiveCard",
		"version": "1.3",
		"body": [
			{
				"type": "Container",
				"items": [
					{
						"type": "TextBlock",
						"text": "¡Hola!",
						"weight": "Bolder",
						"size": "Medium"
					},
					{
						"type": "TextBlock",
						"text": "Esta es la foto del día de la NASA:",
						"wrap": true
					}
				]
			},
			{
				"type": "Image",
				"url": "%s"
			},
			{
				"type": "TextBlock",
				"text": "¿En qué puedo ayudarte?",
				"wrap": true
			}
		]
	}
	`

	response = fmt.Sprintf(response, apodImageURL)
	return response, true
}
