package input

import (
	"fmt"

	generated "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
)

// ParseCardResult is the result of parsing a card request.
func ParseCardRequest(cmd *cobra.Command) (*generated.GetBankCardDetailsRequest, error) {
	label, err := cmd.Flags().GetString("label")
	if err != nil {
		return nil, fmt.Errorf("could not get card label: %w", err)
	}
	if label == "" {
		return nil, fmt.Errorf("please provide a card label")
	}

	return &generated.GetBankCardDetailsRequest{
		CardName: label,
	}, nil
}
