package input

import (
	"testing"

	generated "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestParsePasswordRequest(t *testing.T) {
	type args struct {
		service string
	}
	tests := []struct {
		name    string
		args    args
		want    *generated.GetLoginPasswordPairsRequest
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return error if service name is empty",
			args: args{
				service: "",
			},
			wantErr: assert.Error,
		},
		{
			name: "should return request",
			args: args{
				service: "test",
			},
			want: &generated.GetLoginPasswordPairsRequest{
				ServiceName: "test",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().String("service", tt.args.service, "")

			got, err := ParsePasswordRequest(cmd)

			if !tt.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
