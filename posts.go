package main

import (
	"fmt"
	"os/exec"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

// PostCreateArgs represents arguments for post create command
type PostCreateArgs struct {
	Channel  string `json:"channel" jsonschema:"required,description=Channel to post to (in team:channel format for named channels)"`
	Message  string `json:"message" jsonschema:"required,description=Message text to post"`
	ReplyTo  string `json:"replyTo" jsonschema:"description=Post ID to reply to"`
	AsUserID string `json:"asUserId" jsonschema:"description=User ID to post as (impersonation)"`
}

// PostListArgs represents arguments for post list command
type PostListArgs struct {
	Channel string `json:"channel" jsonschema:"required,description=Channel to list posts from (in team:channel format for named channels)"`
	Number  int    `json:"number" jsonschema:"description=Number of posts to list"`
	ShowIDs bool   `json:"showIds" jsonschema:"description=Show post IDs"`
	Since   string `json:"since" jsonschema:"description=List messages posted after a certain time (ISO 8601)"`
}

// PostDeleteArgs represents arguments for post delete command
type PostDeleteArgs struct {
	PostIDs   []string `json:"postIds" jsonschema:"required,description=IDs of posts to delete"`
	Permanent bool     `json:"permanent" jsonschema:"description=Permanently delete the post and its contents"`
}

// RegisterPostTools registers all post related tools
func RegisterPostTools(server *mcp_golang.Server) error {
	// Register post create tool
	err := server.RegisterTool("post_create", "Create a new post", func(args PostCreateArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"post", "create", "--message", args.Message}

		if args.ReplyTo != "" {
			cmdArgs = append(cmdArgs, "--reply-to", args.ReplyTo)
		}

		// Add the channel as the last argument
		cmdArgs = append(cmdArgs, args.Channel)

		var output string
		var err error
		if args.AsUserID != "" {
			// For impersonation, we handle the special case directly
			// Create the command with the --local and --local-user-id flags
			localArgs := append([]string{"--local", "--local-user-id", args.AsUserID}, cmdArgs...)
			cmd := exec.Command("mmctl", localArgs...)
			outputBytes, cmdErr := cmd.CombinedOutput()
			if cmdErr != nil {
				return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v\nOutput: %s", cmdErr, string(outputBytes)))), nil
			}
			output = string(outputBytes)
		} else {
			// Use standard executeMMCTL for normal post creation
			output, err = executeMMCTL(cmdArgs...)
			if err != nil {
				return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
			}
		}
		if output == "" {
			output = "Post created successfully"
		}

		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register post_create tool: %v", err)
	}

	// Register post list tool
	err = server.RegisterTool("post_list", "List posts in a channel", func(args PostListArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"post", "list", args.Channel}

		if args.Number > 0 {
			cmdArgs = append(cmdArgs, "--number", fmt.Sprintf("%d", args.Number))
		}

		if args.ShowIDs {
			cmdArgs = append(cmdArgs, "--show-ids")
		}

		if args.Since != "" {
			cmdArgs = append(cmdArgs, "--since", args.Since)
		}

		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register post_list tool: %v", err)
	}

	// Register post delete tool
	err = server.RegisterTool("post_delete", "Delete posts", func(args PostDeleteArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"post", "delete"}

		if args.Permanent {
			cmdArgs = append(cmdArgs, "--permanent")
		}

		cmdArgs = append(cmdArgs, args.PostIDs...)

		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register post_delete tool: %v", err)
	}

	return nil
}
