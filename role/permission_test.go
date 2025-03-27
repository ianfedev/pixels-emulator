package role_test

import (
	"pixels-emulator/core/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"pixels-emulator/role"
)

func TestComparePermission(t *testing.T) {
	tests := []struct {
		name       string
		roles      []role.PermissionsCompound
		permission string
		expected   bool
	}{
		{
			name: "Allow exact match",
			roles: []role.PermissionsCompound{{
				Priority:    100,
				Permissions: map[string]struct{}{"pixels.hello": {}},
			}},
			permission: "pixels.hello",
			expected:   true,
		},
		{
			name: "Deny explicitly negated",
			roles: []role.PermissionsCompound{{
				Priority:    50,
				Permissions: map[string]struct{}{"-pixels.hello": {}},
			}},
			permission: "pixels.hello",
			expected:   false,
		},
		{
			name: "Allow wildcard match",
			roles: []role.PermissionsCompound{{
				Priority:    100,
				Permissions: map[string]struct{}{"pixels.hello.*": {}},
			}},
			permission: "pixels.hello.world",
			expected:   true,
		},
		{
			name: "Deny when wildcard negates",
			roles: []role.PermissionsCompound{{
				Priority:    50,
				Permissions: map[string]struct{}{"-pixels.hello.*": {}},
			}},
			permission: "pixels.hello.world",
			expected:   false,
		},
		{
			name: "Allow global wildcard",
			roles: []role.PermissionsCompound{{
				Priority:    100,
				Permissions: map[string]struct{}{"*": {}},
			}},
			permission: "any.permission",
			expected:   true,
		},
		{
			name: "Deny when global wildcard overridden",
			roles: []role.PermissionsCompound{
				{
					Priority:    50,
					Permissions: map[string]struct{}{"-any.permission": {}},
				},
				{
					Priority:    100,
					Permissions: map[string]struct{}{"*": {}},
				},
			},
			permission: "any.permission",
			expected:   false,
		},
		{
			name: "Deny takes precedence over allow in lower priority",
			roles: []role.PermissionsCompound{
				{
					Priority:    50,
					Permissions: map[string]struct{}{"-common.punishments.ban": {}},
				},
				{
					Priority:    100,
					Permissions: map[string]struct{}{"common.punishments.ban": {}},
				},
			},
			permission: "common.punishments.ban",
			expected:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := role.ComparePermission(tt.roles, tt.permission)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestHasPermission(t *testing.T) {
	tests := []struct {
		name       string
		user       model.User
		permission string
		expected   bool
	}{
		{
			name: "User has direct permission",
			user: model.User{
				Roles: []model.Role{
					{
						Priority: 100,
						Permissions: []model.RolePermission{
							{Permission: "pixels.hello"},
						},
					},
				},
			},
			permission: "pixels.hello",
			expected:   true,
		},
		{
			name: "User does not have permission",
			user: model.User{
				Roles: []model.Role{},
			},
			permission: "pixels.hello",
			expected:   false,
		},
		{
			name: "User has wildcard permission",
			user: model.User{
				Roles: []model.Role{
					{
						Priority: 100,
						Permissions: []model.RolePermission{
							{Permission: "pixels.hello.*"},
						},
					},
				},
			},
			permission: "pixels.hello.world",
			expected:   true,
		},
		{
			name: "User has global wildcard",
			user: model.User{
				Roles: []model.Role{
					{
						Priority: 100,
						Permissions: []model.RolePermission{
							{Permission: "*"},
						},
					},
				},
			},
			permission: "any.permission",
			expected:   true,
		},
		{
			name: "Explicitly denied permission",
			user: model.User{
				Roles: []model.Role{
					{
						Priority: 50,
						Permissions: []model.RolePermission{
							{Permission: "-pixels.hello"},
						},
					},
					{
						Priority: 100,
						Permissions: []model.RolePermission{
							{Permission: "pixels.hello"},
						},
					},
				},
			},
			permission: "pixels.hello",
			expected:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := role.HasPermission(tt.user, tt.permission)
			assert.Equal(t, tt.expected, result)
		})
	}
}
