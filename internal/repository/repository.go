package repository

import "github.com/jamiewri/tfctl/internal/config"

type Repository struct {
	App *config.AppConfig
	Tfc TerraformCloud
}