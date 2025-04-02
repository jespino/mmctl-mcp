package main

import (
	"fmt"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

// GroupChannelListArgs represents arguments for group channel list command
type GroupChannelListArgs struct {
	TeamChannel string `json:"teamChannel" jsonschema:"required,description=Team and channel in format 'team:channel'"`
}

// GroupTeamListArgs represents arguments for group team list command
type GroupTeamListArgs struct {
	Team string `json:"team" jsonschema:"required,description=Team name or ID"`
}

// GroupChannelStatusArgs represents arguments for group channel status command
type GroupChannelStatusArgs struct {
	TeamChannel string `json:"teamChannel" jsonschema:"required,description=Team and channel in format 'team:channel'"`
}

// GroupTeamStatusArgs represents arguments for group team status command
type GroupTeamStatusArgs struct {
	Team string `json:"team" jsonschema:"required,description=Team name or ID"`
}

// GroupChannelEnableArgs represents arguments for group channel enable command
type GroupChannelEnableArgs struct {
	TeamChannel string `json:"teamChannel" jsonschema:"required,description=Team and channel in format 'team:channel'"`
}

// GroupChannelDisableArgs represents arguments for group channel disable command
type GroupChannelDisableArgs struct {
	TeamChannel string `json:"teamChannel" jsonschema:"required,description=Team and channel in format 'team:channel'"`
}

// GroupTeamEnableArgs represents arguments for group team enable command
type GroupTeamEnableArgs struct {
	Team string `json:"team" jsonschema:"required,description=Team name or ID"`
}

// GroupTeamDisableArgs represents arguments for group team disable command
type GroupTeamDisableArgs struct {
	Team string `json:"team" jsonschema:"required,description=Team name or ID"`
}

// RegisterGroupTools registers all group related tools
func RegisterGroupTools(server *mcp_golang.Server) error {
	// Register group channel list tool
	err := server.RegisterTool("group_channel_list", "List groups for a channel", func(args GroupChannelListArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"group", "channel", "list", args.TeamChannel}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register group_channel_list tool: %v", err)
	}

	// Register group team list tool
	err = server.RegisterTool("group_team_list", "List groups for a team", func(args GroupTeamListArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"group", "team", "list", args.Team}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register group_team_list tool: %v", err)
	}

	// Register group channel status tool
	err = server.RegisterTool("group_channel_status", "Show group constraint status for a channel", func(args GroupChannelStatusArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"group", "channel", "status", args.TeamChannel}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register group_channel_status tool: %v", err)
	}

	// Register group team status tool
	err = server.RegisterTool("group_team_status", "Show group constraint status for a team", func(args GroupTeamStatusArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"group", "team", "status", args.Team}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register group_team_status tool: %v", err)
	}

	// Register group channel enable tool
	err = server.RegisterTool("group_channel_enable", "Enable group constraints for a channel", func(args GroupChannelEnableArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"group", "channel", "enable", args.TeamChannel}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register group_channel_enable tool: %v", err)
	}

	// Register group channel disable tool
	err = server.RegisterTool("group_channel_disable", "Disable group constraints for a channel", func(args GroupChannelDisableArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"group", "channel", "disable", args.TeamChannel}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register group_channel_disable tool: %v", err)
	}

	// Register group team enable tool
	err = server.RegisterTool("group_team_enable", "Enable group constraints for a team", func(args GroupTeamEnableArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"group", "team", "enable", args.Team}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register group_team_enable tool: %v", err)
	}

	// Register group team disable tool
	err = server.RegisterTool("group_team_disable", "Disable group constraints for a team", func(args GroupTeamDisableArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"group", "team", "disable", args.Team}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register group_team_disable tool: %v", err)
	}

	return nil
}
