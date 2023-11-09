package input

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ParsePasswordResult is the result of parsing a password request.
type ParsePasswordResult struct {
	ServiceName string
	Login       string
	Password    string
	Description string
}

// ParsePasswordRequest parses a password request.
func ParsePasswordRequest(cmd *cobra.Command) (*ParsePasswordResult, error) {
	serviceName, err := cmd.Flags().GetString("service")
	if err != nil {
		return nil, fmt.Errorf("could not get service name: %w", err)
	}
	if serviceName == "" {
		return nil, fmt.Errorf("please provide a service name")
	}
	login, err := cmd.Flags().GetString("login")
	if err != nil {
		return nil, fmt.Errorf("could not get login: %w", err)
	}
	if login == "" {
		return nil, fmt.Errorf("please provide a login")
	}
	password, err := cmd.Flags().GetString("password")
	if err != nil {
		return nil, fmt.Errorf("could not get password: %w", err)
	}
	if password == "" {
		return nil, fmt.Errorf("please provide a password")
	}
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return nil, fmt.Errorf("could not get description: %w", err)
	}

	return &ParsePasswordResult{
		ServiceName: serviceName,
		Login:       login,
		Password:    password,
		Description: description,
	}, nil
}
