/*
 * Nsmf_EventExposure
 *
 * Session Management Event Exposure Service API
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package eventexposure

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/free5gc/openapi"
	smf_context "github.com/free5gc/smf/internal/context"
)

// SubscriptionsPost -
func SubscriptionsPost(c *gin.Context) {
	scopes := []string{"nsmf-event-exposure"}
	_, oauth_err := openapi.CheckOAuth(c.Request.Header.Get("Authorization"), scopes)
	if oauth_err != nil && smf_context.SMF_Self().OAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": oauth_err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// SubscriptionsSubIdDelete -
func SubscriptionsSubIdDelete(c *gin.Context) {
	scopes := []string{"nsmf-event-exposure"}
	_, oauth_err := openapi.CheckOAuth(c.Request.Header.Get("Authorization"), scopes)
	if oauth_err != nil && smf_context.SMF_Self().OAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": oauth_err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// SubscriptionsSubIdGet -
func SubscriptionsSubIdGet(c *gin.Context) {
	scopes := []string{"nsmf-event-exposure"}
	_, oauth_err := openapi.CheckOAuth(c.Request.Header.Get("Authorization"), scopes)
	if oauth_err != nil && smf_context.SMF_Self().OAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": oauth_err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

// SubscriptionsSubIdPut -
func SubscriptionsSubIdPut(c *gin.Context) {
	scopes := []string{"nsmf-event-exposure"}
	_, oauth_err := openapi.CheckOAuth(c.Request.Header.Get("Authorization"), scopes)
	if oauth_err != nil && smf_context.SMF_Self().OAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": oauth_err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
