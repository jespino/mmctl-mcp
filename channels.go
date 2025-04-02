package main

import (
	"fmt"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

// ChannelListArgs represents arguments for channel list command
type ChannelListArgs struct {
	Team string `json:"team" jsonschema:"description=Team name or ID to filter channels by"`
}

// ChannelCreateArgs represents arguments for channel create command
type ChannelCreateArgs struct {
	Team        string `json:"team" jsonschema:"required,description=Team name or ID"`
	Name        string `json:"name" jsonschema:"required,description=Channel name (lowercase, no spaces)"`
	DisplayName string `json:"displayName" jsonschema:"required,description=Channel display name"`
	Header      string `json:"header" jsonschema:"description=Channel header"`
	Purpose     string `json:"purpose" jsonschema:"description=Channel purpose"`
	Private     bool   `json:"private" jsonschema:"description=Create a private channel"`
}

// ChannelSearchArgs represents arguments for channel search command
type ChannelSearchArgs struct {
	Team    string `json:"team" jsonschema:"description=Team name or ID to search in"`
	Channel string `json:"channel" jsonschema:"required,description=Channel name to search for"`
}

// ChannelArchiveArgs represents arguments for channel archive command
type ChannelArchiveArgs struct {
	Channel string `json:"channel" jsonschema:"required,description=Channel name or ID to archive (in team:channel format for named channels)"`
}

// ChannelUnarchiveArgs represents arguments for channel unarchive command
type ChannelUnarchiveArgs struct {
	Channel string `json:"channel" jsonschema:"required,description=Channel name or ID to unarchive (in team:channel format for named channels)"`
}

// RegisterChannelTools registers all channel related tools
func RegisterChannelTools(server *mcp_golang.Server) error {
	// Register channel list tool
	err := server.RegisterTool("channel_list", "List channels in a team", func(args ChannelListArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"channel", "list"}

		if args.Team != "" {
			cmdArgs = append(cmdArgs, args.Team)
		}

		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register channel_list tool: %v", err)
	}

	// Register channel create tool
	err = server.RegisterTool("channel_create", "Create a new channel", func(args ChannelCreateArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"channel", "create", "--team", args.Team, "--name", args.Name, "--display-name", args.DisplayName}

		if args.Header != "" {
			cmdArgs = append(cmdArgs, "--header", args.Header)
		}

		if args.Purpose != "" {
			cmdArgs = append(cmdArgs, "--purpose", args.Purpose)
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
		return fmt.Errorf("failed to register channel_create tool: %v", err)
	}

	// Register channel search tool
	err = server.RegisterTool("channel_search", "Search for a channel", func(args ChannelSearchArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"channel", "search"}

		if args.Team != "" {
			cmdArgs = append(cmdArgs, "--team", args.Team)
		}

		cmdArgs = append(cmdArgs, args.Channel)

		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register channel_search tool: %v", err)
	}

	// Register channel archive tool
	err = server.RegisterTool("channel_archive", "Archive a channel", func(args ChannelArchiveArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"channel", "archive", args.Channel}

		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register channel_archive tool: %v", err)
	}

	// Register channel unarchive tool
	err = server.RegisterTool("channel_unarchive", "Unarchive a channel", func(args ChannelUnarchiveArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"channel", "unarchive", args.Channel}

		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register channel_unarchive tool: %v", err)
	}

	return nil
}
