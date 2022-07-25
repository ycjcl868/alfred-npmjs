package main

import "time"

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
