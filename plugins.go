package main

import (
	"fmt"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

// PluginListArgs represents arguments for plugin list command
type PluginListArgs struct {
	// No specific args for now
}

// PluginEnableArgs represents arguments for plugin enable command
type PluginEnableArgs struct {
	Plugins []string `json:"plugins" jsonschema:"required,description=Plugin IDs to enable"`
}

// PluginDisableArgs represents arguments for plugin disable command
type PluginDisableArgs struct {
	Plugins []string `json:"plugins" jsonschema:"required,description=Plugin IDs to disable"`
}

// PluginMarketplaceListArgs represents arguments for marketplace list command
type PluginMarketplaceListArgs struct {
	Filter    string `json:"filter" jsonschema:"description=Filter plugins by ID, name or description"`
	Page      int    `json:"page" jsonschema:"description=Page number to fetch"`
	PerPage   int    `json:"perPage" jsonschema:"description=Number of plugins per page"`
	LocalOnly bool   `json:"localOnly" jsonschema:"description=Only list local plugins"`
	All       bool   `json:"all" jsonschema:"description=Fetch all plugins"`
}

// RegisterPluginTools registers all plugin related tools
func RegisterPluginTools(server *mcp_golang.Server) error {
	// Register plugin list tool
	err := server.RegisterTool("plugin_list", "List installed plugins", func(args PluginListArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"plugin", "list"}

		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register plugin_list tool: %v", err)
	}

	// Register plugin enable tool
	err = server.RegisterTool("plugin_enable", "Enable plugins", func(args PluginEnableArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"plugin", "enable"}
		cmdArgs = append(cmdArgs, args.Plugins...)

		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register plugin_enable tool: %v", err)
	}

	// Register plugin disable tool
	err = server.RegisterTool("plugin_disable", "Disable plugins", func(args PluginDisableArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"plugin", "disable"}
		cmdArgs = append(cmdArgs, args.Plugins...)

		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register plugin_disable tool: %v", err)
	}

	// Register plugin marketplace list tool
	err = server.RegisterTool("plugin_marketplace_list", "List marketplace plugins", func(args PluginMarketplaceListArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"plugin", "marketplace", "list"}

		if args.Filter != "" {
			cmdArgs = append(cmdArgs, "--filter", args.Filter)
		}

		if args.Page > 0 {
			cmdArgs = append(cmdArgs, "--page", fmt.Sprintf("%d", args.Page))
		}

		if args.PerPage > 0 {
			cmdArgs = append(cmdArgs, "--per-page", fmt.Sprintf("%d", args.PerPage))
		}

		if args.LocalOnly {
			cmdArgs = append(cmdArgs, "--local-only")
		}

		if args.All {
			cmdArgs = append(cmdArgs, "--all")
		}

		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register plugin_marketplace_list tool: %v", err)
	}

	return nil
}
