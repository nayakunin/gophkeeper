//go:generate mockgen -source=service.go -destination=mocks/service.go -package=mocks
package text

import (
	"reflect"
	"testing"
)

func TestNewService(t *testing.T) {
	type args struct {
		encryption         Encryption
		credentialsService CredentialsService
		api                Api
	}
	tests := []struct {
		name string
		args args
		want *Service
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewService(tt.args.encryption, tt.args.credentialsService, tt.args.api); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewService() = %v, want %v", got, tt.want)
			}
		})
	}
}
