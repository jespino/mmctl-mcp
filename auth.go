package main

import (
	"fmt"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

// AuthListArgs represents arguments for auth list command
type AuthListArgs struct {
	// No specific args for now
}

// AuthSetArgs represents arguments for auth set command
type AuthSetArgs struct {
	ServerName string `json:"serverName" jsonschema:"required,description=Server name to set as active"`
}

// AuthCurrentArgs represents arguments for auth current command
type AuthCurrentArgs struct {
	// No specific args for now
}

// RegisterAuthTools registers all auth related tools
func RegisterAuthTools(server *mcp_golang.Server) error {
	// Register auth list tool
	err := server.RegisterTool("auth_list", "List stored credentials", func(args AuthListArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"auth", "list"}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Auth credentials listed successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register auth_list tool: %v", err)
	}

	// Register auth set tool
	err = server.RegisterTool("auth_set", "Set active credentials", func(args AuthSetArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"auth", "set", args.ServerName}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Auth credentials set successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register auth_set tool: %v", err)
	}

	// Register auth current tool
	err = server.RegisterTool("auth_current", "Show current credentials", func(args AuthCurrentArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"auth", "current"}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Current auth credentials retrieved successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register auth_current tool: %v", err)
	}

	return nil
}
