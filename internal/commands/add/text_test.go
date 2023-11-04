package add

import (
	"reflect"
	"testing"

	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
)

func TestService_prepareTextRequest(t *testing.T) {
	type fields struct {
		credentialsService CredentialsService
		encryption         Encryption
	}
	type args struct {
		result        *parseTextResult
		encryptionKey []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.AddTextDataRequest
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				credentialsService: tt.fields.credentialsService,
				encryption:         tt.fields.encryption,
			}
			got, err := s.prepareTextRequest(tt.args.result, tt.args.encryptionKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("prepareTextRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prepareTextRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_textCmd(t *testing.T) {
	type fields struct {
		credentialsService CredentialsService
		encryption         Encryption
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
			}
			if got := s.textCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("textCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}
