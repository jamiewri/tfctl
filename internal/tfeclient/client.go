package tfeclient

import (
	"log"

	"github.com/hashicorp/go-tfe"
	"github.com/jamiewri/tfctl/internal/config"
)

type tfeClient struct {
	TfeClient *tfe.Client
	appConfig config.AppConfig
}

func New(config *config.AppConfig) (*tfeClient) {
  return &tfeClient{
	  TfeClient: nil,
	  appConfig: *config,
  }
}

func (c *tfeClient) InitClient() error {

	// Init go-tfe config
	config := &tfe.Config{
		Token: c.appConfig.TfcToken,
	}

	// Create new tfe client
	client, err := tfe.NewClient(config)
	if err != nil {
		log.Fatal(err)
		return err
	}

	c.TfeClient = client
    return nil
}