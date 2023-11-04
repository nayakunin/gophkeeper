package add

import (
	"reflect"
	"testing"

	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
)

func TestService_passwordCmd(t *testing.T) {
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
			if got := s.passwordCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("passwordCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_parsePasswordRequest(t *testing.T) {
	type fields struct {
		credentialsService CredentialsService
		encryption         Encryption
	}
	type args struct {
		cmd *cobra.Command
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *parsePasswordResult
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
			got, err := s.parsePasswordRequest(tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("parsePasswordRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePasswordRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_preparePasswordRequest(t *testing.T) {
	type fields struct {
		credentialsService CredentialsService
		encryption         Encryption
	}
	type args struct {
		result        *parsePasswordResult
		encryptionKey []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.AddLoginPasswordPairRequest
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
			got, err := s.preparePasswordRequest(tt.args.result, tt.args.encryptionKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("preparePasswordRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("preparePasswordRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
