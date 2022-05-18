package util

import (
	"fmt"

	"github.com/jamiewri/tfctl/internal/models"
)

// PrintWorkspaceNamestags a workspace list and prints the name of each workspace to the console.
func PrintWorkspaceNames(wl models.WorkspaceList) {
	tw := make([]string, 0)
    for _, ws := range wl.Workspaces {
    	tw = append(tw, ws.Name)
    }
    fmt.Println("Targeting the following workspaces", tw)
}


func PrintRunIDs(rl models.RunList) {
	tr := make([]string, 0)
    for _, r := range rl.Runs {
    	tr = append(tr, r.ID)
		fmt.Println(r.Status)
    }
    fmt.Println("Targeting the following runs", tr)
}