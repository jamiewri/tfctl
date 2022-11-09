package command

import (
	"fmt"

	"github.com/jamiewri/tfctl/internal/models"
	"github.com/jamiewri/tfctl/internal/repository"
	"github.com/spf13/cobra"
)

func variablesComand(tfc repository.TerraformCloud) *cobra.Command {

	var w []string

	c := &cobra.Command{
		Use: "variables",
		Short: "List workspaces that match the supplied tags.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Variable handling coming soon..")
			//runListCommand(tfc, w)
		},
	}

    c.Flags().StringSliceVarP(&w, "workspace", "w", []string{}, "")		

	return c
}

func runListCommand(tfc repository.TerraformCloud, w []string) (*models.VariableList){

	// convert ist of strings into tfe.workspace
	ws := tfc.GetWorkspaceFromName(w[0])

	// pass models.Workspace to tfc
	variablesList := tfc.GetVariableListFromWorkspace(*ws)

	for i := range variablesList.Items {
		fmt.Println(variablesList.Items[i].Key)
	}
	return &models.VariableList{}
}