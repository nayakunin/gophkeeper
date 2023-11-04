package input

import (
	"reflect"
	"testing"

	"github.com/spf13/cobra"
)

func TestService_parsePasswordRequest(t *testing.T) {
	type args struct {
		cmd *cobra.Command
	}
	tests := []struct {
		name    string
		args    args
		want    *ParsePasswordResult
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePasswordRequest(tt.args.cmd)
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
