package commands

import (
	"fmt"
	nasaAPI "main/nasaAPI"
)

type helloCommand struct{}

func (c helloCommand) execute() {}

func (c helloCommand) getDescription() string {
	return "Saludar!"
}

func (c helloCommand) setCommandInput(command CommandInput) {}

func (c helloCommand) getOutput() (string, bool) {

	apodImageURL := nasaAPI.GetAPODURL()

	response := `
	{
		"$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
		"type": "AdaptiveCard",
		"version": "1.0",
		"body": [
			{
				"type": "Container",
				"items": [
					{
						"type": "TextBlock",
						"text": "¡Hola!",
						"weight": "Bolder",
						"size": "ExtraLarge"
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
				"type": "ActionSet",
				"actions": [
					{
						"type": "Action.OpenUrl",
						"title": "Más información",
						"url": "https://apod.nasa.gov/apod/astropix.html"
					}
				],
				"height": "stretch",
				"horizontalAlignment": "Center"
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
