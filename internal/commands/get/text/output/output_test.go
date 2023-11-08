//go:generate mockgen -source=output.go -destination=mocks/service.go -package=mocks
package output

import (
	"reflect"
	"testing"

	generated "github.com/nayakunin/gophkeeper/proto"
)

func TestService_MakeResponse(t *testing.T) {
	type fields struct {
		encryption Encryption
	}
	type args struct {
		response      *generated.GetTextDataResponse
		encryptionKey []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []TextResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				encryption: tt.fields.encryption,
			}
			got, err := s.MakeResponse(tt.args.response, tt.args.encryptionKey)
			if (err != nil) != tt.wantErr {
				t.Errorf("Service.MakeResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Service.MakeResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}
