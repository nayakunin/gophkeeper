package data

import (
	"context"
	"errors"
	"testing"

	"github.com/nayakunin/gophkeeper/internal/database"
	"github.com/nayakunin/gophkeeper/internal/grpc/data/mocks"
	"github.com/nayakunin/gophkeeper/pkg/utils/authcommon"
	api "github.com/nayakunin/gophkeeper/proto"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type mockEncryptReply struct {
	encryptedData []byte
	err           error
}

type mockDecryptReply struct {
	decryptedData []byte
	err           error
}

func TestService_AddBinaryData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.WithValue(context.Background(), authcommon.UserIDKey, int64(1))

	type mockAddBinaryDataReply struct {
		err error
	}

	in := &api.AddBinaryDataRequest{
		EncryptedData: []byte("encrypted data"),
		Description:   "description",
	}

	testcases := []struct {
		name                   string
		mockEncryptReply       *mockEncryptReply
		mockAddBinaryDataReply *mockAddBinaryDataReply
		hasError               bool
	}{{
		name: "success",
		mockEncryptReply: &mockEncryptReply{
			encryptedData: []byte("encrypted data"),
		},
		mockAddBinaryDataReply: &mockAddBinaryDataReply{},
	}, {
		name: "encrypt error",
		mockEncryptReply: &mockEncryptReply{
			err: errors.New("encrypt error"),
		},
		hasError: true,
	}, {
		name: "add binary data error",
		mockEncryptReply: &mockEncryptReply{
			encryptedData: []byte("encrypted data"),
		},
		mockAddBinaryDataReply: &mockAddBinaryDataReply{
			err: errors.New("add binary data error"),
		},
		hasError: true,
	}}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockStorage := mocks.NewMockStorage(ctrl)
			if tc.mockAddBinaryDataReply != nil {
				mockStorage.EXPECT().AddBinaryData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(tc.mockAddBinaryDataReply.err)
			}
			mockEncryption := mocks.NewMockEncryption(ctrl)
			if tc.mockEncryptReply != nil {
				mockEncryption.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return(tc.mockEncryptReply.encryptedData, tc.mockEncryptReply.err)
			}
			service := NewService(mockStorage, mockEncryption)

			_, err := service.AddBinaryData(ctx, in)
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_AddBankCardDetail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.WithValue(context.Background(), authcommon.UserIDKey, int64(1))

	in := &api.AddBankCardDetailRequest{
		CardName:            "card name",
		EncryptedCardNumber: []byte("encrypted card number"),
		EncryptedExpiryDate: []byte("encrypted expiry date"),
		EncryptedCvc:        []byte("encrypted cvc"),
		Description:         "description",
	}

	type mockAddBankCardDetailReply struct {
		err error
	}

	testcases := []struct {
		name                       string
		mockEncryptCardNumberReply *mockEncryptReply
		mockEncryptExpiryReply     *mockEncryptReply
		mockEncryptCVCReply        *mockEncryptReply
		mockAddBankCardDetailReply *mockAddBankCardDetailReply
		hasError                   bool
	}{{
		name: "success",
		mockEncryptCardNumberReply: &mockEncryptReply{
			encryptedData: []byte("encrypted card number"),
		},
		mockEncryptExpiryReply: &mockEncryptReply{
			encryptedData: []byte("encrypted expiry date"),
		},
		mockEncryptCVCReply: &mockEncryptReply{
			encryptedData: []byte("encrypted cvc"),
		},
		mockAddBankCardDetailReply: &mockAddBankCardDetailReply{},
	}, {
		name: "encrypt card number error",
		mockEncryptCardNumberReply: &mockEncryptReply{
			err: errors.New("encrypt card number error"),
		},
		hasError: true,
	}, {
		name: "encrypt expiry date error",
		mockEncryptCardNumberReply: &mockEncryptReply{
			encryptedData: []byte("encrypted card number"),
		},
		mockEncryptExpiryReply: &mockEncryptReply{
			err: errors.New("encrypt expiry date error"),
		},
		hasError: true,
	}, {
		name: "encrypt cvc error",
		mockEncryptCardNumberReply: &mockEncryptReply{
			encryptedData: []byte("encrypted card number"),
		},
		mockEncryptExpiryReply: &mockEncryptReply{
			encryptedData: []byte("encrypted expiry date"),
		},
		mockEncryptCVCReply: &mockEncryptReply{
			err: errors.New("encrypt cvc error"),
		},
		hasError: true,
	}, {
		name: "add bank card detail error",
		mockEncryptCardNumberReply: &mockEncryptReply{
			encryptedData: []byte("encrypted card number"),
		},
		mockEncryptExpiryReply: &mockEncryptReply{
			encryptedData: []byte("encrypted expiry date"),
		},
		mockEncryptCVCReply: &mockEncryptReply{
			encryptedData: []byte("encrypted cvc"),
		},
		mockAddBankCardDetailReply: &mockAddBankCardDetailReply{
			err: errors.New("add bank card detail error"),
		},
		hasError: true,
	}}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockStorage := mocks.NewMockStorage(ctrl)
			if tc.mockAddBankCardDetailReply != nil {
				mockStorage.EXPECT().AddBankCardDetails(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(tc.mockAddBankCardDetailReply.err)
			}
			mockEncryption := mocks.NewMockEncryption(ctrl)
			if tc.mockEncryptCardNumberReply != nil {
				mockEncryption.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return(tc.mockEncryptCardNumberReply.encryptedData, tc.mockEncryptCardNumberReply.err)
			}
			if tc.mockEncryptExpiryReply != nil {
				mockEncryption.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return(tc.mockEncryptExpiryReply.encryptedData, tc.mockEncryptExpiryReply.err)
			}
			if tc.mockEncryptCVCReply != nil {
				mockEncryption.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return(tc.mockEncryptCVCReply.encryptedData, tc.mockEncryptCVCReply.err)
			}
			service := NewService(mockStorage, mockEncryption)

			_, err := service.AddBankCardDetail(ctx, in)
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_AddLoginPasswordPair(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.WithValue(context.Background(), authcommon.UserIDKey, int64(1))

	in := &api.AddLoginPasswordPairRequest{
		ServiceName:       "service name",
		Login:             "login",
		EncryptedPassword: []byte("encrypted password"),
		Description:       "description",
	}

	type mockAddLoginPasswordPairReply struct {
		err error
	}

	testcases := []struct {
		name                          string
		mockEncryptPasswordReply      *mockEncryptReply
		mockAddLoginPasswordPairReply *mockAddLoginPasswordPairReply
		hasError                      bool
	}{{
		name: "success",
		mockEncryptPasswordReply: &mockEncryptReply{
			encryptedData: []byte("encrypted password"),
		},
		mockAddLoginPasswordPairReply: &mockAddLoginPasswordPairReply{},
	}, {
		name: "encrypt password error",
		mockEncryptPasswordReply: &mockEncryptReply{
			err: errors.New("encrypt password error"),
		},
		hasError: true,
	}, {
		name: "add login password pair error",
		mockEncryptPasswordReply: &mockEncryptReply{
			encryptedData: []byte("encrypted password"),
		},
		mockAddLoginPasswordPairReply: &mockAddLoginPasswordPairReply{
			err: errors.New("add login password pair error"),
		},
		hasError: true,
	}}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockStorage := mocks.NewMockStorage(ctrl)
			if tc.mockAddLoginPasswordPairReply != nil {
				mockStorage.EXPECT().AddLoginPasswordPair(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(tc.mockAddLoginPasswordPairReply.err)
			}
			mockEncryption := mocks.NewMockEncryption(ctrl)
			if tc.mockEncryptPasswordReply != nil {
				mockEncryption.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return(tc.mockEncryptPasswordReply.encryptedData, tc.mockEncryptPasswordReply.err)
			}
			service := NewService(mockStorage, mockEncryption)

			_, err := service.AddLoginPasswordPair(ctx, in)
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_AddTextData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.WithValue(context.Background(), authcommon.UserIDKey, int64(1))

	in := &api.AddTextDataRequest{
		EncryptedText: []byte("encrypted text"),
		Description:   "description",
	}

	type mockAddTextDataReply struct {
		err error
	}

	testcases := []struct {
		name                 string
		mockEncryptReply     *mockEncryptReply
		mockAddTextDataReply *mockAddTextDataReply
		hasError             bool
	}{{
		name: "success",
		mockEncryptReply: &mockEncryptReply{
			encryptedData: []byte("encrypted text"),
		},
		mockAddTextDataReply: &mockAddTextDataReply{},
	}, {
		name: "encrypt error",
		mockEncryptReply: &mockEncryptReply{
			err: errors.New("encrypt error"),
		},
		hasError: true,
	}, {
		name: "add text data error",
		mockEncryptReply: &mockEncryptReply{
			encryptedData: []byte("encrypted text"),
		},
		mockAddTextDataReply: &mockAddTextDataReply{
			err: errors.New("add text data error"),
		},
		hasError: true,
	}}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockStorage := mocks.NewMockStorage(ctrl)
			if tc.mockAddTextDataReply != nil {
				mockStorage.EXPECT().AddTextData(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(tc.mockAddTextDataReply.err)
			}
			mockEncryption := mocks.NewMockEncryption(ctrl)
			if tc.mockEncryptReply != nil {
				mockEncryption.EXPECT().Encrypt(gomock.Any(), gomock.Any()).Return(tc.mockEncryptReply.encryptedData, tc.mockEncryptReply.err)
			}
			service := NewService(mockStorage, mockEncryption)

			_, err := service.AddTextData(ctx, in)
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestService_GetBinaryData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.WithValue(context.Background(), authcommon.UserIDKey, int64(1))

	type mockGetBinaryDataReply struct {
		binaryData []database.BinaryData
		err        error
	}

	testcases := []struct {
		name                   string
		out                    *api.GetBinaryDataResponse
		mockGetBinaryDataReply *mockGetBinaryDataReply
		mockDecryptReply       *mockDecryptReply
		hasError               bool
	}{{
		name: "success",
		out: &api.GetBinaryDataResponse{
			BinaryData: []*api.GetBinaryDataResponseItem{{
				Id:            1,
				EncryptedData: []byte("decrypted data"),
				Description:   "description",
			}, {
				Id:            2,
				EncryptedData: []byte("decrypted data"),
				Description:   "description",
			}},
		},
		mockGetBinaryDataReply: &mockGetBinaryDataReply{
			binaryData: []database.BinaryData{{
				ID:            1,
				EncryptedData: []byte("encrypted data"),
				Description:   "description",
			}, {
				ID:            2,
				EncryptedData: []byte("encrypted data"),
				Description:   "description",
			}},
		},
		mockDecryptReply: &mockDecryptReply{
			decryptedData: []byte("decrypted data"),
		},
	}, {
		name: "get binary data error",
		mockGetBinaryDataReply: &mockGetBinaryDataReply{
			err: errors.New("get binary data error"),
		},
		hasError: true,
	}, {
		name: "decrypt error",
		mockGetBinaryDataReply: &mockGetBinaryDataReply{
			binaryData: []database.BinaryData{{
				ID:            1,
				EncryptedData: []byte("encrypted data"),
				Description:   "description",
			}, {
				ID:            2,
				EncryptedData: []byte("encrypted data"),
				Description:   "description",
			}},
		},
		mockDecryptReply: &mockDecryptReply{
			err: errors.New("decrypt error"),
		},
		hasError: true,
	}}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockStorage := mocks.NewMockStorage(ctrl)
			if tc.mockGetBinaryDataReply != nil {
				mockStorage.EXPECT().GetBinaryData(gomock.Any(), gomock.Any()).Return(tc.mockGetBinaryDataReply.binaryData, tc.mockGetBinaryDataReply.err)
			}
			mockEncryption := mocks.NewMockEncryption(ctrl)
			if tc.mockDecryptReply != nil {
				mockEncryption.EXPECT().Decrypt(gomock.Any(), gomock.Any()).Return(tc.mockDecryptReply.decryptedData, tc.mockDecryptReply.err).AnyTimes()
			}
			service := NewService(mockStorage, mockEncryption)

			out, err := service.GetBinaryData(ctx, &api.Empty{})
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tc.out != nil {
				assert.Equal(t, tc.out, out)
			}
		})
	}
}

func TestService_GetBankCardDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.WithValue(context.Background(), authcommon.UserIDKey, int64(1))

	type mockGetBankCardDetailsReply struct {
		bankCardDetails []database.BankCardDetail
		err             error
	}

	defaultBankCardDetails := []database.BankCardDetail{{
		ID:                  1,
		CardName:            "card name",
		EncryptedCardNumber: []byte("encrypted card number"),
		EncryptedExpiryDate: []byte("encrypted expiry date"),
		EncryptedCVC:        []byte("encrypted cvc"),
		Description:         "description",
	}}

	testcases := []struct {
		name                        string
		out                         *api.GetBankCardDetailsResponse
		mockGetBankCardDetailsReply *mockGetBankCardDetailsReply
		mockDecryptCardNumberReply  *mockDecryptReply
		mockDecryptExpiryReply      *mockDecryptReply
		mockDecryptCVCReply         *mockDecryptReply
		hasError                    bool
	}{{
		name: "success",
		mockGetBankCardDetailsReply: &mockGetBankCardDetailsReply{
			bankCardDetails: defaultBankCardDetails,
		},
		mockDecryptCardNumberReply: &mockDecryptReply{
			decryptedData: []byte("decrypted card number"),
		},
		mockDecryptExpiryReply: &mockDecryptReply{
			decryptedData: []byte("decrypted expiry date"),
		},
		mockDecryptCVCReply: &mockDecryptReply{
			decryptedData: []byte("decrypted cvc"),
		},
		out: &api.GetBankCardDetailsResponse{
			BankCardDetails: []*api.BankCardDetail{{
				Id:                  1,
				CardName:            "card name",
				EncryptedCardNumber: []byte("decrypted card number"),
				EncryptedExpiryDate: []byte("decrypted expiry date"),
				EncryptedCvc:        []byte("decrypted cvc"),
				Description:         "description",
			}},
		},
	}, {
		name: "get bank card details error",
		mockGetBankCardDetailsReply: &mockGetBankCardDetailsReply{
			err: errors.New("get bank card details error"),
		},
		hasError: true,
	}, {
		name: "decrypt card number error",
		mockGetBankCardDetailsReply: &mockGetBankCardDetailsReply{
			bankCardDetails: defaultBankCardDetails,
		},
		mockDecryptCardNumberReply: &mockDecryptReply{
			err: errors.New("decrypt card number error"),
		},
		hasError: true,
	}, {
		name: "decrypt expiry date error",
		mockGetBankCardDetailsReply: &mockGetBankCardDetailsReply{
			bankCardDetails: defaultBankCardDetails,
		},
		mockDecryptCardNumberReply: &mockDecryptReply{
			decryptedData: []byte("decrypted card number"),
		},
		mockDecryptExpiryReply: &mockDecryptReply{
			err: errors.New("decrypt expiry date error"),
		},
		hasError: true,
	}, {
		name: "decrypt cvc error",
		mockGetBankCardDetailsReply: &mockGetBankCardDetailsReply{
			bankCardDetails: defaultBankCardDetails,
		},
		mockDecryptCardNumberReply: &mockDecryptReply{
			decryptedData: []byte("decrypted card number"),
		},
		mockDecryptExpiryReply: &mockDecryptReply{
			decryptedData: []byte("decrypted expiry date"),
		},
		mockDecryptCVCReply: &mockDecryptReply{
			err: errors.New("decrypt cvc error"),
		},
		hasError: true,
	}}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockStorage := mocks.NewMockStorage(ctrl)
			if tc.mockGetBankCardDetailsReply != nil {
				mockStorage.EXPECT().GetBankCardDetails(gomock.Any(), gomock.Any(), gomock.Any()).Return(tc.mockGetBankCardDetailsReply.bankCardDetails, tc.mockGetBankCardDetailsReply.err)
			}
			mockEncryption := mocks.NewMockEncryption(ctrl)
			if tc.mockDecryptCardNumberReply != nil {
				mockEncryption.EXPECT().Decrypt(gomock.Any(), gomock.Any()).Return(tc.mockDecryptCardNumberReply.decryptedData, tc.mockDecryptCardNumberReply.err)
			}
			if tc.mockDecryptExpiryReply != nil {
				mockEncryption.EXPECT().Decrypt(gomock.Any(), gomock.Any()).Return(tc.mockDecryptExpiryReply.decryptedData, tc.mockDecryptExpiryReply.err)
			}
			if tc.mockDecryptCVCReply != nil {
				mockEncryption.EXPECT().Decrypt(gomock.Any(), gomock.Any()).Return(tc.mockDecryptCVCReply.decryptedData, tc.mockDecryptCVCReply.err)
			}
			service := NewService(mockStorage, mockEncryption)

			out, err := service.GetBankCardDetails(ctx, &api.GetBankCardDetailsRequest{
				CardName: "card name",
			})
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tc.out != nil {
				assert.EqualValues(t, tc.out.GetBankCardDetails(), out.GetBankCardDetails())
			}
		})
	}
}

