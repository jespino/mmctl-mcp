package main

import (
	"fmt"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

// RoleSystemAdminArgs represents arguments for roles system-admin command
type RoleSystemAdminArgs struct {
	Users []string `json:"users" jsonschema:"required,description=Users to promote to system admin (username, email, or user ID)"`
}

// RoleMemberArgs represents arguments for roles member command
type RoleMemberArgs struct {
	Users []string `json:"users" jsonschema:"required,description=Users to demote to member (username, email, or user ID)"`
}

// RegisterRoleTools registers all role related tools
func RegisterRoleTools(server *mcp_golang.Server) error {
	// Register roles system-admin tool
	err := server.RegisterTool("role_system_admin", "Make users system admins", func(args RoleSystemAdminArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"roles", "system-admin"}
		cmdArgs = append(cmdArgs, args.Users...)
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Users promoted to system admin successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register role_system_admin tool: %v", err)
	}

	// Register roles member tool 
	err = server.RegisterTool("role_member", "Demote users to members", func(args RoleMemberArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"roles", "member"}
		cmdArgs = append(cmdArgs, args.Users...)
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Users demoted to member role successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register role_member tool: %v", err)
	}

	return nil
}
