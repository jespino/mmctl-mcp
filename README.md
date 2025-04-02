# mmctl-mcp - Mattermost CLI for Model Context Protocol

[MCP](https://github.com/metoro-io/mcp-golang) server that provides access to [Mattermost](https://mattermost.com/) administration capabilities via mmctl. This plugin allows AI models like Claude to manage, configure, and interact with Mattermost servers through a structured interface.

## Features

- Complete Mattermost administrative capabilities via Claude
- 50+ tools covering all aspects of Mattermost management
- Structured JSON schema for effective AI model interaction
- Local-mode operation through mmctl
- Direct Claude API integration for complex analysis tasks

## Setup Guide

### Prerequisites

- Go 1.20 or higher
- Mattermost server with mmctl installed and configured
- Anthropic API key (optional, for Claude API integration)

### Installation

```bash
# Using Go
go install github.com/jespino/mmctl-mcp@latest

# Using Git
git clone https://github.com/jespino/mmctl-mcp.git
cd mmctl-mcp
go build -o mmctl-mcp .
```

### Configuration and Usage

Before using mmctl-mcp, ensure mmctl is properly configured with your Mattermost server:

```bash
mmctl auth login https://your-mattermost-server.com --name your-server
mmctl auth current
```

The plugin uses mmctl in local mode, so no additional configuration is needed.

For Claude API integration, set your Anthropic API key in the environment:

```bash
export ANTHROPIC_API_KEY=your-api-key
```

## Claude Integration

### Using Claude Desktop App

Add mmctl-mcp to Claude Desktop with this configuration:

```json
{
  "mcpServers": {
    "mattermost-admin": {
      "command": "mmctl-mcp"
    }
  }
}
```

### Using Claude Web Interface

1. Start mmctl-mcp in your terminal
2. In Claude web interface, use /mcp command
3. Follow prompts to connect to your local MCP server

## Available Tools

The plugin provides tools in the following categories:

| Category | Description | Example Tools |
|----------|-------------|--------------|
| System | General system operations | system_info |
| Authentication | Manage authentication | auth_list, auth_current, auth_set |
| Teams | Team management | team_list, team_create, team_search |
| Channels | Channel operations | channel_list, channel_create, channel_archive |
| Users | User management | user_list, user_search, user_create |
| Posts | Message management | post_create, post_list, post_delete |
| Plugins | Plugin management | plugin_list, plugin_enable, plugin_disable |
| Configuration | Server configuration | config_get, config_set, config_show |
| Permissions | Role permissions | permission_add, permission_remove |
| Roles | User roles | role_system_admin, role_member |
| Webhooks | Webhook management | webhook_list, webhook_create_incoming |
| Bots | Bot management | bot_list, bot_create, bot_enable |
| Groups | Group management | group_channel_list, group_team_list |
| Jobs | Server jobs | job_list, job_update |
| LDAP | LDAP integration | ldap_sync, ldap_idmigrate |
| SAML | SAML configuration | saml_auth_data_reset |
| OAuth | OAuth applications | oauth_list |
| License | License management | license_remove, license_upload, license_upload_string |
| Claude | Claude API integration | claude_prompt, claude_file_analysis |

## Examples

### Managing Teams and Channels

```
User: "Create a new team called 'engineering' with display name 'Engineering Team'"
Claude: Uses team_create tool to make the team

User: "Create a private channel in the engineering team called 'security-updates'"
Claude: Uses channel_create tool with the private flag
```

### User Management

```
User: "List all users on the system"
Claude: Uses user_list tool to show all users

User: "Create a new user account for john@example.com"
Claude: Uses user_create tool to add the user
```

### Plugin Administration

```
User: "What plugins are currently enabled?"
Claude: Uses plugin_list tool to show enabled plugins

User: "Disable the jira plugin"
Claude: Uses plugin_disable tool to disable the specified plugin
```

### License Management

```
User: "Upload a new license file for our Mattermost server"
Claude: Uses license_upload tool to upload the license file

User: "Remove the current license and revert to Team Edition"
Claude: Uses license_remove tool to remove the current license
```

## Development & Debugging

### Local Development Setup

Clone the repository and run locally:

```bash
git clone https://github.com/jespino/mmctl-mcp.git
cd mmctl-mcp
go run main.go
```

### Testing with MCP Inspector

Use the MCP Inspector to test your mmctl-mcp server:

```bash
npx @modelcontextprotocol/inspector go run main.go
```

### Building for Distribution

```bash
go build -o mmctl-mcp .
```

## Security Considerations

- The plugin requires properly configured mmctl credentials
- All operations use mmctl's authentication and authorization
- Consider using a dedicated Mattermost account with appropriate permissions
- When using Claude API integration, secure your API key properly

## License

MIT License - see [LICENSE](LICENSE) file.

## Topics

`mcp` `mattermost` `mmctl` `claude` `ai-tools` `administration` `server-management`