func TestService_GetLoginPasswordPairs(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.WithValue(context.Background(), authcommon.UserIDKey, int64(1))

	type mockGetLoginPasswordPairsReply struct {
		loginPasswordPairs []database.LoginPasswordPair
		err                error
	}

	defaultLoginPasswordPairs := []database.LoginPasswordPair{{
		ID:                1,
		ServiceName:       "service name",
		Login:             "login",
		EncryptedPassword: []byte("encrypted password"),
	}}

	testcases := []struct {
		name                           string
		out                            *api.GetLoginPasswordPairsResponse
		mockGetLoginPasswordPairsReply *mockGetLoginPasswordPairsReply
		mockDecryptPasswordReply       *mockDecryptReply
		hasError                       bool
	}{{
		name: "success",
		mockGetLoginPasswordPairsReply: &mockGetLoginPasswordPairsReply{
			loginPasswordPairs: defaultLoginPasswordPairs,
		},
		mockDecryptPasswordReply: &mockDecryptReply{
			decryptedData: []byte("decrypted password"),
		},
	}, {
		name: "get login password pairs error",
		mockGetLoginPasswordPairsReply: &mockGetLoginPasswordPairsReply{
			err: errors.New("get login password pairs error"),
		},
		hasError: true,
	}, {
		name: "decrypt password error",
		mockGetLoginPasswordPairsReply: &mockGetLoginPasswordPairsReply{
			loginPasswordPairs: defaultLoginPasswordPairs,
		},
		mockDecryptPasswordReply: &mockDecryptReply{
			err: errors.New("decrypt password error"),
		},
		hasError: true,
	}}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockStorage := mocks.NewMockStorage(ctrl)
			if tc.mockGetLoginPasswordPairsReply != nil {
				mockStorage.EXPECT().GetLoginPasswordPairs(gomock.Any(), gomock.Any(), gomock.Any()).Return(tc.mockGetLoginPasswordPairsReply.loginPasswordPairs, tc.mockGetLoginPasswordPairsReply.err)
			}
			mockEncryption := mocks.NewMockEncryption(ctrl)
			if tc.mockDecryptPasswordReply != nil {
				mockEncryption.EXPECT().Decrypt(gomock.Any(), gomock.Any()).Return(tc.mockDecryptPasswordReply.decryptedData, tc.mockDecryptPasswordReply.err)
			}
			service := NewService(mockStorage, mockEncryption)

			out, err := service.GetLoginPasswordPairs(ctx, &api.GetLoginPasswordPairsRequest{
				ServiceName: "service name",
			})
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tc.out != nil {
				assert.EqualValues(t, tc.out.GetLoginPasswordPairs(), out.GetLoginPasswordPairs())
			}
		})
	}
}

