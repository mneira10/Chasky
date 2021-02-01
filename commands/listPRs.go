package commands

import (
	azDevops "main/azDevops"
)

type listPRsCommand struct {
	prs map[azDevops.Project]azDevops.PullRequests
}

func (listPRsCommand) setCommandInput(commInput CommandInput) {}

func (listPRsCommand) getDescription() string {
	return "Retorna todos los pull reuests activos en los proyectos de iMedical."
}

func (env listPRsCommand) execute() {
	env.prs = azDevops.GetPRs()
}

func (env listPRsCommand) getOutput() (string, bool) {
	response := `{
		"$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
		"type": "AdaptiveCard",
		"version": "1.3",
		"body": [
			{
				"type": "Container",
				"items": [
					{
						"type": "TextBlock",
						"text": "Projecto1",
						"wrap": true,
						"weight": "Bolder",
						"size": "ExtraLarge"
					}
					,{
						"type": "Container",
						"items": [
							{
								"type": "TextBlock",
								"text": "PR1",
								"wrap": true,
								"size": "Large",
								"weight": "Bolder"
							},
							{
								"type": "TextBlock",
								"text": "Una descipción muy larga del PR. Es Larga pero informativa para darles a los .. aslkdalskdj alskdn alksdj aslkd ",
								"wrap": true
							},
							{
								"type": "FactSet",
								"facts": [
									{
										"title": "Creado por:",
										"value": "Mauricio Neira"
									},
									{
										"title": "Fecha de creación:",
										"value": "21-02-12"
									},
									{
										"title": "Merge status:",
										"value": "Succeded"
									},
									{
										"title": "Source branch:",
										"value": "alguna-rama.com"
									},
									{
										"title": "Target branch:",
										"value": "alguna-otra-rama.com"
									}
								]
							},
							{
								"type": "TextBlock",
								"text": "Reviewers",
								"wrap": true,
								"size": "Medium",
								"weight": "Bolder"
							},
							{
								"type": "TextBlock",
								"text": "- Persona 1\n- Persona 2\n- Persona 3",
								"wrap": true
							},
							{
								"type": "ActionSet",
								"actions": [
									{
										"type": "Action.OpenUrl",
										"title": "Abrir PR"
									}
								]
							}
						]
					}
				]
			}
		]
	}`

	return response, true
}
