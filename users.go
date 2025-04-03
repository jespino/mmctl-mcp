package main

import (
	"fmt"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

// UserSearchArgs represents arguments for user search command
type UserSearchArgs struct {
	Terms []string `json:"terms" jsonschema:"required,description=Terms to search for (email, username, or user ID)"`
}

// UserCreateArgs represents arguments for user create command
type UserCreateArgs struct {
	Email           string `json:"email" jsonschema:"required,description=Email address for the user"`
	Username        string `json:"username" jsonschema:"required,description=Username for the user"`
	Password        string `json:"password" jsonschema:"required,description=Password for the user"`
	FirstName       string `json:"firstName" jsonschema:"description=First name for the user"`
	LastName        string `json:"lastName" jsonschema:"description=Last name for the user"`
	Nickname        string `json:"nickname" jsonschema:"description=Nickname for the user"`
	Locale          string `json:"locale" jsonschema:"description=Locale (e.g., en, fr) for the user"`
	SystemAdmin     bool   `json:"systemAdmin" jsonschema:"description=Whether to make the user a system admin"`
	EmailVerified   bool   `json:"emailVerified" jsonschema:"description=Whether to mark the email as verified"`
	Guest           bool   `json:"guest" jsonschema:"description=Whether to create as a guest user"`
	DisableWelcomeEmail bool `json:"disableWelcomeEmail" jsonschema:"description=Whether to disable the welcome email"`
}

// UserActivateArgs represents arguments for user activate command
type UserActivateArgs struct {
	Users []string `json:"users" jsonschema:"required,description=Users to activate (email, username, or ID)"`
}

// UserDeactivateArgs represents arguments for user deactivate command
type UserDeactivateArgs struct {
	Users []string `json:"users" jsonschema:"required,description=Users to deactivate (email, username, or ID)"`
}

// UserEmailArgs represents arguments for changing user email
type UserEmailArgs struct {
	User     string `json:"user" jsonschema:"required,description=User to change email for (username, email, or ID)"`
	NewEmail string `json:"newEmail" jsonschema:"required,description=New email address"`
}

// UserAddTeamArgs represents arguments for adding users to a team
type UserAddTeamArgs struct {
	Team  string   `json:"team" jsonschema:"required,description=Team name or ID"`
	Users []string `json:"users" jsonschema:"required,description=Users to add (usernames, emails, or IDs)"`
}

// UserAddChannelArgs represents arguments for adding users to a channel
type UserAddChannelArgs struct {
	Channel string   `json:"channel" jsonschema:"required,description=Channel name or ID (in team:channel format for named channels)"`
	Users   []string `json:"users" jsonschema:"required,description=Users to add (usernames, emails, or IDs)"`
}

// RegisterUserTools registers all user related tools
func RegisterUserTools(server *mcp_golang.Server) error {
	// Register user search tool
	err := server.RegisterTool("user_search", "Search for users", func(args UserSearchArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"user", "search"}
		cmdArgs = append(cmdArgs, args.Terms...)
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "User search completed successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register user_search tool: %v", err)
	}

	// Register user create tool
	err = server.RegisterTool("user_create", "Create a new user", func(args UserCreateArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"user", "create", 
			"--email", args.Email,
			"--username", args.Username,
			"--password", args.Password,
		}
		
		if args.FirstName != "" {
			cmdArgs = append(cmdArgs, "--firstname", args.FirstName)
		}
		
		if args.LastName != "" {
			cmdArgs = append(cmdArgs, "--lastname", args.LastName)
		}
		
		if args.Nickname != "" {
			cmdArgs = append(cmdArgs, "--nickname", args.Nickname)
		}
		
		if args.Locale != "" {
			cmdArgs = append(cmdArgs, "--locale", args.Locale)
		}
		
		if args.SystemAdmin {
			cmdArgs = append(cmdArgs, "--system-admin")
		}
		
		if args.EmailVerified {
			cmdArgs = append(cmdArgs, "--email-verified")
		}
		
		if args.Guest {
			cmdArgs = append(cmdArgs, "--guest")
		}
		
		if args.DisableWelcomeEmail {
			cmdArgs = append(cmdArgs, "--disable-welcome-email")
		}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "User created successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register user_create tool: %v", err)
	}

	// Register user activate tool
	err = server.RegisterTool("user_activate", "Activate users", func(args UserActivateArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"user", "activate"}
		cmdArgs = append(cmdArgs, args.Users...)
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Users activated successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register user_activate tool: %v", err)
	}

	// Register user deactivate tool
	err = server.RegisterTool("user_deactivate", "Deactivate users", func(args UserDeactivateArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"user", "deactivate"}
		cmdArgs = append(cmdArgs, args.Users...)
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Users deactivated successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register user_deactivate tool: %v", err)
	}

	// Register user email tool
	err = server.RegisterTool("user_email", "Change a user's email", func(args UserEmailArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"user", "email", args.User, args.NewEmail}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "User email changed successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register user_email tool: %v", err)
	}

	// Register user add team tool
	err = server.RegisterTool("user_add_team", "Add users to a team", func(args UserAddTeamArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"user", "add", args.Team}
		cmdArgs = append(cmdArgs, args.Users...)
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Users added to team successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register user_add_team tool: %v", err)
	}

	// Register user add channel tool
	err = server.RegisterTool("user_add_channel", "Add users to a channel", func(args UserAddChannelArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"channel", "add", args.Channel}
		cmdArgs = append(cmdArgs, args.Users...)
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		if output == "" {
			output = "Users added to channel successfully"
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register user_add_channel tool: %v", err)
	}

	return nil
}
