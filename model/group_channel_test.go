package model

import (
	"testing"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/setting/group_channel_setting"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func setupGroupChannelTest(t *testing.T) {
	t.Helper()
	truncateTables(t)
	require.NoError(t, DB.Exec("DELETE FROM abilities").Error)
	require.NoError(t, DB.Exec("DELETE FROM channels").Error)

	previous := group_channel_setting.GetGroupChannelsCopy()
	t.Cleanup(func() { group_channel_setting.SetGroupChannels(previous) })
	group_channel_setting.SetGroupChannels(nil)
}

func createGroupChannel(t *testing.T, name string, group string, models string) *Channel {
	t.Helper()
	channel := &Channel{
		Name:   name,
		Key:    "test-key",
		Status: common.ChannelStatusEnabled,
		Models: models,
		Group:  group,
	}
	require.NoError(t, DB.Create(channel).Error)
	require.NoError(t, channel.UpdateAbilities(nil))
	return channel
}

func TestGroupEnabledModelsUsesModelBinding(t *testing.T) {
	setupGroupChannelTest(t)
	group := "bound-model-group"
	bound := createGroupChannel(t, "bound", group, "model-a,model-shared")
	other := createGroupChannel(t, "other", group, "model-b,model-shared")

	group_channel_setting.SetGroupChannels(map[string]map[string][]int{
		group: {"model-a": {bound.Id}},
	})

	assert.ElementsMatch(t, []string{"model-a", "model-b", "model-shared"}, GetGroupEnabledModels(group))
	assert.True(t, GroupHasModelBinding(group, "model-a"))
	assert.False(t, GroupHasModelBinding(group, "model-b"))

	group_channel_setting.SetGroupChannels(map[string]map[string][]int{
		group: {group_channel_setting.WildcardModel: {bound.Id}},
	})
	assert.ElementsMatch(t, []string{"model-a", "model-shared"}, GetGroupEnabledModels(group))

	group_channel_setting.SetGroupChannels(map[string]map[string][]int{group: {}})
	assert.False(t, GroupHasModelBinding(group, "model-a"))
	assert.ElementsMatch(t, []string{"model-a", "model-b", "model-shared"}, GetGroupEnabledModels(group))

	var count int64
	require.NoError(t, DB.Model(&Ability{}).Where("channel_id IN ?", []int{bound.Id, other.Id}).Count(&count).Error)
	assert.EqualValues(t, 4, count)
}

func TestChannelAllowedForGroupWithModelBinding(t *testing.T) {
	setupGroupChannelTest(t)
	group := "bound-allow-group"
	bound := createGroupChannel(t, "bound", group, "model-a,model-b")
	other := createGroupChannel(t, "other", group, "model-a,model-b")

	previousCache := common.MemoryCacheEnabled
	common.MemoryCacheEnabled = false
	t.Cleanup(func() { common.MemoryCacheEnabled = previousCache })

	group_channel_setting.SetGroupChannels(map[string]map[string][]int{
		group: {"model-a": {bound.Id}},
	})

	assert.True(t, ChannelAllowedForGroup(group, "model-a", bound.Id))
	assert.False(t, ChannelAllowedForGroup(group, "model-a", other.Id))
	assert.True(t, ChannelAllowedForGroup(group, "model-b", other.Id))
	assert.False(t, ChannelAllowedForGroup(group, "model-missing", bound.Id))

	group_channel_setting.SetGroupChannels(nil)
	assert.True(t, ChannelAllowedForGroup(group, "model-a", other.Id))
}

func TestGetChannelHonorsModelBindingDBPath(t *testing.T) {
	setupGroupChannelTest(t)
	group := "bound-db-group"
	legacy := createGroupChannel(t, "legacy", group, "shared-model,legacy-only-model")
	bound := createGroupChannel(t, "bound", group, "shared-model")

	previousCache := common.MemoryCacheEnabled
	common.MemoryCacheEnabled = false
	t.Cleanup(func() { common.MemoryCacheEnabled = previousCache })

	group_channel_setting.SetGroupChannels(map[string]map[string][]int{
		group: {"shared-model": {bound.Id}},
	})

	for i := 0; i < 10; i++ {
		channel, err := GetChannel(group, "shared-model", 0, nil)
		require.NoError(t, err)
		require.NotNil(t, channel)
		assert.Equal(t, bound.Id, channel.Id, "bound model must not select unbound channel %d", legacy.Id)
	}

	channel, err := GetChannel(group, "legacy-only-model", 0, nil)
	require.NoError(t, err)
	require.NotNil(t, channel)
	assert.Equal(t, legacy.Id, channel.Id)

	group_channel_setting.SetGroupChannels(map[string]map[string][]int{
		group: {group_channel_setting.WildcardModel: {bound.Id}},
	})
	channel, err = GetChannel(group, "shared-model", 0, nil)
	require.NoError(t, err)
	require.NotNil(t, channel)
	assert.Equal(t, bound.Id, channel.Id)
	channel, err = GetChannel(group, "legacy-only-model", 0, nil)
	require.NoError(t, err)
	assert.Nil(t, channel)
}

func TestGetRandomSatisfiedChannelHonorsModelBindingCachePath(t *testing.T) {
	setupGroupChannelTest(t)
	group := "bound-cache-group"
	legacy := createGroupChannel(t, "legacy", group, "cache-model,other-model")
	bound := createGroupChannel(t, "bound", group, "cache-model")

	previousCache := common.MemoryCacheEnabled
	common.MemoryCacheEnabled = true
	t.Cleanup(func() { common.MemoryCacheEnabled = previousCache })

	group_channel_setting.SetGroupChannels(map[string]map[string][]int{
		group: {"cache-model": {bound.Id}},
	})
	InitChannelCache()

	for i := 0; i < 20; i++ {
		channel, err := GetRandomSatisfiedChannel(group, "cache-model", 0, nil)
		require.NoError(t, err)
		require.NotNil(t, channel)
		assert.Equal(t, bound.Id, channel.Id, "bound model must not select unbound channel %d", legacy.Id)
	}

	channel, err := GetRandomSatisfiedChannel(group, "other-model", 0, nil)
	require.NoError(t, err)
	require.NotNil(t, channel)
	assert.Equal(t, legacy.Id, channel.Id)
}
