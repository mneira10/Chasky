package azDevopsAPI

import "time"

type PullRequest struct {
	Repository struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		URL     string `json:"url"`
		Project struct {
			ID             string `json:"id"`
			Name           string `json:"name"`
			State          string `json:"state"`
			Visibility     string `json:"visibility"`
			LastUpdateTime string `json:"lastUpdateTime"`
		} `json:"project"`
	} `json:"repository"`
	PullRequestID int    `json:"pullRequestId"`
	CodeReviewID  int    `json:"codeReviewId"`
	Status        string `json:"status"`
	CreatedBy     struct {
		DisplayName string `json:"displayName"`
		URL         string `json:"url"`
		Links       struct {
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"_links"`
		ID         string `json:"id"`
		UniqueName string `json:"uniqueName"`
		ImageURL   string `json:"imageUrl"`
		Descriptor string `json:"descriptor"`
	} `json:"createdBy"`
	CreationDate          time.Time `json:"creationDate"`
	Title                 string    `json:"title"`
	Description           string    `json:"description"`
	SourceRefName         string    `json:"sourceRefName"`
	TargetRefName         string    `json:"targetRefName"`
	MergeStatus           string    `json:"mergeStatus"`
	IsDraft               bool      `json:"isDraft"`
	MergeID               string    `json:"mergeId"`
	LastMergeSourceCommit struct {
		CommitID string `json:"commitId"`
		URL      string `json:"url"`
	} `json:"lastMergeSourceCommit"`
	LastMergeTargetCommit struct {
		CommitID string `json:"commitId"`
		URL      string `json:"url"`
	} `json:"lastMergeTargetCommit"`
	LastMergeCommit struct {
		CommitID string `json:"commitId"`
		URL      string `json:"url"`
	} `json:"lastMergeCommit"`
	Reviewers []struct {
		ReviewerURL string `json:"reviewerUrl"`
		Vote        int    `json:"vote"`
		HasDeclined bool   `json:"hasDeclined"`
		IsFlagged   bool   `json:"isFlagged"`
		DisplayName string `json:"displayName"`
		URL         string `json:"url"`
		Links       struct {
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"_links"`
		ID         string `json:"id"`
		UniqueName string `json:"uniqueName"`
		ImageURL   string `json:"imageUrl"`
	} `json:"reviewers"`
	URL                string `json:"url"`
	SupportsIterations bool   `json:"supportsIterations"`
	AutoCompleteSetBy  struct {
		DisplayName string `json:"displayName"`
		URL         string `json:"url"`
		Links       struct {
			Avatar struct {
				Href string `json:"href"`
			} `json:"avatar"`
		} `json:"_links"`
		ID         string `json:"id"`
		UniqueName string `json:"uniqueName"`
		ImageURL   string `json:"imageUrl"`
		Descriptor string `json:"descriptor"`
	} `json:"autoCompleteSetBy,omitempty"`
	Labels []struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Active bool   `json:"active"`
	} `json:"labels,omitempty"`
}

type PullRequests struct {
	Value []PullRequest `json:"value"`
	Count int           `json:"count"`
}

type Project struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	URL            string    `json:"url"`
	State          string    `json:"state"`
	Revision       int       `json:"revision"`
	Visibility     string    `json:"visibility"`
	LastUpdateTime time.Time `json:"lastUpdateTime"`
	Description    string    `json:"description,omitempty"`
}

type Projects struct {
	Count int       `json:"count"`
	Value []Project `json:"value"`
}
