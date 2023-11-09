package input

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestService_parsePasswordRequest(t *testing.T) {
	type args struct {
		service     string
		login       string
		password    string
		description string
	}
	tests := []struct {
		name    string
		args    args
		want    *ParsePasswordResult
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "should return error if service is empty",
			args: args{
				service:     "",
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error if login is empty",
			args: args{
				service:     "service",
				login:       "",
			},
			wantErr: assert.Error,
		},
		{
			name: "should return error if password is empty",
			args: args{
				service:     "service",
				login:       "login",
				password:    "",
			},
			wantErr: assert.Error,
		},
		{
			name: "should return no error",
			args: args{
				service:     "service",
				login:       "login",
				password:    "password",
				description: "description",
			},
			want: &ParsePasswordResult{
				ServiceName: "service",
				Login:       "login",
				Password:    "password",
				Description: "description",
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := &cobra.Command{}
			cmd.Flags().String("service", tt.args.service, "")
			cmd.Flags().String("login", tt.args.login, "")
			cmd.Flags().String("password", tt.args.password, "")
			cmd.Flags().String("description", tt.args.description, "")

			got, err := ParsePasswordRequest(cmd)
			if !tt.wantErr(t, err) {
				return
			}

			assert.Equal(t, tt.want, got)
		})
	}
}
