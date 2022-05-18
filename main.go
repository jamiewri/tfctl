package main

import (
	"os"

	"github.com/jamiewri/tfctl/internal/command"
	"github.com/jamiewri/tfctl/internal/config"
	"github.com/jamiewri/tfctl/internal/repository"
	"github.com/jamiewri/tfctl/internal/tfeclient"
)

var appConfig config.AppConfig

func main () {


    // Saving env args in appConfig
    appConfig.TfcOrg = os.Getenv("TFC_ORGANIZATION")
    appConfig.TfcToken = os.Getenv("TFC_TEAM_TOKEN")

	// Create new tfe client with appConfig
	c := tfeclient.New(&appConfig)
	c.InitClient()

	// Create new instance of Terraform Cloud abstraction layer
	tfc := repository.NewTerraformCloud(c.TfeClient, appConfig)

	// Parse args and run through command tree
	command.Execute(&appConfig, tfc)
}