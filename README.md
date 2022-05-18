# tfctl

## Purpose
This CLI tool is used to orchestrate Terraform Cloud runs. Its goal is not to cover the entire administation of Terraform Cloud, but instead to solve a very simple workflow. 

That workflow is to trigger common actions against workspaces by referencing tags only.

For example, I would like to create a destroy run on all workspace that match the tag 'demo', or I would like to cancel all runs on workspaces that match the tags 'dev'.

If you are looking for a complete CLI tool that orchestrates with Terraform Cloud, I'd recommend that you check out [tecli](https://github.com/awslabs/tecli) and [tfx](https://github.com/straubt1/tfx).

## Requirements
Required environment variables
| Env | Description |
|-----|-------------|
| TFC_ORGANIZATION | The name of your Terraform Cloud organization |
| TFC_TEAM_TOKEN | [Terraform Cloud team token](https://www.terraform.io/cloud-docs/users-teams-organizations/users#api-tokens) |


## Example Usage
List all workspaces that match the tag `test`.
```bash
$ tfctl search -t test
Searching organization: devopstower
tfc-aws-network-dev
tfc-aws-virtual-machine-dev
tfc-aws-virtual-machine-prod
tfc-aws-network-prod
```

Run a plan on all workspaces that match the tag `test`.
```bash
$ tfctl plan -t test
Searching organization: devopstower
Targeting the following workspaces [tfc-aws-network-dev]
Plan only started: ws-BvDUwoSB4cELmfmc - run-pt5NodM5GTckfMor
```

Run a destroy on all workspaces that match the tag `test`.
```bash
$ tfctl destroy -t test
Searching organization: devopstower
Targeting the following workspaces [tfc-aws-network-dev]
Destroy run started: ws-BvDUwoSB4cELmfmc - run-kF62JTa21eqQUdjt
```

Cancel (or discard) all runs on all workspaces that match both tags `azure` and `demo`.
```bash
$ tfctl cancel -t azure,demo
Searching organization: devopstower
Targeting the following workspaces [tfc-azure-webapp]
Run cancelled: run-Z7kVEntjzwz3ddMw
Run discarded: run-KUyRJxwJE5fWxEWX
```

## Help Menu
```bash
$ tfctl --help
Usage:
  tfctl [command]

Available Commands:
  apply       Start an apply run, this run will automatically apply.
  cancel      Cancel all run on any workspace that matches the supplied tag.
  completion  Generate the autocompletion script for the specified shell
  destroy     Start a destroy run, auto-apply disable by default.
  help        Help about any command
  plan        Start a plan only run, auto-apply disable by default.
  search      List workspaces that match the supplied tags.
```

## Build
```bash
go build -o tfctl .
```