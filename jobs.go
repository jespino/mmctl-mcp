package main

import (
	"fmt"
	"strings"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

// JobListArgs represents arguments for job list command
type JobListArgs struct {
	Page    int      `json:"page" jsonschema:"description=Page number"`
	PerPage int      `json:"perPage" jsonschema:"description=Number of jobs per page"`
	All     bool     `json:"all" jsonschema:"description=Fetch all jobs"`
	JobIDs  []string `json:"jobIds" jsonschema:"description=List of job IDs to filter by"`
	JobType string   `json:"jobType" jsonschema:"description=Filter by job type"`
	Status  string   `json:"status" jsonschema:"description=Filter by job status"`
}

// JobUpdateArgs represents arguments for job update command
type JobUpdateArgs struct {
	JobID  string `json:"jobId" jsonschema:"required,description=ID of the job to update"`
	Status string `json:"status" jsonschema:"required,description=New status for the job (pending, cancel_requested, canceled)"`
	Force  bool   `json:"force" jsonschema:"description=Force the status update, overriding restrictions"`
}

// RegisterJobTools registers all job related tools
func RegisterJobTools(server *mcp_golang.Server) error {
	// Register job list tool
	err := server.RegisterTool("job_list", "List jobs", func(args JobListArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"job", "list"}

		if args.Page > 0 {
			cmdArgs = append(cmdArgs, "--page", fmt.Sprintf("%d", args.Page))
		}

		if args.PerPage > 0 {
			cmdArgs = append(cmdArgs, "--per-page", fmt.Sprintf("%d", args.PerPage))
		}

		if args.All {
			cmdArgs = append(cmdArgs, "--all")
		}

		if len(args.JobIDs) > 0 {
			cmdArgs = append(cmdArgs, "--ids", fmt.Sprintf("%s", strings.Join(args.JobIDs, ",")))
		}

		if args.JobType != "" {
			cmdArgs = append(cmdArgs, "--type", args.JobType)
		}

		if args.Status != "" {
			cmdArgs = append(cmdArgs, "--status", args.Status)
		}

		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Jobs listed successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register job_list tool: %v", err)
	}

	// Register job update tool
	err = server.RegisterTool("job_update", "Update job status", func(args JobUpdateArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"job", "update", args.JobID, args.Status}

		if args.Force {
			cmdArgs = append(cmdArgs, "--force")
		}

		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Job updated successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register job_update tool: %v", err)
	}

	return nil
}
