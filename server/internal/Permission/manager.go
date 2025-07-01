package Permission

import (
	"log"
	"server/internal/Core"
)

var PermissionStore *Core.Store[[]string]

func LoadPermissionStore() {
	store, err := Core.Load[[]string]("stores/permissions.json", []string{})
	if err != nil {
		log.Fatalf("failed to load permission store: %v", err)
	}
	PermissionStore = store
}
