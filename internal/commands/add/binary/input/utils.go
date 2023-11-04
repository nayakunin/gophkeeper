package input

import (
	"fmt"

	"github.com/spf13/cobra"
)

type ParseBinaryResult struct {
	Filepath    string
	Description string
}

func ParseBinaryRequest(cmd *cobra.Command) (*ParseBinaryResult, error) {
	filepath, err := cmd.Flags().GetString("filepath")
	if err != nil {
		return nil, fmt.Errorf("could not get filepath: %w", err)
	}
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return nil, fmt.Errorf("could not get description: %w", err)
	}

	return &ParseBinaryResult{
		Filepath:    filepath,
		Description: description,
	}, nil
}
