package input

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_ParseCardRequest(t *testing.T) {
	type args struct {
		label       string
		number      string
		expiration  string
		cvc         string
		description string
	}
	tests := []struct {
		name    string
		args    args
		want    *ParseCardResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return error when label is empty",
			args: args{
				label: "",
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error when number is empty",
			args: args{
				label:  "label",
				number: "",
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error when expiration is empty",
			args: args{
				label:      "label",
				number:     "number",
				expiration: "",
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error when cvc is empty",
			args: args{
				label:      "label",
				number:     "number",
				expiration: "expiration",
				cvc:        "",
			},
			wantErr: assert.Error,
		},
		{
			name: "should return no error",
			args: args{
				label:       "label",
				number:      "number",
				expiration:  "expiration",
				cvc:         "cvc",
				description: "description",
			},
			want: &ParseCardResult{
				Name:        "label",
				Number:      "number",
				Expiration:  "expiration",
				Cvc:         "cvc",
				Description: "description",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}

			cmd.Flags().String("label", tt.args.label, "")
			cmd.Flags().String("number", tt.args.number, "")
			cmd.Flags().String("expiration", tt.args.expiration, "")
			cmd.Flags().String("cvc", tt.args.cvc, "")
			cmd.Flags().String("description", tt.args.description, "")

			got, err := ParseCardRequest(cmd)
			if !tt.wantErr(t, err) {
				t.Fail()
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