func TestService_GetTextData(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ctx := context.WithValue(context.Background(), authcommon.UserIDKey, int64(1))

	type mockGetTextDataReply struct {
		textData []database.TextData
		err      error
	}

	defaultTextData := []database.TextData{{
		ID:            1,
		EncryptedText: []byte("encrypted text"),
		Description:   "description",
	}}

	testcases := []struct {
		name                 string
		out                  *api.GetTextDataResponse
		mockGetTextDataReply *mockGetTextDataReply
		mockDecryptReply     *mockDecryptReply
		hasError             bool
	}{{
		name: "success",
		out: &api.GetTextDataResponse{
			TextData: []*api.GetTextDataResponseItem{{
				Id:            1,
				EncryptedText: []byte("decrypted text"),
				Description:   "description",
			}},
		},
		mockGetTextDataReply: &mockGetTextDataReply{
			textData: defaultTextData,
		},
		mockDecryptReply: &mockDecryptReply{
			decryptedData: []byte("decrypted text"),
		},
	}, {
		name: "get text data error",
		mockGetTextDataReply: &mockGetTextDataReply{
			err: errors.New("get text data error"),
		},
		hasError: true,
	}, {
		name: "decrypt error",
		mockGetTextDataReply: &mockGetTextDataReply{
			textData: defaultTextData,
		},
		mockDecryptReply: &mockDecryptReply{
			err: errors.New("decrypt error"),
		},
		hasError: true,
	}}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockStorage := mocks.NewMockStorage(ctrl)
			if tc.mockGetTextDataReply != nil {
				mockStorage.EXPECT().GetTextData(gomock.Any(), gomock.Any()).Return(tc.mockGetTextDataReply.textData, tc.mockGetTextDataReply.err)
			}
			mockEncryption := mocks.NewMockEncryption(ctrl)
			if tc.mockDecryptReply != nil {
				mockEncryption.EXPECT().Decrypt(gomock.Any(), gomock.Any()).Return(tc.mockDecryptReply.decryptedData, tc.mockDecryptReply.err)
			}
			service := NewService(mockStorage, mockEncryption)

			out, err := service.GetTextData(ctx, &api.Empty{})
			if tc.hasError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if tc.out != nil {
				assert.Equal(t, tc.out, out)
			}
		})
	}
}
