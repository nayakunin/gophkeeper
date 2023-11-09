package main

import (
	"github.com/nayakunin/gophkeeper/internal/commands"
	"github.com/nayakunin/gophkeeper/internal/commands/transport"
	"github.com/nayakunin/gophkeeper/internal/services/credentials"
	"github.com/nayakunin/gophkeeper/internal/services/encryption"
	"github.com/nayakunin/gophkeeper/internal/services/localstorage"
)

func main() {
	credentialsService, err := credentials.NewService()
	if err != nil {
		panic(err)
	}

	localStorageService := localstorage.NewStorage(credentialsService)
	encryptionService := encryption.NewService()
	apiService := transport.NewService()

	root := commands.NewRoot(localStorageService, encryptionService, apiService)

	root.Execute()
}
