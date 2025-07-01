package Permission

import (
	"server/internal/Database"
	"strings"
)

func CheckIfUserHasPermission(user *Database.User, permission string) bool {
	// Direct user permissions
	if hasPermission(user.Permissions, permission) {
		return true
	}

	// Role-based permissions
	for _, role := range user.Roles {
		if hasPermission(role.Permissions, permission) {
			return true
		}
	}

	return false
}

func hasPermission(perms []string, target string) bool {
	for _, perm := range perms {
		if perm == target {
			return true
		}
		// Handle wildcard like app.*
		if strings.HasSuffix(perm, ".*") {
			prefix := strings.TrimSuffix(perm, ".*")
			if strings.HasPrefix(target, prefix+".") {
				return true
			}
		}
	}
	return false
}
