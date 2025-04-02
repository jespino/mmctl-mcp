package main

import (
	"fmt"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

// OAuthListArgs represents arguments for oauth list command
type OAuthListArgs struct {
	Page    int `json:"page" jsonschema:"description=Page number to fetch"`
	PerPage int `json:"perPage" jsonschema:"description=Number of items per page"`
}

// RegisterOAuthTools registers all oauth related tools
func RegisterOAuthTools(server *mcp_golang.Server) error {
	// Register oauth list tool
	err := server.RegisterTool("oauth_list", "List OAuth2 applications", func(args OAuthListArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"oauth", "list"}
		
		if args.Page > 0 {
			cmdArgs = append(cmdArgs, "--page", fmt.Sprintf("%d", args.Page))
		}
		
		if args.PerPage > 0 {
			cmdArgs = append(cmdArgs, "--per-page", fmt.Sprintf("%d", args.PerPage))
		}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register oauth_list tool: %v", err)
	}

	return nil
}
