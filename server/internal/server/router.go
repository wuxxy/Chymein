package server

import (
	"server/internal/Auth"
	"server/internal/Core"
	"server/internal/User"
)

func Router(server *Core.Server) {
	server.RegisterRoute("GET", "/status", Status)

	if !server.Config.IsSetup {
		server.RegisterRoute("POST", "/create_admin", User.CreateSuperAdmin)
	}

	server.RegisterRoute("POST", "/Auth/login", Auth.Login)
	server.RegisterRoute("GET", "/Auth/me", Auth.Me)
}
