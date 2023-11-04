package api

import (
	"reflect"
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/add/card/input"
	api "github.com/nayakunin/gophkeeper/proto"
)

func TestService_prepareCardRequest(t *testing.T) {
	type fields struct {
		encryption Encryption
	}
	type args struct {
		data          *input.ParseCardResult
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
				encryption: tt.fields.encryption,
			}
			got, err := s.PrepareCardRequest(tt.args.data, tt.args.encryptionKey)
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
