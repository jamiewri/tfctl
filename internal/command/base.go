package command

import (
	"fmt"
	"os"

	"github.com/jamiewri/tfctl/internal/config"
	"github.com/jamiewri/tfctl/internal/repository"
	"github.com/spf13/cobra"
)

// BaseCmd represents the base command when called without any subcommands
var BaseCommand = &cobra.Command{
	Use:   "tfctl",
	Short: "Terraform Cloud orchestration",
	Long: "This CLI tool is used to orchestrate Terraform Cloud and Terraform Enterprise.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute(appConfig *config.AppConfig, tfc repository.TerraformCloud) {
 
	// Output expected behaviour
	fmt.Println("Searching organization:", appConfig.TfcOrg)

	BaseCommand.AddCommand(searchCommand(tfc))
	BaseCommand.AddCommand(planCommand(tfc))
	BaseCommand.AddCommand(applyCommand(tfc))
	BaseCommand.AddCommand(destroyCommand(tfc))
	BaseCommand.AddCommand(cancelCommand(tfc))

	if err := BaseCommand.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	  }
}