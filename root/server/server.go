package server

import (
	"account-management/root/config"
	"account-management/root/database"
	"account-management/root/domain/repository"
)


// Init set router
func Init() {
	config := config.GetConfig()
	r := NewRouter()

	dataStoreInterface := database.NewPostgresql()
	repository.NewUserRepository(dataStoreInterface)

	r.Run(config.GetString("server.port"))
}
