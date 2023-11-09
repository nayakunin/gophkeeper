package input

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ParseTextResult is the result of parsing a text request.
type ParseTextResult struct {
	Text        string
	Description string
}

// ParseTextRequest parses a text request.
func ParseTextRequest(cmd *cobra.Command) (*ParseTextResult, error) {
	text, err := cmd.Flags().GetString("text")
	if err != nil {
		return nil, fmt.Errorf("could not get text: %w", err)
	}
	if text == "" {
		return nil, fmt.Errorf("please provide a text")
	}
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return nil, fmt.Errorf("could not get description: %w", err)
	}

	return &ParseTextResult{
		Text:        text,
		Description: description,
	}, nil
}
