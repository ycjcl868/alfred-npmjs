# alfred-npmjs

[Github](https://github.com/ycjcl868/alfred-npmjs)
[中文 README](README-zh_CN.md)

> [Alfred](https://www.alfredapp.com) workflow to search for npm packages with [npmjs.com](https://www.npmjs.com/)

npm is the package manager for JavaScript and the world’s largest software registry. alfred-npmjs can search and reach the package repo page more quickly than npmjs.com

![](https://raw.githubusercontent.com/ycjcl868/alfred-npmjs/gh-pages/images/snapshot.png)

Using `CMD` + `Enter` to copy package `"name": "^version"` to clipboard.

![](https://user-images.githubusercontent.com/13595509/180936962-3fbd3f6f-0dd6-4424-ace5-173794f27cce.png)

## Download and installation
Download the latest version from [Github releases page](https://github.com/ycjcl868/alfred-npmjs/releases/latest) or [packal download](http://www.packal.org/workflow/npmsearch)

## Features
- For accurate search (default show 3 packages, you can modify the max number)
- Display the packages' avator images
- Cache package lists, if the sum of avator images is greater than 10, the images downloaded could be deleted. (default cache 10 images files)

## Usage
In Alfred, type `npm`, <kbd>Space</kbd> , `package` your want to search. (eg: `npm lodash`)

Select a package and press <kbd>Enter</kbd> to go to the package `npm.js` repo.

![](https://raw.githubusercontent.com/ycjcl868/alfred-npmjs/gh-pages/images/usage.gif)

## Contributing

### Run project

The workflow is written in [Go](https://golang.org/) and uses [AwGo](https://github.com/deanishe/awgo) library for all Alfred related things.

It uses [modd](https://github.com/cortesi/modd) to ease its development.

1. Clone repo
2. Run `build.sh` (makes symbolic link of [`workflow`](workflow) directory)
3. Run `modd` (starts a process that automatically builds the workflow with `build.sh` on any changes you make to `.go` files, this builds and places a binary inside [`workflow`](workflow) directory.)
4. Make changes to code or modify Alfred objects to do what you want! Open debugger in Alfred or run the workflow with `workflow:log` passed in as argument to see the logs Alfred produces.

## Changelog

#### v1.1.4

- Support automatic request retry.
- Fix keyword with `/` not searching.

#### v1.1.3

- Fix copy text format

#### v1.1.2
- feat: add `NPM_REGISTRY` environment variable to change registry host

#### v1.1.0
- refactor: using golang
- feat：speed up icon downloading

#### v1.0.4
- fix: search URL not work
#### v1.0.3
- fix: description optional bug

#### v1.0.2
- fix: searching error when input a package name including slash.

#### v1.0.1
- provide a faster search HK Proxy for Chinese User

#### v1.0.0
- init project

## Contributing
[GitHub issues](https://github.com/ycjcl868/alfred-npmjs/issues)
