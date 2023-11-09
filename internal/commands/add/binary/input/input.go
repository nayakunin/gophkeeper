package input

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ParseBinaryResult is the result of parsing a binary request.
type ParseBinaryResult struct {
	Filepath    string
	Description string
}

// ParseBinaryRequest parses a binary request.
func ParseBinaryRequest(cmd *cobra.Command) (*ParseBinaryResult, error) {
	filepath, err := cmd.Flags().GetString("filepath")
	if err != nil {
		return nil, fmt.Errorf("could not get filepath: %w", err)
	}
	if filepath == "" {
		return nil, fmt.Errorf("please provide a filepath")
	}
	description, _ := cmd.Flags().GetString("description")

	return &ParseBinaryResult{
		Filepath:    filepath,
		Description: description,
	}, nil
}
