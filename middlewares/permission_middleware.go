package middlewares

import (
	"ametory-crud/config"
	"ametory-crud/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PermissionMiddleware checks if the user has the required permissions
func PermissionMiddleware(requiredPermissions []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authData, exists := c.Get("auth")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Authentication data not found"})
			c.Abort()
			return
		}
		if !config.App.Server.UseACL {
			c.Next()
			return
		}

		auth := authData.(models.Auth)

		userPermissions, err := auth.GetPermissions() // assuming Auth model has Permissions field
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user permissions"})
			c.Abort()
			return
		}
		if !hasRequiredPermissions(userPermissions, requiredPermissions) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// SuperAdminMiddleware checks if the user is a super admin
func SuperAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authData, exists := c.Get("auth")
		if !exists {
			c.JSON(http.StatusForbidden, gin.H{"error": "Authentication data not found"})
			c.Abort()
			return
		}

		auth := authData.(models.Auth)

		isSuperAdmin, err := auth.IsSuperAdmin()
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		if !isSuperAdmin {
			c.JSON(http.StatusForbidden, gin.H{"error": "Insufficient permissions"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// hasRequiredPermissions checks if the user has all required permissions
func hasRequiredPermissions(userPermissions []string, requiredPermissions []string) bool {
	permissionMap := make(map[string]bool)
	for _, perm := range userPermissions {
		permissionMap[perm] = true
	}

	for _, reqPerm := range requiredPermissions {
		if !permissionMap[reqPerm] {
			return false
		}
	}
	return true
}
