package add

import (
	"fmt"

	"github.com/nayakunin/gophkeeper/constants"
	"github.com/nayakunin/gophkeeper/pkg/utils"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type parsePasswordResult struct {
	ServiceName string
	Login       string
	Password    string
	Description string
}

func (s *Service) parsePasswordRequest(cmd *cobra.Command) (*parsePasswordResult, error) {
	serviceName, err := cmd.Flags().GetString("service")
	if err != nil {
		return nil, fmt.Errorf("could not get service name: %w", err)
	}
	if serviceName == "" {
		return nil, fmt.Errorf("please provide a service name")
	}
	login, err := cmd.Flags().GetString("login")
	if err != nil {
		return nil, fmt.Errorf("could not get login: %w", err)
	}
	if login == "" {
		return nil, fmt.Errorf("please provide a login")
	}
	password, err := cmd.Flags().GetString("password")
	if err != nil {
		return nil, fmt.Errorf("could not get password: %w", err)
	}
	if password == "" {
		return nil, fmt.Errorf("please provide a password")
	}
	description, err := cmd.Flags().GetString("description")
	if err != nil {
		return nil, fmt.Errorf("could not get description: %w", err)
	}

	return &parsePasswordResult{
		ServiceName: serviceName,
		Login:       login,
		Password:    password,
		Description: description,
	}, nil
}

func (s *Service) preparePasswordRequest(result *parsePasswordResult, encryptionKey []byte) (*api.AddLoginPasswordPairRequest, error) {
	encryptedPassword, err := s.encryption.Encrypt([]byte(result.Password), encryptionKey)
	if err != nil {
		return nil, fmt.Errorf("could not encrypt password: %w", err)
	}

	return &api.AddLoginPasswordPairRequest{
		ServiceName:       result.ServiceName,
		Login:             result.Login,
		EncryptedPassword: encryptedPassword,
		Description:       result.Description,
	}, nil
}

func (s *Service) passwordCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "password",
		Short: "Add a new password",
		RunE: func(cmd *cobra.Command, args []string) error {
			token, encryptionKey, err := s.credentialsService.GetCredentials()
			if err != nil {
				return fmt.Errorf("unable to get credentials: %w", err)
			}

			tmpResult, err := s.parsePasswordRequest(cmd)
			if err != nil {
				return fmt.Errorf("could not parse request: %w", err)
			}

			request, err := s.preparePasswordRequest(tmpResult, encryptionKey)
			if err != nil {
				return fmt.Errorf("could not prepare request: %w", err)
			}

			conn, err := grpc.Dial(constants.GrpcURL, grpc.WithInsecure())
			if err != nil {
				return fmt.Errorf("could not connect: %w", err)
			}
			defer conn.Close()

			client := api.NewDataServiceClient(conn)
			md := utils.GetRequestMetadata(token)
			ctx := metadata.NewOutgoingContext(cmd.Context(), md)
			_, err = client.AddLoginPasswordPair(ctx, request)
			if err != nil {
				return fmt.Errorf("could not add password: %w", err)
			}

			fmt.Println("Password added")
			return nil
		},
	}

	cmd.Flags().StringP("service", "s", "", "Service name")
	_ = cmd.MarkFlagRequired("service")
	cmd.Flags().StringP("login", "l", "", "Login")
	_ = cmd.MarkFlagRequired("login")
	cmd.Flags().StringP("password", "p", "", "Password")
	_ = cmd.MarkFlagRequired("password")
	cmd.Flags().StringP("description", "d", "", "Description")

	return cmd
}
