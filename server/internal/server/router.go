package server

import (
	"server/internal/Admin"
	"server/internal/Auth"
	"server/internal/Content"
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
	server.RegisterRoute("GET", "/Admin/users", Admin.AllUsers)
	server.RegisterRoute("POST", "/Admin/containers", Admin.CreateContainer)
	server.RegisterRoute("GET", "/Content/content", Content.GetContent)
	server.RegisterRoute("GET", "/Admin/containers", Admin.AllContainers)
}
