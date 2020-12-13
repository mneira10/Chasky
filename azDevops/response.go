package azDevopsAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const newline = "   \n"
const doubleLineBreak = "\n\n\u200C"

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

	bodyData := getRequest(url, user, token)
	return parseProjects(bodyData)
}

func getPullRequests(project string, user string, token string) PullRequests {
	prsUrl := "https://dev.azure.com/imedicalservices/" + project + "/_apis/git/pullrequests?api-version=5.0"

	bodyData := getRequest(prsUrl, user, token)
	return parsePullRequests(bodyData)
}

// GetPRs returns a string with all of the pull requests in iMedical's repositories
func GetPRs() string {
	projectsUrl := "https://dev.azure.com/imedicalservices/_apis/projects/?api-version=5.0"
	token := "ut6dkcefggjykf46o4a25xeypni72mbhqxlsxpgvy63z7q4m5uwa"
	user := "mauricio.neira"

	projects := getProjects(projectsUrl, token, user)
	numPRs := 0

	loc, err := time.LoadLocation("America/Bogota")
	if err != nil {
		panic(err)
	}

	var response string
	for i := 0; i < projects.Count; i++ {
		projectName := projects.Value[i].Name
		prs := getPullRequests(projectName, user, token)

		response += "# __" + projectName + "__" + newline + doubleLineBreak + newline

		for j := 0; j < prs.Count; j++ {
			pr := prs.Value[j]
			response += "__" + pr.Title + "__" + newline + doubleLineBreak + newline

			response += "Created by: " + pr.CreatedBy.DisplayName + newline
			response += "Created at: " + pr.CreationDate.In(loc).Format("2006-01-02T15:04:05-0700") + newline
			response += "Merge status: " + pr.MergeStatus + newline
			response += "Source branch: " + pr.SourceRefName + newline
			response += "Target branch: " + pr.TargetRefName + newline
			response += "URL: " + pr.URL + newline

			response += "Reviewers: " + newline
			for i, reviewer := range pr.Reviewers {
				response += fmt.Sprintf("%d. %s"+newline, i+1, reviewer.DisplayName)
			}

			response += doubleLineBreak + newline

			numPRs++
		}
		response += newline
	}
	response += fmt.Sprintf("%d Total active pull requests", numPRs)
	fmt.Println(response)
	return response

}
