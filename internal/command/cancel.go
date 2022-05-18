package command

import (
	"fmt"

	"github.com/jamiewri/tfctl/internal/repository"
	"github.com/jamiewri/tfctl/internal/util"
	"github.com/spf13/cobra"
)

func cancelCommand(tfc repository.TerraformCloud) *cobra.Command {

	var tags []string

	c := &cobra.Command{
		Use: "cancel",
		Short: "Cancel all run on any workspace that matches the supplied tag.",
		Run: func(cmd *cobra.Command, args []string) {
			runCancelCommand(tfc, tags)
		},
	}

    c.Flags().StringSliceVarP(&tags, "tags", "t", []string{}, "")		

	return c
}


// runCancelComand either discards or cancels run based on status
// Unsure what to do so do nothing.
//  confirmed
//  cost_estimated
//  fetching
//  planned
//  policy_soft_failed
//  post_plan_running
//  post_plan_completed

// No Action
// - applied
// - errored
// - discarded
// - cancelled
// - planned_and_finished


// Cancel
// - planning
// - plan_queued
// - apply_queued
// - pending 
// - cost_estimating
// - applying
// - policy_checking

// Discard
// - policy_checked
// - policy_override


func runCancelCommand(tfc repository.TerraformCloud, tags []string) {

	// Use tags to list to find list of workspaces
	wl, err := tfc.GetWorkspacesFromTags(tags)
	if err != nil {
		fmt.Println(err)
	}

	// Tell the user what workspaces we are targeting.
	util.PrintWorkspaceNames(wl)

	// Run a plan on every workspace
	for i := range wl.Workspaces {
	
		//find all runs in workspace
		rl, err := tfc.GetRunsFromWorkspace(wl.Workspaces[i])
		if err != nil {
			fmt.Println(err)
		}

		for i := range rl.Runs {

			s := rl.Runs[i].Status

            // Discard run
            if s == "policy_checked" ||
			   s == "policy_override" {

			     tfc.DiscardRun(rl.Runs[i])
			     continue

            }
            
            // Cancel run
            if s == "planning" ||
			   s == "plan_queued" ||
			   s == "apply_queued" ||
			   s == "pending" ||
			   s == "cost_estimating" ||
			   s == "applying" ||
			   s == "policy_checking" {

			     tfc.CancelRun(rl.Runs[i])
			     continue

            }
		}
	}
}
