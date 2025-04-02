package main

import (
	"fmt"

	mcp_golang "github.com/metoro-io/mcp-golang"
)

// LicenseRemoveArgs represents arguments for license remove command
type LicenseRemoveArgs struct {
	// No specific args for now
}

// LicenseUploadArgs represents arguments for license upload command
type LicenseUploadArgs struct {
	LicensePath string `json:"licensePath" jsonschema:"required,description=Path to the license file"`
}

// LicenseUploadStringArgs represents arguments for license upload-string command
type LicenseUploadStringArgs struct {
	LicenseString string `json:"licenseString" jsonschema:"required,description=License string to upload"`
}

// RegisterLicenseTools registers all license related tools
func RegisterLicenseTools(server *mcp_golang.Server) error {
	// Register license remove tool
	err := server.RegisterTool("license_remove", "Remove the current license", func(args LicenseRemoveArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"license", "remove"}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register license_remove tool: %v", err)
	}

	// Register license upload tool
	err = server.RegisterTool("license_upload", "Upload a license file", func(args LicenseUploadArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"license", "upload", args.LicensePath}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register license_upload tool: %v", err)
	}

	// Register license upload-string tool
	err = server.RegisterTool("license_upload_string", "Upload a license from a string", func(args LicenseUploadStringArgs) (*mcp_golang.ToolResponse, error) {
		cmdArgs := []string{"license", "upload-string", args.LicenseString}
		
		output, err := executeMMCTL(cmdArgs...)
		if err != nil {
			return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(fmt.Sprintf("Error: %v", err))), nil
		}
		return mcp_golang.NewToolResponse(mcp_golang.NewTextContent(output)), nil
	})
	if err != nil {
		return fmt.Errorf("failed to register license_upload_string tool: %v", err)
	}

	return nil
}
