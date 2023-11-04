package input

import (
	"fmt"

	"github.com/spf13/cobra"
)

type ParseCardResult struct {
	Name        string
	Number      string
	Expiration  string
	Cvc         string
	Description string
}

func ParseCardRequest(cmd *cobra.Command) (*ParseCardResult, error) {
	name, err := cmd.Flags().GetString("label")
	if err != nil {
		return nil, fmt.Errorf("could not get card name: %w", err)
	}
	if name == "" {
		return nil, fmt.Errorf("please provide a card name")
	}
	number, err := cmd.Flags().GetString("number")
	if err != nil {
		return nil, fmt.Errorf("could not get card number: %w", err)
	}
	if number == "" {
		return nil, fmt.Errorf("please provide a card number")
	}
	expiration, err := cmd.Flags().GetString("expiration")
	if err != nil {
		return nil, fmt.Errorf("could not get card expiration date: %w", err)
	}
	if expiration == "" {
		return nil, fmt.Errorf("please provide a card expiration date")
	}
	cvc, err := cmd.Flags().GetString("cvc")
	if err != nil {
		return nil, fmt.Errorf("could not get card CVC: %w", err)
	}
	if cvc == "" {
		return nil, fmt.Errorf("please provide a card CVC")
	}
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return nil, fmt.Errorf("could not get card description: %w", err)
	}

	return &ParseCardResult{
		Name:        name,
		Number:      number,
		Expiration:  expiration,
		Cvc:         cvc,
		Description: description,
	}, nil
}
