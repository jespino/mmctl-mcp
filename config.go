package main

import (
	"fmt"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

// ConfigGetArgs represents arguments for config get command
type ConfigGetArgs struct {
	Path string `json:"path" jsonschema:"required,description=Configuration setting path in dot notation (e.g., 'SqlSettings.DriverName')"`
}

// ConfigSetArgs represents arguments for config set command
type ConfigSetArgs struct {
	Path   string   `json:"path" jsonschema:"required,description=Configuration setting path in dot notation (e.g., 'SqlSettings.DriverName')"`
	Values []string `json:"values" jsonschema:"required,description=Value(s) to set for the configuration setting"`
}

// ConfigShowArgs represents arguments for config show command
type ConfigShowArgs struct {
	// No specific args for now
}

// RegisterConfigTools registers all configuration related tools
func RegisterConfigTools(server *mcp_golang.Server) error {
	// Register config get tool
	err := server.RegisterTool("config_get", "Get a configuration setting", func(args ConfigGetArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"config", "get", args.Path}

		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register config_get tool: %v", err)
	}

	// Register config set tool
	err = server.RegisterTool("config_set", "Set a configuration setting", func(args ConfigSetArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"config", "set", args.Path}
		cmdArgs = append(cmdArgs, args.Values...)

		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register config_set tool: %v", err)
	}

	// Register config show tool
	err = server.RegisterTool("config_show", "Show the server configuration", func(args ConfigShowArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"config", "show"}

		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register config_show tool: %v", err)
	}

	return nil
}
