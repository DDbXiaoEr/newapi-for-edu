package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/setting/group_channel_setting"
	"github.com/QuantumNous/new-api/setting/ratio_setting"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func TestUpdateGroupChannelBindingAPI(t *testing.T) {
	gin.SetMode(gin.TestMode)

	originalDB := model.DB
	originalCache := common.MemoryCacheEnabled
	originalOptionMap := common.OptionMap
	originalRatios := ratio_setting.GroupRatio2JSONString()
	originalBindings := group_channel_setting.GetGroupChannelsCopy()
	common.MemoryCacheEnabled = false
	if common.OptionMap == nil {
		common.OptionMap = map[string]string{}
	}

	dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared", strings.ReplaceAll(t.Name(), "/", "_"))
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&model.Channel{}, &model.Option{}))
	model.DB = db

	t.Cleanup(func() {
		model.DB = originalDB
		common.MemoryCacheEnabled = originalCache
		common.OptionMap = originalOptionMap
		group_channel_setting.SetGroupChannels(originalBindings)
		require.NoError(t, ratio_setting.UpdateGroupRatioByJSONString(originalRatios))
		sqlDB, sqlErr := db.DB()
		if sqlErr == nil {
			_ = sqlDB.Close()
		}
	})

	require.NoError(t, ratio_setting.UpdateGroupRatioByJSONString(`{"default":1,"vip":2}`))
	channel := &model.Channel{Name: "bind-channel", Key: "k", Status: common.ChannelStatusEnabled, Models: "m", Group: "default"}
	require.NoError(t, db.Create(channel).Error)

	body := fmt.Sprintf(`{"group":"default","channel_ids":[%d]}`, channel.Id)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest(http.MethodPut, "/api/group/channels", strings.NewReader(body))

	UpdateGroupChannelBinding(ctx)

	require.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), `"success":true`)
	ids, ok := group_channel_setting.GetBoundChannelIDs("default")
	require.True(t, ok)
	assert.Equal(t, []int{channel.Id}, ids)

	// Reject an unknown group.
	recorder = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest(http.MethodPut, "/api/group/channels", strings.NewReader(`{"group":"missing","channel_ids":[1]}`))
	UpdateGroupChannelBinding(ctx)
	assert.Contains(t, recorder.Body.String(), `"success":false`)
}
