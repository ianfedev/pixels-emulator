package role

import (
	"pixels-emulator/core/model"
	"sort"
	"strings"
)

// PermissionsCompound holds a set of permissions with a priority.
type PermissionsCompound struct {
	Priority    int
	Permissions map[string]struct{}
}

// ComparePermission checks if a permission is granted based on role priorities.
func ComparePermission(roles []PermissionsCompound, permission string) bool {
	var allowed bool

	// Sort roles by priority (ascending, lower is stronger)
	sort.Slice(roles, func(i, j int) bool {
		return roles[i].Priority < roles[j].Priority
	})

	for _, role := range roles {
		for perm := range role.Permissions {
			if perm == "*" {
				allowed = true
			} else if perm == "-"+permission || matchWildcard("-"+permission, perm) {
				return false // Explicitly denied
			} else if perm == permission || matchWildcard(permission, perm) {
				allowed = true
			}
		}
	}
	return allowed
}

// matchWildcard checks if a permission matches a wildcard.
func matchWildcard(target, pattern string) bool {
	if strings.HasSuffix(pattern, ".*") {
		return strings.HasPrefix(target, strings.TrimSuffix(pattern, ".*"))
	}
	return target == pattern
}

// HasPermission check if user has permission
func HasPermission(user model.User, permission string) bool {
	var roles []PermissionsCompound

	for _, role := range user.Roles {
		perms := make(map[string]struct{})
		for _, perm := range role.Permissions {
			perms[perm.Permission] = struct{}{}
		}
		roles = append(roles, PermissionsCompound{
			Priority:    role.Priority,
			Permissions: perms,
		})
	}

	return ComparePermission(roles, permission)
}
