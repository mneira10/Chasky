package azDevopsAPI

import (
	"encoding/json"
	"io/ioutil"
	"log"
	utils "main/utils"
	"net/http"
)

const newline = "   \n"
const doubleLineBreak = "\n\n\u200C"
const projectsURL = "https://dev.azure.com/imedicalservices/_apis/projects/?api-version=5.0"

func getRequest(url string, user string, token string) []byte {
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(user, token)

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	return body
}

func parseProjects(bodyData []byte) Projects {

	var projects Projects
	err := json.Unmarshal(bodyData, &projects)
	if err != nil {
		log.Fatal(err)
	}

	return projects
}

func parsePullRequests(bodyData []byte) PullRequests {
	var pullRequests PullRequests
	err := json.Unmarshal(bodyData, &pullRequests)
	if err != nil {
		log.Fatal(err)
	}

	return pullRequests
}

func getProjects(url string, user string, token string) Projects {
	log.Println("Getting all projects")
	bodyData := getRequest(url, user, token)
	return parseProjects(bodyData)
}

func getPullRequests(project string, user string, token string) PullRequests {
	log.Printf("Getting all prs for project: %s", project)
	prsURL := "https://dev.azure.com/imedicalservices/" + project + "/_apis/git/pullrequests?api-version=5.0"

	bodyData := getRequest(prsURL, user, token)
	return parsePullRequests(bodyData)
}

// GetPRs returns a string with all of the pull requests in iMedical's repositories
func GetPRs() map[Project]PullRequests {

	log.Printf("Getting ALL pull requests")
	token := utils.GetEnvironmentVariable("AZ_DEVOPS_TOKEN")
	user := utils.GetEnvironmentVariable("AZ_DEVOPS_USER")

	projects := getProjects(projectsURL, token, user)

	var projectsToPRs map[Project]PullRequests
	projectsToPRs = make(map[Project]PullRequests)

	// for every project, get the pull requests
	for i := 0; i < projects.Count; i++ {

		project := projects.Value[i]
		projectName := project.Name

		prs := getPullRequests(projectName, user, token)

		projectsToPRs[project] = prs

	}

	log.Println("Succesfully retrieved all pull requests.")
	return projectsToPRs

}
