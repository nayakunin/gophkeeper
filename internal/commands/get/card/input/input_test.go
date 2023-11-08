package input

import (
	"reflect"
	"testing"

	generated "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
)

func TestParseCardRequest(t *testing.T) {
	type args struct {
		cmd *cobra.Command
	}
	tests := []struct {
		name    string
		args    args
		want    *generated.GetBankCardDetailsRequest
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCardRequest(tt.args.cmd)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseCardRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseCardRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}
