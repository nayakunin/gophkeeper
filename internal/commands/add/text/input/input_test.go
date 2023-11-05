package input

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestParseTextRequest(t *testing.T) {
	type args struct {
		text        string
		description string
	}
	tests := []struct {
		name    string
		args    args
		want    *ParseTextResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return error if text is empty",
			args: args{
				text:        "",
				description: "test",
			},
			wantErr: assert.Error,
		},
		{
			name: "should return no error",
			args: args{
				text:        "test",
				description: "test",
			},
			want: &ParseTextResult{
				Text:        "test",
				Description: "test",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().String("text", tt.args.text, "")
			cmd.Flags().String("description", tt.args.description, "")

			got, err := ParseTextRequest(cmd)

			if !tt.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
