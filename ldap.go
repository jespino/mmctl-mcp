package main

import (
	"fmt"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

// LDAPSyncArgs represents arguments for ldap sync command
type LDAPSyncArgs struct {
	IncludeRemovedMembers bool `json:"includeRemovedMembers" jsonschema:"description=Include members who left or were removed from a group-synced team/channel"`
}

// LDAPIDMigrateArgs represents arguments for ldap idmigrate command
type LDAPIDMigrateArgs struct {
	IdAttribute string `json:"idAttribute" jsonschema:"required,description=New ID attribute to migrate to (e.g., 'objectGUID')"`
}

// RegisterLDAPTools registers all LDAP related tools
func RegisterLDAPTools(server *mcp_golang.Server) error {
	// Register ldap sync tool
	err := server.RegisterTool("ldap_sync", "Sync LDAP users and groups", func(args LDAPSyncArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"ldap", "sync"}
		
		if args.IncludeRemovedMembers {
			cmdArgs = append(cmdArgs, "--include-removed-members")
		}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "LDAP sync completed successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register ldap_sync tool: %v", err)
	}

	// Register ldap idmigrate tool
	err = server.RegisterTool("ldap_idmigrate", "Migrate LDAP ID attribute", func(args LDAPIDMigrateArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"ldap", "idmigrate", args.IdAttribute}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "LDAP ID migration completed successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register ldap_idmigrate tool: %v", err)
	}

	return nil
}
