package main

import (
	"fmt"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

// WebhookListArgs represents arguments for webhook list command
type WebhookListArgs struct {
	Team string `json:"team" jsonschema:"description=Team name or ID to filter webhooks by"`
}

// WebhookShowArgs represents arguments for webhook show command
type WebhookShowArgs struct {
	WebhookID string `json:"webhookId" jsonschema:"required,description=ID of the webhook to show"`
}

// WebhookCreateIncomingArgs represents arguments for creating incoming webhooks
type WebhookCreateIncomingArgs struct {
	Channel     string `json:"channel" jsonschema:"required,description=Channel ID"`
	User        string `json:"user" jsonschema:"required,description=User ID (creator)"`
	DisplayName string `json:"displayName" jsonschema:"description=Incoming webhook display name"`
	Description string `json:"description" jsonschema:"description=Incoming webhook description"`
	LockToChannel bool   `json:"lockToChannel" jsonschema:"description=Lock webhook to channel"`
	Icon         string `json:"icon" jsonschema:"description=Icon URL"`
}

// WebhookCreateOutgoingArgs represents arguments for creating outgoing webhooks
type WebhookCreateOutgoingArgs struct {
	Team         string   `json:"team" jsonschema:"required,description=Team name or ID"`
	Channel      string   `json:"channel" jsonschema:"description=Channel name or ID"`
	User         string   `json:"user" jsonschema:"required,description=User username, email, or ID"`
	DisplayName  string   `json:"displayName" jsonschema:"required,description=Outgoing webhook display name"`
	Description  string   `json:"description" jsonschema:"description=Outgoing webhook description"`
	TriggerWords []string `json:"triggerWords" jsonschema:"required,description=Words to trigger webhook"`
	TriggerWhen  string   `json:"triggerWhen" jsonschema:"description=When to trigger webhook (exact or start)"`
	Urls         []string `json:"urls" jsonschema:"required,description=Callback URLs"`
	ContentType  string   `json:"contentType" jsonschema:"description=Content-type for the webhook"`
	Icon         string   `json:"icon" jsonschema:"description=Icon URL"`
}

// WebhookDeleteArgs represents arguments for webhook delete command
type WebhookDeleteArgs struct {
	WebhookID string `json:"webhookId" jsonschema:"required,description=ID of the webhook to delete"`
}

// RegisterWebhookTools registers all webhook related tools
func RegisterWebhookTools(server *mcp_golang.Server) error {
	// Register webhook list tool
	err := server.RegisterTool("webhook_list", "List webhooks", func(args WebhookListArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"webhook", "list"}
		
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
		return fmt.Errorf("failed to register webhook_list tool: %v", err)
	}

	// Register webhook show tool
	err = server.RegisterTool("webhook_show", "Show webhook details", func(args WebhookShowArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"webhook", "show", args.WebhookID}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register webhook_show tool: %v", err)
	}

	// Register webhook create-incoming tool
	err = server.RegisterTool("webhook_create_incoming", "Create incoming webhook", func(args WebhookCreateIncomingArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"webhook", "create-incoming", "--channel", args.Channel, "--user", args.User}
		
		if args.DisplayName != "" {
			cmdArgs = append(cmdArgs, "--display-name", args.DisplayName)
		}
		
		if args.Description != "" {
			cmdArgs = append(cmdArgs, "--description", args.Description)
		}
		
		if args.LockToChannel {
			cmdArgs = append(cmdArgs, "--lock-to-channel")
		}
		
		if args.Icon != "" {
			cmdArgs = append(cmdArgs, "--icon", args.Icon)
		}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register webhook_create_incoming tool: %v", err)
	}

	// Register webhook create-outgoing tool
	err = server.RegisterTool("webhook_create_outgoing", "Create outgoing webhook", func(args WebhookCreateOutgoingArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"webhook", "create-outgoing", "--team", args.Team, "--user", args.User, "--display-name", args.DisplayName}
		
		if args.Channel != "" {
			cmdArgs = append(cmdArgs, "--channel", args.Channel)
		}
		
		if args.Description != "" {
			cmdArgs = append(cmdArgs, "--description", args.Description)
		}
		
		for _, word := range args.TriggerWords {
			cmdArgs = append(cmdArgs, "--trigger-word", word)
		}
		
		if args.TriggerWhen != "" {
			cmdArgs = append(cmdArgs, "--trigger-when", args.TriggerWhen)
		}
		
		for _, url := range args.Urls {
			cmdArgs = append(cmdArgs, "--url", url)
		}
		
		if args.ContentType != "" {
			cmdArgs = append(cmdArgs, "--content-type", args.ContentType)
		}
		
		if args.Icon != "" {
			cmdArgs = append(cmdArgs, "--icon", args.Icon)
		}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register webhook_create_outgoing tool: %v", err)
	}

	// Register webhook delete tool
	err = server.RegisterTool("webhook_delete", "Delete webhook", func(args WebhookDeleteArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"webhook", "delete", args.WebhookID}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register webhook_delete tool: %v", err)
	}

	return nil
}
