package card

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

func TestService_GetCmd(t *testing.T) {
	type fields struct {
		output             Output
		credentialsService CredentialsService
		encryption         Encryption
		api                Api
	}
	tests := []struct {
		name   string
		fields fields
		want   *cobra.Command
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				output:             tt.fields.output,
				credentialsService: tt.fields.credentialsService,
				encryption:         tt.fields.encryption,
				api:                tt.fields.api,
			}
			if got := s.GetCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
