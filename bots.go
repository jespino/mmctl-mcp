package main

import (
	"fmt"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

// BotListArgs represents arguments for bot list command
type BotListArgs struct {
	All      bool `json:"all" jsonschema:"description=Include all bots (including deleted and orphaned)"`
	Orphaned bool `json:"orphaned" jsonschema:"description=Only show orphaned bots"`
}

// BotCreateArgs represents arguments for bot create command
type BotCreateArgs struct {
	Username    string `json:"username" jsonschema:"required,description=Username for the new bot"`
	DisplayName string `json:"displayName" jsonschema:"description=Display name for the bot"`
	Description string `json:"description" jsonschema:"description=Description for the bot"`
	WithToken   bool   `json:"withToken" jsonschema:"description=Auto-generate access token for the bot"`
}

// BotAssignArgs represents arguments for bot assign command
type BotAssignArgs struct {
	Bot       string `json:"bot" jsonschema:"required,description=Bot username to assign"`
	NewOwner  string `json:"newOwner" jsonschema:"required,description=New owner username"`
}

// BotDisableArgs represents arguments for bot disable command
type BotDisableArgs struct {
	Bot string `json:"bot" jsonschema:"required,description=Bot username to disable"`
}

// BotEnableArgs represents arguments for bot enable command
type BotEnableArgs struct {
	Bot string `json:"bot" jsonschema:"required,description=Bot username to enable"`
}

// RegisterBotTools registers all bot related tools
func RegisterBotTools(server *mcp_golang.Server) error {
	// Register bot list tool
	err := server.RegisterTool("bot_list", "List bots", func(args BotListArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"bot", "list"}
		
		if args.All {
			cmdArgs = append(cmdArgs, "--all")
		}
		
		if args.Orphaned {
			cmdArgs = append(cmdArgs, "--orphaned")
		}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Bots listed successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register bot_list tool: %v", err)
	}

	// Register bot create tool
	err = server.RegisterTool("bot_create", "Create a new bot", func(args BotCreateArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"bot", "create", args.Username}
		
		if args.DisplayName != "" {
			cmdArgs = append(cmdArgs, "--display-name", args.DisplayName)
		}
		
		if args.Description != "" {
			cmdArgs = append(cmdArgs, "--description", args.Description)
		}
		
		if args.WithToken {
			cmdArgs = append(cmdArgs, "--with-token")
		}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Bot created successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register bot_create tool: %v", err)
	}

	// Register bot assign tool
	err = server.RegisterTool("bot_assign", "Assign a bot to a new owner", func(args BotAssignArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"bot", "assign", args.Bot, args.NewOwner}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Bot assigned to new owner successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register bot_assign tool: %v", err)
	}

	// Register bot disable tool
	err = server.RegisterTool("bot_disable", "Disable a bot", func(args BotDisableArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"bot", "disable", args.Bot}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Bot disabled successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register bot_disable tool: %v", err)
	}

	// Register bot enable tool
	err = server.RegisterTool("bot_enable", "Enable a bot", func(args BotEnableArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"bot", "enable", args.Bot}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Bot enabled successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register bot_enable tool: %v", err)
	}

	return nil
}
