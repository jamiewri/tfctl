package command

import (
	"fmt"
	"log"

	"github.com/hashicorp/go-tfe"
	"github.com/jamiewri/tfctl/internal/repository"
	"github.com/spf13/cobra"
)

func searchCommand(tfc repository.TerraformCloud) *cobra.Command {

	var tags []string

	c := &cobra.Command{
		Use: "search",
		Short: "List workspaces that match the supplied tags.",
		Run: func(cmd *cobra.Command, args []string) {
			runSearchCommand(tfc, tags)
		},
	}

    c.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, "")		

	return c
}

func runSearchCommand(tfc repository.TerraformCloud, tags []string) (*tfe.WorkspaceList){

	// Search for matching workspaces
	workspacesList, err := tfc.GetWorkspacesFromTags(tags)
	if err != nil {
		log.Fatal(err)
	}

	// Output matching workspaces
	for i := range workspacesList.Items {
		fmt.Println(workspacesList.Items[i].Name)
	}
	return workspacesList
}
