package api

import (
	"reflect"
	"testing"

	"github.com/nayakunin/gophkeeper/internal/commands/add/text/input"
	api "github.com/nayakunin/gophkeeper/proto"
)

func TestService_prepareTextRequest(t *testing.T) {
	type fields struct {
		encryption Encryption
	}
	type args struct {
		result        *input.ParseTextResult
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
				encryption: tt.fields.encryption,
			}
			got, err := s.PrepareTextRequest(tt.args.result, tt.args.encryptionKey)
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
