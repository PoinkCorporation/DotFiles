package tests

import (
	permv1 "permissions/gen/go/permissions"
	"permissions/tests/suite"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHasPermission_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)
	userID := randomUserID()
	_, err := st.PermissionsClient.AssignRole(ctx, &permv1.AssignRoleRequest{
		UserId: userID,
		Role:   "admin",
	})
	require.NoError(t, err)
	resp, err := st.PermissionsClient.HasPermission(ctx, &permv1.HasPermissionRequest{
		UserId: userID,
		Role:   "chat:delete",
	})
	require.NoError(t, err)
	assert.True(t, resp.GetExists())
}
func TestUserRoles_HappyPath(t *testing.T) {
	ctx, st := suite.New(t)
	userID := randomUserID()
	_, err := st.PermissionsClient.AssignRole(ctx, &permv1.AssignRoleRequest{
		UserId: userID,
		Role:   "user",
	})
	require.NoError(t, err)
	resp, err := st.PermissionsClient.UserRoles(ctx, &permv1.UserRolesRequest{
		UserId: userID,
	})
	require.NoError(t, err)
	assert.Contains(t, resp.GetRoles(), "user")
}
func TestHasPermission_FailCases(t *testing.T) {
	ctx, st := suite.New(t)
	tests := []struct {
		name        string
		userID      int64
		role        string
		expectedErr string
	}{
		{
			name:        "Empty UserID",
			userID:      0,
			role:        "chat:read",
			expectedErr: "user_id is required",
		},
		{
			name:        "Empty Role",
			userID:      randomUserID(),
			role:        "",
			expectedErr: "role is required",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := st.PermissionsClient.HasPermission(ctx, &permv1.HasPermissionRequest{
				UserId: tt.userID,
				Role:   tt.role,
			})
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}
func TestAssignRole_FailCases(t *testing.T) {
	ctx, st := suite.New(t)
	tests := []struct {
		name        string
		userID      int64
		role        string
		expectedErr string
	}{
		{
			name:        "Empty UserID",
			userID:      0,
			role:        "user",
			expectedErr: "user_id is required",
		},
		{
			name:        "Empty Role",
			userID:      randomUserID(),
			role:        "",
			expectedErr: "role is required",
		},
		{
			name:        "Non-Existent Role",
			userID:      randomUserID(),
			role:        "superadmin",
			expectedErr: "role not found",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := st.PermissionsClient.AssignRole(ctx, &permv1.AssignRoleRequest{
				UserId: tt.userID,
				Role:   tt.role,
			})
			require.Error(t, err)
			require.Contains(t, err.Error(), tt.expectedErr)
		})
	}
}
func TestHasPermission_NoRole(t *testing.T) {
	ctx, st := suite.New(t)
	resp, err := st.PermissionsClient.HasPermission(ctx, &permv1.HasPermissionRequest{
		UserId: randomUserID(),
		Role:   "chat:delete",
	})
	require.NoError(t, err)
	assert.False(t, resp.GetExists())
}
