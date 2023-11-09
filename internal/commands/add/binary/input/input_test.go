package input

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_ParseBinaryRequest(t *testing.T) {
	type args struct {
		filepath    string
		description string
	}

	tests := []struct {
		name    string
		args    *args
		want    *ParseBinaryResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name:    "should return error when no flags are set",
			wantErr: assert.Error,
		},
		{
			name: "should return error when filepath is an empty string",
			args: &args{
				filepath: "",
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error when no filepath is set",
			args: &args{
				description: "",
			},
			wantErr: assert.Error,
		},
		{
			name: "should return no error when both flags are set",
			args: &args{
				filepath:    "filepath",
				description: "",
			},
			want: &ParseBinaryResult{
				Filepath:    "filepath",
				Description: "",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			if tt.args != nil {
				cmd.Flags().String("filepath", tt.args.filepath, "")
				cmd.Flags().String("description", tt.args.description, "")
			}

			got, err := ParseBinaryRequest(cmd)
			if !tt.wantErr(t, err) {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
