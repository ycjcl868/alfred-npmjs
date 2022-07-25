package main

import (
	"log"
	"os"
	"os/exec"

	aw "github.com/deanishe/awgo"
)

const UPDATE_JOB_NAME = "checkForUpdate"

func showUpdateStatus() {
	if doCheck {
		wf.Configure(aw.TextErrors(true))
		log.Println("Checking for updates...")
		if err := wf.CheckForUpdate(); err != nil {
			wf.FatalError(err)
		}
		return
	}

	if wf.UpdateCheckDue() && !wf.IsRunning(UPDATE_JOB_NAME) {
		log.Println("Running update check in background...")

		cmd := exec.Command(os.Args[0], "-check")
		if err := wf.RunInBackground(UPDATE_JOB_NAME, cmd); err != nil {
			log.Printf("Error starting update check: %s", err)
		}
	}

	if query != "" {
		return
	}

	if wf.UpdateAvailable() {
		wf.Configure(aw.SuppressUIDs(true))
		log.Println("Update available!")
		wf.NewItem("An new version is available!").
			Subtitle("⇥ or ↩ to install update").
			Valid(false).
			Autocomplete("workflow:update").
			Icon(iconAvailable)
	}
}
