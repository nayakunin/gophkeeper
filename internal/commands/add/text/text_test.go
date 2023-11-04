package text

import (
	"reflect"
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/add"
	"github.com/spf13/cobra"
)

func TestService_textCmd(t *testing.T) {
	type fields struct {
		credentialsService add.CredentialsService
		encryption         add.Encryption
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
			s := &add.Service{
				credentialsService: tt.fields.credentialsService,
				encryption:         tt.fields.encryption,
			}
			if got := s.textCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("textCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
