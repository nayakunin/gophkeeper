package add

import (
	"reflect"
	"testing"

	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
)

func TestService_cardCmd(t *testing.T) {
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
			if got := s.cardCmd(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("cardCmd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_parseCardRequest(t *testing.T) {
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
		want    *parseCardResult
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
			got, err := s.parseCardRequest(tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCardRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseCardRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_prepareCardRequest(t *testing.T) {
	type fields struct {
		credentialsService CredentialsService
		encryption         Encryption
	}
	type args struct {
		data          *parseCardResult
		encryptionKey []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *api.AddBankCardDetailRequest
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
			got, err := s.prepareCardRequest(tt.args.data, tt.args.encryptionKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("prepareCardRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prepareCardRequest() got = %v, want %v", got, tt.want)
			}
		})
	}
}
