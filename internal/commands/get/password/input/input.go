package input

import (
	"fmt"

	generated "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
)

// ParsePasswordResult is the result of parsing a password request.
func ParsePasswordRequest(cmd *cobra.Command) (*generated.GetLoginPasswordPairsRequest, error) {
	serviceName, err := cmd.Flags().GetString("service")
	if err != nil {
		return nil, fmt.Errorf("could not get service name: %w", err)
	}
	if serviceName == "" {
		return nil, fmt.Errorf("please provide a service name")
	}

	return &generated.GetLoginPasswordPairsRequest{
		ServiceName: serviceName,
	}, nil
}