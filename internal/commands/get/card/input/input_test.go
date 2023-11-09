package input

import (
	"testing"

	generated "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestParseCardRequest(t *testing.T) {
	type args struct {
		label string
	}
	tests := []struct {
		name    string
		args    args
		want    *generated.GetBankCardDetailsRequest
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return error if card label is empty",
			args: args{
				label: "",
			},
			wantErr: assert.Error,
		},
		{
			name: "should return card request",
			args: args{
				label: "test",
			},
			want: &generated.GetBankCardDetailsRequest{
				CardName: "test",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}

			cmd.Flags().String("label", tt.args.label, "")

			got, err := ParseCardRequest(cmd)

			if !tt.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
