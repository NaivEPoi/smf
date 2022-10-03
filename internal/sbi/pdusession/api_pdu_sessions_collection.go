/*
 * Nsmf_PDUSession
 *
 * SMF PDU Session Service
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package pdusession

import (
	"net/http"

	"github.com/free5gc/openapi"
	smf_context "github.com/free5gc/smf/internal/context"
	"github.com/gin-gonic/gin"
)

// PostPduSessions - Create
func PostPduSessions(c *gin.Context) {
	scopes := []string{"nsmf-pdusession"}
	_, oauth_err := openapi.CheckOAuth(c.Request.Header.Get("Authorization"), scopes)
	if oauth_err != nil && smf_context.SMF_Self().OAuth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": oauth_err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
