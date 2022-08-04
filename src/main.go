package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/imroc/req/v3"
)

var (
	npmApiHost = "https://registry.npmjs.com"
	repo       = "ycjcl868/alfred-npmjs"

	doCheck       bool
	iconAvailable = &aw.Icon{Value: "update-available.png"}
	query         string
	wf            *aw.Workflow
)

const (
	PACKAGE_NUM        = 5
	VARIABLES_REGISTRY = "NPM_REGISTRY"
)

func SearchNpmPackages(keyword string) (NpmRepoSearchResponse, error) {
	repoResp := NpmRepoSearchResponse{}
	client := req.C()

	resp, err := client.R().
		SetRetryCount(2).
		SetResult(&repoResp).
		SetQueryParams(map[string]string{
			"text": keyword,
			"from": "0",
			"size": fmt.Sprint(PACKAGE_NUM),
		}).
		Get(fmt.Sprintf("%s/-/v1/search", npmApiHost))

	log.Println(resp.Request.URL)
	log.Println(err)

	if err != nil {
		return repoResp, err
	}

	if !resp.IsSuccess() {
		log.Println("bad response status", resp.Status)
	}

	return repoResp, nil
}

func run() {
	registryFromEnv, _ := wf.Config.Env.Lookup(VARIABLES_REGISTRY)

	if registryFromEnv != "" {
		npmApiHost = strings.TrimSuffix(registryFromEnv, "/")
	}

	log.Printf("npmApiHost: %s", npmApiHost)

	wf.Args() // call to handle magic actions
	flag.Parse()
	query = flag.Arg(0)

	log.Printf("query: %s\n", query)
	if query != "" {
		resp, _ := SearchNpmPackages(query)
		for _, value := range resp.Objects {
			copyText := fmt.Sprintf("\"%s\": \"^%s\"", value.Package.Name, value.Package.Version)
			title := fmt.Sprintf("%s %s", value.Package.Name, value.Package.Version)
			item := wf.NewItem(title).Subtitle(value.Package.Description).Copytext(copyText).Quicklook(value.Package.Links.Npm).Arg(value.Package.Links.Npm).Valid(true)

			// copy
			item.Cmd().Arg(copyText).Subtitle("Press Enter to copy this \"name\": \"^version\" to clipboard").Valid(true)
		}
	}

	if query != "" {
		wf.Filter(query)
	}

	wf.WarnEmpty("No matching items", "Try a different query?")
	wf.SendFeedback()
}

func init() {
	flag.BoolVar(&doCheck, "check", false, "check for a new version")
	wf = aw.New()
	// wf = aw.New(update.GitHub(repo), aw.HelpURL(repo+"/issues"))
}

func main() {
	wf.Run(run)
}
