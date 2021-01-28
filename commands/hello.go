package commands

type helloCommand struct {
}

func (c helloCommand) Execute() {}

func (c helloCommand) GetDescription() string {
	return "Greet humans!"
}

func (c helloCommand) SetCommandInput(command CommandInput) {}

func (c helloCommand) GetOutput() (string, bool) {
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
						"text": "Hello, Human",
						"weight": "Bolder",
						"size": "Medium"
					},
					{
						"type": "TextBlock",
						"text": "Here is a random image of a dog:",
						"wrap": true
					}
				]
			},
			{
				"type": "Image",
				"url": "https://placedog.net/500/280?random"
			},
			{
				"type": "TextBlock",
				"text": "Have a nice day!",
				"wrap": true
			}
		]
	}
	`
	return response, true
}

func (c helloCommand) Help() (string, bool) {
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
							"text": "This function greets humans",
							"weight": "Bolder",
							"size": "Medium"
						},
						{
							"type": "TextBlock",
							"text": "And returns images of dogs. Like this one!",
							"wrap": true
						}
					]
				},
				{
					"type": "Image",
					"url": "https://placedog.net/500/280?random"
				}
			]
		}
	`
	return response, true
}
