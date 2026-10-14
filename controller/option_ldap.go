package controller

import (
	"errors"
	"strings"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/i18n"
	"github.com/QuantumNous/new-api/service"

	"github.com/gin-gonic/gin"
)

var testLDAPLogin = service.TestLDAPLogin

type ldapLoginTestRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestLDAPLogin(c *gin.Context) {
	var request ldapLoginTestRequest
	if err := common.DecodeJson(c.Request.Body, &request); err != nil {
		common.ApiErrorI18n(c, i18n.MsgInvalidParams)
		return
	}

	username := strings.TrimSpace(request.Username)
	if username == "" || request.Password == "" {
		common.ApiErrorMsg(c, "LDAP username and password are required")
		return
	}

	result, err := testLDAPLogin(c.Request.Context(), username, request.Password)
	if err != nil {
		common.ApiErrorMsg(c, ldapLoginTestMessage(err))
		return
	}
	if result == nil {
		result = &service.LDAPLoginTestResult{}
	}
	if result.AttributeNames == nil {
		result.AttributeNames = []string{}
	}

	common.ApiSuccess(c, result)
}

func ldapLoginTestMessage(err error) string {
	switch {
	case errors.Is(err, service.ErrLDAPInvalidCredentials):
		return "LDAP username or password is incorrect"
	case errors.Is(err, service.ErrLDAPNoResult):
		return "LDAP user not found"
	case errors.Is(err, service.ErrLDAPMultipleResults):
		return "Multiple LDAP users matched this username"
	case errors.Is(err, service.ErrLDAPConfig):
		return "LDAP configuration error"
	default:
		return err.Error()
	}
}
