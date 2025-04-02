package main

import (
	"fmt"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

// TeamCreateArgs represents arguments for team create command
type TeamCreateArgs struct {
	Name        string `json:"name" jsonschema:"required,description=Team name (lowercase, no spaces)"`
	DisplayName string `json:"displayName" jsonschema:"required,description=Team display name"`
	Email       string `json:"email" jsonschema:"description=Administrator email address"`
	Private     bool   `json:"private" jsonschema:"description=Create a private team"`
}

// TeamSearchArgs represents arguments for team search command
type TeamSearchArgs struct {
	Terms []string `json:"terms" jsonschema:"required,description=Terms to search for"`
}

// TeamModifyArgs represents arguments for team modify command
type TeamModifyArgs struct {
	Team    string `json:"team" jsonschema:"required,description=Team to modify (name or ID)"`
	Private bool   `json:"private" jsonschema:"description=Make the team private"`
	Public  bool   `json:"public" jsonschema:"description=Make the team public"`
}

// TeamRenameArgs represents arguments for team rename command
type TeamRenameArgs struct {
	Team        string `json:"team" jsonschema:"required,description=Team to rename (name or ID)"`
	DisplayName string `json:"displayName" jsonschema:"required,description=New display name"`
}

// RegisterTeamTools registers all team related tools
func RegisterTeamTools(server *mcp_golang.Server) error {
	// Register team create tool
	err := server.RegisterTool("team_create", "Create a new team", func(args TeamCreateArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"team", "create", "--name", args.Name, "--display-name", args.DisplayName}
		
		if args.Email != "" {
			cmdArgs = append(cmdArgs, "--email", args.Email)
		}
		
		if args.Private {
			cmdArgs = append(cmdArgs, "--private")
		}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register team_create tool: %v", err)
	}

	// Register team search tool
	err = server.RegisterTool("team_search", "Search for teams", func(args TeamSearchArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"team", "search"}
		cmdArgs = append(cmdArgs, args.Terms...)
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register team_search tool: %v", err)
	}

	// Register team modify tool
	err = server.RegisterTool("team_modify", "Modify a team", func(args TeamModifyArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"team", "modify", args.Team}
		
		if args.Private {
			cmdArgs = append(cmdArgs, "--private")
		}
		
		if args.Public {
			cmdArgs = append(cmdArgs, "--public")
		}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register team_modify tool: %v", err)
	}

	// Register team rename tool
	err = server.RegisterTool("team_rename", "Rename a team", func(args TeamRenameArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"team", "rename", args.Team, "--display-name", args.DisplayName}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register team_rename tool: %v", err)
	}

	return nil
}
