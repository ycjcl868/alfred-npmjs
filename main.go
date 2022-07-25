package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	aw "github.com/deanishe/awgo"
)

var (
	repo = "ycjcl868/alfred-npmjs"

	doCheck       bool
	iconAvailable = &aw.Icon{Value: "update-available.png"}
	query         string
	wf            *aw.Workflow
)

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
	query = wf.Args()[0]

	// showUpdateStatus()
	log.Printf("query: %s\n", query)
	if query != "" {
		wf.ClearData()
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
	flag.BoolVar(&doCheck, "check", false, "check for a new version")
	wf = aw.New()
}

func main() {
	wf.Run(run)
}
