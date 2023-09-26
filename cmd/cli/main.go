package main

import (
	"github.com/nayakunin/gophkeeper/internal/commands"
	"github.com/nayakunin/gophkeeper/internal/services/credentials"
	"github.com/nayakunin/gophkeeper/internal/services/encryption"
	"github.com/nayakunin/gophkeeper/internal/services/localstorage"
)

func main() {
	credentialsService := credentials.NewService()
	localStorageService := localstorage.NewStorage(credentialsService)
	encryptionService := encryption.NewService()

	root := commands.NewRoot(localStorageService, encryptionService)

	root.Execute()
}
