package util

import (
	"fmt"

	"github.com/hashicorp/go-tfe"
)

// PrintWorkspaceNamestags a workspace list and prints the name of each workspace to the console.
func PrintWorkspaceNames(wl *tfe.WorkspaceList) {
	tw := make([]string, 0)
    for _, ws := range wl.Items {
    	tw = append(tw, ws.Name)
    }
    fmt.Println("Targeting the following workspaces", tw)
}


func PrintRunIDs(rl *tfe.RunList) {
	tr := make([]string, 0)
    for _, r := range rl.Items {
    	tr = append(tr, r.ID)
		fmt.Println(r.Status)
    }
    fmt.Println("Targeting the following runs", tr)
}