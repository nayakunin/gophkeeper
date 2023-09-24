package main

import (
	"github.com/nayakunin/gophkeeper/internal/commands"
	"github.com/nayakunin/gophkeeper/internal/credentials"
	"github.com/nayakunin/gophkeeper/internal/localstorage"
)

func main() {
	credentialsService := credentials.NewService()
	localStorageService := localstorage.NewStorage(credentialsService)

	root := commands.NewRoot(localStorageService)

	root.Execute()
}
