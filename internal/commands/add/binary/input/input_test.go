package input

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_ParseBinaryRequest(t *testing.T) {
	type args struct {
		cmd *cobra.Command
	}

	cmd0 := &cobra.Command{}

	cmd1 := &cobra.Command{}
	cmd1.Flags().String("filepath", "", "")

	cmd2 := &cobra.Command{}
	cmd2.Flags().String("description", "", "")

	cmd3 := &cobra.Command{}
	cmd3.Flags().String("filepath", "", "")
	cmd3.Flags().String("description", "", "")

	tests := []struct {
		name    string
		args    args
		want    *ParseBinaryResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return error when no flags are set",
			args: args{
				cmd: cmd0,
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "should return error when only filepath flag is set",
			args: args{
				cmd: cmd1,
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "should return error when only description flag is set",
			args: args{
				cmd: cmd2,
			},
			want:    nil,
			wantErr: assert.Error,
		},
		{
			name: "should return no error when both flags are set",
			args: args{
				cmd: cmd3,
			},
			want: &ParseBinaryResult{
				Filepath:    "",
				Description: "",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseBinaryRequest(tt.args.cmd)
			if !tt.wantErr(t, err, fmt.Sprintf("parseBinaryRequest(%v)", tt.args.cmd)) {
				return
			}
			assert.Equalf(t, tt.want, got, "parseBinaryRequest(%v)", tt.args.cmd)
		})
	}
}
