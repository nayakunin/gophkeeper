package text

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

func TestService_GetCmd(t *testing.T) {
	type fields struct {
		credentialsService CredentialsService
		encryption         Encryption
		api                Api
		output             Output
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
				credentialsService: tt.fields.credentialsService,
				encryption:         tt.fields.encryption,
				api:                tt.fields.api,
				output:             tt.fields.output,
			}
			if got := s.GetCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.GetCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
