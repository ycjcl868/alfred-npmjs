package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	aw "github.com/deanishe/awgo"
)

type NpmRepoSearchResponse struct {
	Objects []Objects `json:"objects"`
	Total   int       `json:"total"`
	Time    string    `json:"time"`
}
type Links struct {
	Npm        string `json:"npm"`
	Homepage   string `json:"homepage"`
	Repository string `json:"repository"`
	Bugs       string `json:"bugs"`
}
type Publisher struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
type Maintainers struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}
type Detail struct {
	Quality     float64 `json:"quality"`
	Popularity  float64 `json:"popularity"`
	Maintenance float64 `json:"maintenance"`
}

// type Score struct {
// 	Final  float64 `json:"final"`
// 	Detail Detail  `json:"detail"`
// }
type Package struct {
	Name string `json:"name"`
	// Scope       string        `json:"scope"`
	Version     string        `json:"version"`
	Description string        `json:"description"`
	Keywords    []string      `json:"keywords"`
	Date        time.Time     `json:"date"`
	Links       Links         `json:"links"`
	Author      Author        `json:"author"`
	Publisher   Publisher     `json:"publisher"`
	Maintainers []Maintainers `json:"maintainers"`
}
type Author struct {
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
	URL   string `json:"url",omitempty`
}
type Objects struct {
	Package Package `json:"package,omitempty"`
	// Score       Score   `json:"score"`
	// SearchScore float64 `json:"searchScore"`
}

var wf *aw.Workflow

const (
	NPM_API_HOST = "https://registry.npmjs.com"

	PACKAGE_NUM = 5
)

func SearchNpmPackages(keyword string) (NpmRepoSearchResponse, error) {
	url := fmt.Sprintf("%s/-/v1/search?text=%s&from=0&size=%d", NPM_API_HOST, keyword, PACKAGE_NUM)

	repoResp := NpmRepoSearchResponse{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Connection", "Keep-Alive")
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return repoResp, err
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &repoResp)
	if err != nil {
		return repoResp, err
	}

	return repoResp, nil
}

func run() {
	query := wf.Args()[0]

	wf.Configure(aw.SuppressUIDs(true))

	if query != "" {
		resp, _ := SearchNpmPackages(query)
		for _, value := range resp.Objects {
			title := fmt.Sprintf("%s %s", value.Package.Name, value.Package.Version)
			item := wf.NewItem(title).Subtitle(value.Package.Description).Copytext(title).Quicklook(value.Package.Links.Npm).Arg(value.Package.Links.Npm).Valid(true)

			// copy
			copyText := fmt.Sprintf("\"%s\": \"^%s\"", value.Package.Name, value.Package.Version)
			item.Cmd().Arg(copyText).Subtitle("Press Enter to copy this \"name\": \"^version\" to clipboard").Valid(true)
		}
	}

	wf.WarnEmpty("No matching items", "Try a different query?")
	wf.SendFeedback()
}

func init() {
	wf = aw.New()
}

func main() {
	wf.Run(run)
}
