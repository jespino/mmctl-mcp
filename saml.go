package main

import (
	"fmt"
	"strings"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

// SAMLAuthDataResetArgs represents arguments for saml auth-data-reset command
type SAMLAuthDataResetArgs struct {
	IncludeDeleted bool     `json:"includeDeleted" jsonschema:"description=Include deleted users"`
	DryRun         bool     `json:"dryRun" jsonschema:"description=Perform a dry run without making changes"`
	Users          []string `json:"users" jsonschema:"description=Comma-separated list of user IDs to reset"`
	Yes            bool     `json:"yes" jsonschema:"description=Skip confirmation"`
}

// RegisterSAMLTools registers all SAML related tools
func RegisterSAMLTools(server *mcp_golang.Server) error {
	// Register saml auth-data-reset tool
	err := server.RegisterTool("saml_auth_data_reset", "Reset SAML AuthData field to email", func(args SAMLAuthDataResetArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"saml", "auth-data-reset"}

		if args.IncludeDeleted {
			cmdArgs = append(cmdArgs, "--include-deleted")
		}

		if args.DryRun {
			cmdArgs = append(cmdArgs, "--dry-run")
		}

		if len(args.Users) > 0 {
			cmdArgs = append(cmdArgs, "--users", fmt.Sprintf("%s", strings.Join(args.Users, ",")))
		}

		if args.Yes {
			cmdArgs = append(cmdArgs, "--yes")
		}

		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register saml_auth_data_reset tool: %v", err)
	}

	return nil
}
