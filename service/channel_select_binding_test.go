package service

import (
	"net/http/httptest"
	"testing"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/constant"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/setting/group_channel_setting"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestResolvePinnedGroupHonorsBinding(t *testing.T) {
	db := setupChannelSelectAutoGroupsTest(t)
	const modelName = "pin-binding-model"
	createChannelSelectAutoGroupsChannel(t, db, 2201, "vip", modelName)
	createChannelSelectAutoGroupsChannel(t, db, 2202, "default", modelName)
	model.InitChannelCache()

	previous := group_channel_setting.GetGroupChannelsCopy()
	t.Cleanup(func() { group_channel_setting.SetGroupChannels(previous) })
	group_channel_setting.SetGroupChannels(map[string][]int{"default": {2202}})

	gin.SetMode(gin.TestMode)

	newCtx := func() *gin.Context {
		ctx, _ := gin.CreateTestContext(httptest.NewRecorder())
		common.SetContextKey(ctx, constant.ContextKeyUserGroup, "default")
		common.SetContextKey(ctx, constant.ContextKeyTokenAutoGroups, []string{"vip", "default"})
		return ctx
	}

	// Non-auto: only the bound channel is allowed.
	if _, allowed := resolvePinnedGroup(newCtx(), "default", modelName, 2202); !allowed {
		t.Fatalf("bound channel 2202 must be allowed for group default")
	}
	if _, allowed := resolvePinnedGroup(newCtx(), "default", modelName, 2201); allowed {
		t.Fatalf("unbound channel 2201 must be rejected for group default")
	}

	// Auto: 2202 resolves through the bound "default" group, 2201 through the
	// unbound "vip" legacy membership.
	ctx := newCtx()
	group, allowed := resolvePinnedGroup(ctx, "auto", modelName, 2202)
	assert.True(t, allowed)
	assert.Equal(t, "default", group)
	assert.Equal(t, "default", common.GetContextKeyString(ctx, constant.ContextKeyAutoGroup))

	ctx = newCtx()
	group, allowed = resolvePinnedGroup(ctx, "auto", modelName, 2201)
	assert.True(t, allowed)
	assert.Equal(t, "vip", group)
	assert.Equal(t, "vip", common.GetContextKeyString(ctx, constant.ContextKeyAutoGroup))
}
