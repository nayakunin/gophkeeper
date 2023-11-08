package binary

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

func TestService_GetCmd(t *testing.T) {
	type fields struct {
		encryption         Encryption
		credentialsService CredentialsService
		output             Output
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
				encryption:         tt.fields.encryption,
				credentialsService: tt.fields.credentialsService,
				output:             tt.fields.output,
				api:                tt.fields.api,
			}
			if got := s.GetCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
