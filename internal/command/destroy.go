package command

import (
	"fmt"

	"github.com/jamiewri/tfctl/internal/repository"
	"github.com/jamiewri/tfctl/internal/util"
	"github.com/spf13/cobra"
)

func destroyCommand(tfc repository.TerraformCloud) *cobra.Command {

	var tags []string

	c := &cobra.Command{
		Use: "destroy",
		Short: "Start a destroy run, auto-apply disable by default.",
		Run: func(cmd *cobra.Command, args []string) {
			runDestroyCommand(tfc, tags)
		},
	}

    c.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, "")		

	return c
}

func runDestroyCommand(tfc repository.TerraformCloud, tags []string) {

	// Use tags to list to find list of workspaces
	wl, err := tfc.GetWorkspacesFromTags(tags)
	if err != nil {
		fmt.Println(err)
	}

	// Tell the user what workspaces we are targeting.
	util.PrintWorkspaceNames(wl)

	// Run a destroy on every workspace
	for i := range wl.Items {
		tfc.StartDestroy(wl.Items[i])
	}
}
