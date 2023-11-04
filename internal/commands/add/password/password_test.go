package password

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

func TestService_passwordCmd(t *testing.T) {
	type fields struct {
		credentialsService CredentialsService
		encryption         Encryption
		apiPreparer        Api
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
				apiPreparer:        tt.fields.apiPreparer,
			}
			if got := s.GetCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("passwordCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
