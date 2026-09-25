package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/model"
	"github.com/QuantumNous/new-api/pkg/testdb"
	"github.com/QuantumNous/new-api/setting/group_channel_setting"
	"github.com/QuantumNous/new-api/setting/ratio_setting"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUpdateGroupChannelBindingAPI(t *testing.T) {
	gin.SetMode(gin.TestMode)

	originalCache := common.MemoryCacheEnabled
	originalOptionMap := common.OptionMap
	originalRatios := ratio_setting.GroupRatio2JSONString()
	originalBindings := group_channel_setting.GetGroupChannelsCopy()
	common.MemoryCacheEnabled = false
	if common.OptionMap == nil {
		common.OptionMap = map[string]string{}
	}
	db := testdb.OpenBound(t, &model.DB, &model.LOG_DB, model.InitColumnNames, testdb.Options{Models: []any{&model.Channel{}, &model.Option{}}})
	t.Cleanup(func() {
		common.MemoryCacheEnabled = originalCache
		common.OptionMap = originalOptionMap
		group_channel_setting.SetGroupChannels(originalBindings)
		require.NoError(t, ratio_setting.UpdateGroupRatioByJSONString(originalRatios))
	})

	require.NoError(t, ratio_setting.UpdateGroupRatioByJSONString(`{"default":1,"vip":2}`))
	channel := &model.Channel{Name: "bind-channel", Key: "k", Status: common.ChannelStatusEnabled, Models: "m", Group: "default"}
	require.NoError(t, db.Create(channel).Error)

	body := fmt.Sprintf(`{"group":"default","model":"gpt-4","channel_ids":[%d]}`, channel.Id)
	recorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest(http.MethodPut, "/api/group/channels", strings.NewReader(body))

	UpdateGroupChannelBinding(ctx)

	require.Equal(t, http.StatusOK, recorder.Code)
	assert.Contains(t, recorder.Body.String(), `"success":true`)
	ids, ok := group_channel_setting.GetBoundChannelIDs("default", "gpt-4")
	require.True(t, ok)
	assert.Equal(t, []int{channel.Id}, ids)

	recorder = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest(http.MethodPut, "/api/group/channels", strings.NewReader(`{"group":"missing","model":"gpt-4","channel_ids":[1]}`))
	UpdateGroupChannelBinding(ctx)
	assert.Contains(t, recorder.Body.String(), `"success":false`)

	recorder = httptest.NewRecorder()
	ctx, _ = gin.CreateTestContext(recorder)
	ctx.Request = httptest.NewRequest(http.MethodPut, "/api/group/channels", strings.NewReader(`{"group":"default","channel_ids":[1]}`))
	UpdateGroupChannelBinding(ctx)
	assert.Contains(t, recorder.Body.String(), `"success":false`)
}
