package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/hashicorp/go-tfe"
	"github.com/jamiewri/tfctl/internal/config"
	"github.com/jamiewri/tfctl/internal/models"
)

type TerraformCloud interface {
	GetWorkspacesFromTags([]string) (models.WorkspaceList, error)
	GetRunsFromWorkspace(w models.Workspace) (models.RunList, error)
	StartPlan(models.Workspace)
	StartApply(models.Workspace)
	StartDestroy(models.Workspace)
	CancelRun(r models.Run)
	DiscardRun(r models.Run)
    GetVariableListFromWorkspace(w tfe.Workspace) (*tfe.VariableList)
	GetWorkspaceFromName(ws string) (*tfe.Workspace)
}

type terraformCloud struct {
	tfcClient *tfe.Client
	appConfig config.AppConfig
}

func NewTerraformCloud(tfcClient *tfe.Client, ac config.AppConfig) (TerraformCloud) {
	return &terraformCloud{
		tfcClient: tfcClient,
		appConfig: ac,
	}
}

func (t *terraformCloud) GetRunsFromWorkspace(w models.Workspace) (models.RunList, error) {

	ctx := context.Background()

	listOptions := &tfe.ListOptions{
		PageNumber: 1,
		PageSize: 50,
	}

	runListOptions := &tfe.RunListOptions{
		ListOptions: *listOptions,
	}

	rl, err := t.tfcClient.Runs.List(ctx, w.ID, runListOptions)
	if err != nil {
		fmt.Println(err)
	}

	runs := make([]models.Run, 0)

	for _, r := range rl.Items {
		run := models.Run{}
		run.ID = r.ID
		run.Status = string(r.Status)
		runs = append(runs, run)
	}
	
    return models.RunList{
	    Runs: runs,
    }, nil
}


func (t *terraformCloud) GetWorkspacesFromTags(tags []string) (models.WorkspaceList, error) {

	// Create context
	ctx := context.Background()

	listOptions := &tfe.ListOptions{
		PageNumber: 1,
		PageSize: 50,
	}

	workspaceListOptions := &tfe.WorkspaceListOptions{
		ListOptions: *listOptions,
		Tags: strings.Join(tags, ","),
	}

	wl, err := t.tfcClient.Workspaces.List(ctx, t.appConfig.TfcOrg, workspaceListOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Create slice of models.workspaces
	workspaces := make([]models.Workspace, 0)

    // Translate tfe.Workspace struct into models.workspaceList
	for _, workspace := range wl.Items {
	    ws := models.Workspace{}
		ws.ID = workspace.ID
		ws.Name = workspace.Name
		workspaces = append(workspaces, ws)
	}

	return models.WorkspaceList{
		Workspaces: workspaces,
	}, nil
}

// GetWorkspaceFromName takes a workspace name and returns a workspace struct
func (t *terraformCloud) GetWorkspaceFromName(ws string) (*tfe.Workspace) {
	ctx := context.Background()

	workspace, err := t.tfcClient.Workspaces.Read(ctx, t.appConfig.TfcOrg, ws)
	if err != nil {
	 fmt.Println(err)
    }

	return workspace
} 

// GetVariableListFromWorkspace takes a workspace and retruns a variable list string
func (t *terraformCloud) GetVariableListFromWorkspace(w tfe.Workspace) (*tfe.VariableList) {
	ctx := context.Background()

	listOptions := &tfe.ListOptions{
		PageNumber: 1,
		PageSize: 50,
	}

	variableListOptions := tfe.VariableListOptions{
		ListOptions: *listOptions,
	}

	vl, err := t.tfcClient.Variables.List(ctx, w.ID, &variableListOptions)
	if err != nil {
		fmt.Println(err)
	}
	 return vl
}

// StartPlan starts a plan with auto-apply set to false
func (t *terraformCloud) StartPlan(w models.Workspace) {

	ctx := context.Background()

	o := tfe.RunCreateOptions{
		Workspace: nil,
		Message: tfe.String("Plan only run started from tfctl"),
		AutoApply: tfe.Bool(false),
	}

	o.Workspace = t.GetWorkspaceFromName(w.Name)

    p, err := t.tfcClient.Runs.Create(ctx, o)
    if err != nil {
    	fmt.Println(err)
    	os.Exit(1)
    }

	fmt.Println("Plan only started:", p.Workspace.ID, "-", p.ID)
}

// StartApply starts a plan with auto-apply set to true
func (t *terraformCloud) StartApply(w models.Workspace) {

	ctx := context.Background()

	o := tfe.RunCreateOptions{
		Workspace: nil,
		Message: tfe.String("Apply run started from tfctl"),
		AutoApply: tfe.Bool(true),
	}

	o.Workspace = t.GetWorkspaceFromName(w.Name)

    p, err := t.tfcClient.Runs.Create(ctx, o)
    if err != nil {
    	fmt.Println(err)
    	os.Exit(1)
    }

	fmt.Println("Apply started:", p.Workspace.ID, "-", p.ID)
}


func (t *terraformCloud) StartDestroy(w models.Workspace) {
	ctx := context.Background()

	o := tfe.RunCreateOptions{
		Workspace: nil,
		Message: tfe.String("Destroy run started from tfctl"),
		AutoApply: tfe.Bool(true),
		IsDestroy: tfe.Bool(true),
	}

	o.Workspace = t.GetWorkspaceFromName(w.Name)

    p, err := t.tfcClient.Runs.Create(ctx, o)
    if err != nil {
    	fmt.Println(err)
    	os.Exit(1)
    }

	fmt.Println("Destroy run started:", p.Workspace.ID, "-", p.ID)
}

// CancelRun will cancel the supplied run
func (t *terraformCloud) CancelRun(r models.Run) {

	ctx := context.Background()

    rco := tfe.RunCancelOptions{
    	Comment: tfe.String("Run cancelled via tfctl"),
    }
    err := t.tfcClient.Runs.Cancel(ctx, r.ID, rco)
    if err != nil {
    	fmt.Println(err)
    }

	fmt.Println("Run cancelled:", r.ID)
}

//DiscardRun will discard the supplied run
func (t *terraformCloud) DiscardRun(r models.Run) {
	ctx := context.Background()

	rdo := tfe.RunDiscardOptions{
		Comment: tfe.String("Run discarded via tfctl"),
	}
 	
   	err := t.tfcClient.Runs.Discard(ctx, r.ID, rdo)
   	if err != nil {
   		fmt.Println(err)
   	}

	fmt.Println("Run discarded:", r.ID)
}