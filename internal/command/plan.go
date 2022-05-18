package command

import (
	"fmt"

	"github.com/jamiewri/tfctl/internal/repository"
	"github.com/jamiewri/tfctl/internal/util"
	"github.com/spf13/cobra"
)

func planCommand(tfc repository.TerraformCloud) *cobra.Command {

	var tags []string

	c := &cobra.Command{
		Use: "plan",
		Short: "Start a plan only run, auto-apply is disabled.",
		Run: func(cmd *cobra.Command, args []string) {
			runPlanCommand(tfc, tags)
		},
	}

    c.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, "")

	return c
}

func runPlanCommand(tfc repository.TerraformCloud, tags []string) {

	// Use tags to list to find list of workspaces
	wl, err := tfc.GetWorkspacesFromTags(tags)
	if err != nil {
		fmt.Println(err)
	}

	// Tell the user what workspaces we are targeting.
	util.PrintWorkspaceNames(wl)

	// Run a plan on every workspace
	for _, w := range wl.Workspaces {
		tfc.StartPlan(w)
	}
}
