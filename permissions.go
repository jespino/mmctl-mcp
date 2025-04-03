package main

import (
	"fmt"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

// PermissionAddArgs represents arguments for permissions add command
type PermissionAddArgs struct {
	Role       string   `json:"role" jsonschema:"required,description=Role to add permissions to"`
	Permissions []string `json:"permissions" jsonschema:"required,description=Permissions to add to the role"`
}

// PermissionRemoveArgs represents arguments for permissions remove command
type PermissionRemoveArgs struct {
	Role       string   `json:"role" jsonschema:"required,description=Role to remove permissions from"`
	Permissions []string `json:"permissions" jsonschema:"required,description=Permissions to remove from the role"`
}

// PermissionResetArgs represents arguments for permissions reset command
type PermissionResetArgs struct {
	Role string `json:"role" jsonschema:"required,description=Role to reset permissions for"`
}

// RegisterPermissionTools registers all permission related tools
func RegisterPermissionTools(server *mcp_golang.Server) error {
	// Register permissions add tool
	err := server.RegisterTool("permission_add", "Add permissions to a role", func(args PermissionAddArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"permissions", "add", args.Role}
		cmdArgs = append(cmdArgs, args.Permissions...)
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Permissions added successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register permission_add tool: %v", err)
	}

	// Register permissions remove tool
	err = server.RegisterTool("permission_remove", "Remove permissions from a role", func(args PermissionRemoveArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"permissions", "remove", args.Role}
		cmdArgs = append(cmdArgs, args.Permissions...)
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Permissions removed successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register permission_remove tool: %v", err)
	}

	// Register permissions reset tool
	err = server.RegisterTool("permission_reset", "Reset permissions for a role", func(args PermissionResetArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"permissions", "reset", args.Role}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Permissions reset successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register permission_reset tool: %v", err)
	}

	return nil
}
