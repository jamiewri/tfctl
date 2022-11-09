package command

import (
	"fmt"

	"github.com/jamiewri/tfctl/internal/repository"
	"github.com/jamiewri/tfctl/internal/util"
	"github.com/spf13/cobra"
)

func applyCommand(tfc repository.TerraformCloud) *cobra.Command {

	var tags []string

	c := &cobra.Command{
		Use: "apply",
		Short: "Start an apply run, this run will automatically apply.",
		Run: func(cmd *cobra.Command, args []string) {
			runApplyCommand(tfc, tags)
		},
	}

    c.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, "")		

	return c
}

func runApplyCommand(tfc repository.TerraformCloud, tags []string) {

	// Use tags to list to find list of workspaces
	wl, err := tfc.GetWorkspacesFromTags(tags)
	if err != nil {
		fmt.Println(err)
	}

	// Tell the user what workspaces we are targeting.
	util.PrintWorkspaceNames(wl)

	// Run a plan on every workspace
	for i := range wl.Items {
		tfc.StartApply(wl.Items[i])
	}
}
