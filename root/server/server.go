package server

import (
	"account-management/root/config"
)


// Init set router
func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(config.GetString("server.port"))
}
