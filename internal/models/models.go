package models

type WorkspaceList struct {
	Workspaces []Workspace
}

type Workspace struct {
	ID string
	Name string
}

type RunList struct {
	Runs []Run
}

type Run struct {
	ID string
	Status string
}
