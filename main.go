package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	mcp_golang "github.com/metoro-io/mcp-golang"
	"github.com/metoro-io/mcp-golang/transport/stdio"
)

// MMCTLCommand represents arguments for mmctl commands
type MMCTLCommand struct {
	Command string `json:"command" jsonschema:"required,description=The mmctl command to run (e.g. 'user list', 'team list')"`
}

// SystemInfoArgs represents arguments for system info command
type SystemInfoArgs struct {
	Detail bool `json:"detail" jsonschema:"description=Whether to include detailed information"`
}

// UserListArgs represents arguments for user list command
type UserListArgs struct {
	Team     string `json:"team" jsonschema:"description=Filter users by team"`
	Inactive bool   `json:"inactive" jsonschema:"description=Show only inactive users"`
	Page     int    `json:"page" jsonschema:"description=Page number"`
	PerPage  int    `json:"perPage" jsonschema:"description=Number of users per page"`
}

// TeamListArgs represents arguments for team list command
type TeamListArgs struct {
	// No specific args for now
}

// executeMMCTL runs the mmctl command with given arguments
func executeMMCTL(args ...string) (string, error) {
	// Always add --local mode
	localArgs := append([]string{"--local"}, args...)
	cmd := exec.Command("mmctl", localArgs...)
	
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("error executing mmctl: %w\nOutput: %s", err, string(output))
	}
	
	return string(output), nil
}

func main() {
	server := mcp_golang.NewServer(stdio.NewStdioServerTransport())

	// Register a generic mmctl command tool
	err := server.RegisterTool("mmctl", "Run any mmctl command with --local mode", func(args MMCTLCommand) (*mcp_golang.ToolResponse, error) {
		cmdArgs := strings.Fields(args.Command)
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to register mmctl tool: %v\n", err)
		os.Exit(1)
	}

	// Register system version tool
	err = server.RegisterTool("system_info", "Get Mattermost system version information", func(args SystemInfoArgs) (*mcp_golang.ToolResponse, error) {
		var output string
		var err error
		
		if args.Detail {
			output, err = executeMMCTL("system", "status")
		} else {
			output, err = executeMMCTL("system", "version")
		}
		
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to register system_info tool: %v\n", err)
		os.Exit(1)
	}

	// Register user list tool
	err = server.RegisterTool("user_list", "List Mattermost users", func(args UserListArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"user", "list"}
		
		if args.Team != "" {
			cmdArgs = append(cmdArgs, "--team", args.Team)
		}
		
		if args.Inactive {
			cmdArgs = append(cmdArgs, "--inactive")
		}
		
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
		fmt.Fprintf(os.Stderr, "Failed to register user_list tool: %v\n", err)
		os.Exit(1)
	}

	// Register team list tool
	err = server.RegisterTool("team_list", "List Mattermost teams", func(args TeamListArgs) (*mcp_golang.ToolResponse, error) {
		output, err := executeMMCTL("team", "list")
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to register team_list tool: %v\n", err)
		os.Exit(1)
	}

	// Register the new tool categories
	if err := RegisterChannelTools(server); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to register channel tools: %v\n", err)
		os.Exit(1)
	}

	if err := RegisterUserTools(server); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to register user tools: %v\n", err)
		os.Exit(1)
	}

	if err := RegisterPostTools(server); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to register post tools: %v\n", err)
		os.Exit(1)
	}

	if err := RegisterPluginTools(server); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to register plugin tools: %v\n", err)
		os.Exit(1)
	}

	if err := RegisterTeamTools(server); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to register team tools: %v\n", err)
		os.Exit(1)
	}

	if err := RegisterConfigTools(server); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to register config tools: %v\n", err)
		os.Exit(1)
	}

	if err := RegisterWebhookTools(server); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to register webhook tools: %v\n", err)
		os.Exit(1)
	}

	if err := RegisterRoleTools(server); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to register role tools: %v\n", err)
		os.Exit(1)
	}

	if err := RegisterJobTools(server); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to register job tools: %v\n", err)
		os.Exit(1)
	}

	if err := RegisterPermissionTools(server); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to register permission tools: %v\n", err)
		os.Exit(1)
	}

	// Start the server
	err = server.Serve()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to start server: %v\n", err)
		os.Exit(1)
	}

	// Block forever
	select {}
}
